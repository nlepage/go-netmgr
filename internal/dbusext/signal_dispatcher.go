package dbusext

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/godbus/dbus/v5"
)

type (
	SignalKey struct {
		path dbus.ObjectPath
		name string
	}

	outChan struct {
		value   reflect.Value
		convert func(interface{}) reflect.Value
	}

	SignalDispatcher struct {
		l    sync.RWMutex
		in   <-chan *dbus.Signal
		outs map[SignalKey]map[interface{}]outChan
	}
)

var SignalDispatcherKey = struct{}{}

func NewSignalDispatcher() *SignalDispatcher {
	return &SignalDispatcher{
		outs: make(map[SignalKey]map[interface{}]outChan),
	}
}

func (sm *SignalDispatcher) Signal(conn *dbus.Conn, path dbus.ObjectPath, iface, member string, elemType reflect.Type, out interface{}, convert interface{}) error {
	oc, err := newOutChan(out, elemType, convert)
	if err != nil {
		return err
	}

	sm.l.Lock()
	defer sm.l.Unlock()

	if sm.in == nil {
		var in = make(chan *dbus.Signal)
		sm.in = in
		conn.Signal(in)
		go sm.pipe(conn.Context().Done())
	}

	var k = SignalKey{path, iface + "." + member}

	if _, ok := sm.outs[k]; !ok {
		sm.outs[k] = make(map[interface{}]outChan)
		if err := conn.AddMatchSignal(
			dbus.WithMatchObjectPath(path),
			dbus.WithMatchInterface(iface),
			dbus.WithMatchMember(member),
		); err != nil {
			return err
		}
	}
	if _, ok := sm.outs[k][out]; !ok {
		sm.outs[k][out] = oc
	}

	return nil
}

func (sm *SignalDispatcher) pipe(done <-chan struct{}) {
	for {
		select {
		case s := <-sm.in:
			sm.pipeSignal(s)
		case <-done:
			return
		}
	}
}

func (sm *SignalDispatcher) pipeSignal(s *dbus.Signal) {
	sm.l.RLock()
	defer sm.l.RUnlock()

	if outs, ok := sm.outs[SignalKey{s.Path, s.Name}]; ok {
		for _, ch := range outs {
			for _, v := range s.Body {
				ch.Send(v)
			}
		}
	}
}

func newOutChan(ch interface{}, inType reflect.Type, convert interface{}) (outChan, error) {
	chType := reflect.TypeOf(ch)
	if chType.Kind() != reflect.Chan {
		return outChan{}, errors.New("ch is not a chan")
	}

	oc := outChan{
		value: reflect.ValueOf(ch),
	}

	chElemType := chType.Elem()
	if convert != nil {
		convertType := reflect.TypeOf(convert)
		if convertType.Kind() != reflect.Func {
			return outChan{}, errors.New("convert is not a func")
		}
		if convertType.NumIn() != 1 || convertType.In(0) != inType || convertType.NumOut() != 1 || convertType.Out(0) != chElemType {
			return outChan{}, fmt.Errorf("convert type should be func(%s) %s", inType, chElemType)
		}
		convertValue := reflect.ValueOf(convert)
		oc.convert = func(v interface{}) reflect.Value {
			return convertValue.Call([]reflect.Value{reflect.ValueOf(v)})[0]
		}
	} else {
		if chElemType == inType {
			oc.convert = reflect.ValueOf
		} else {
			if !inType.ConvertibleTo(chElemType) {
				return outChan{}, fmt.Errorf("%s is not convertible to %s", inType, chElemType)
			}
			oc.convert = func(v interface{}) reflect.Value {
				return reflect.ValueOf(v).Convert(chElemType)
			}
		}
	}

	return oc, nil
}

func (oc outChan) Send(v interface{}) {
	oc.value.Send(oc.convert(v))
}
