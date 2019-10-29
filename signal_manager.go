package wpa

import (
	"sync"

	"github.com/godbus/dbus/v5"
)

type signalKey struct {
	name string
	path dbus.ObjectPath
}

type SignalManager struct {
	conn *dbus.Conn
	l    sync.Mutex
	in   <-chan *dbus.Signal
	outs map[signalKey]map[chan<- []interface{}]bool
}

func NewSignalManager(conn *dbus.Conn) *SignalManager {
	return &SignalManager{
		conn: conn,
		outs: make(map[signalKey]map[chan<- []interface{}]bool),
	}
}

func (sm *SignalManager) Signal(iface, member string, path dbus.ObjectPath, out chan<- []interface{}) error {
	sm.l.Lock()
	defer sm.l.Unlock()

	if sm.in == nil {
		var in = make(chan *dbus.Signal, 10)
		sm.in = in
		sm.conn.Signal(in)
		go sm.pipe()
	}

	var k = signalKey{iface + "." + member, path}

	if _, ok := sm.outs[k]; !ok {
		sm.outs[k] = make(map[chan<- []interface{}]bool)
		if err := sm.conn.AddMatchSignal(
			dbus.WithMatchInterface(iface),
			dbus.WithMatchMember(member),
			dbus.WithMatchObjectPath(path),
		); err != nil {
			return err
		}
	}
	if _, ok := sm.outs[k][out]; !ok {
		sm.outs[k][out] = true
	}

	return nil
}

func (sm *SignalManager) pipe() {
	for {
		select {
		case s := <-sm.in:
			sm.pipeSignal(s)
		case <-sm.conn.Context().Done():
		}
	}
}

func (sm *SignalManager) pipeSignal(s *dbus.Signal) {
	sm.l.Lock()
	defer sm.l.Unlock()

	var k = signalKey{s.Name, s.Path}

	if outs, ok := sm.outs[k]; ok {
		for out := range outs {
			out <- s.Body
		}
	}
}
