package main

import (
	"fmt"
	"time"

	wpa "github.com/nlepage/go-netmgr"
)

func main() {
	w, err := wpa.SystemWPA()
	if err != nil {
		panic(err)
	}
	defer w.Close()

	ifaces, err := w.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, iface := range ifaces {
		ifname, err := iface.Ifname()
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s:\n", ifname)

		nets, err := iface.Networks()
		if err != nil {
			panic(err)
		}

		for _, net := range nets {
			props, err := net.Properties()
			if err != nil {
				panic(err)
			}

			fmt.Printf("\t%#v\n", props)
		}

		ch := make(chan bool)
		if err := iface.ScanDone(ch); err != nil {
			panic(err)
		}
		go func() {
			if ok := <-ch; ok {
				fmt.Println("Scan Done !")
			}
		}()

		if err := iface.Scan(wpa.ScanOptions{
			Type: wpa.ScanActive,
		}); err != nil {
			panic(err)
		}
		fmt.Println("Scanned!")
	}

	time.Sleep(time.Minute)
}
