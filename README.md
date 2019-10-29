<h1 align="center">Welcome to go-wpa 👋</h1>
<p>
  <a href="https://godoc.org/github.com/nlepage/go-wpa" target="_blank">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" />
  </a>
  <a href="https://spdx.org/licenses/Apache-2.0.html" target="_blank">
    <img alt="License: Apache 2.0" src="https://img.shields.io/badge/License-Apache 2.0-yellow.svg" />
  </a>
  <a href="https://twitter.com/njblepage" target="_blank">
    <img alt="Twitter: njblepage" src="https://img.shields.io/twitter/follow/njblepage.svg?style=social" />
  </a>
</p>

> Go bindings for wpa_supplicant D-Bus API

## Install

```sh
go get -u github.com/nlepage/go-wpa
```

## Usage

```go
package main

import (
	"fmt"

	wpa "github.com/nlepage/go-wpa"
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

	for i, iface := range ifaces {
		ifname, err := iface.Ifname()
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d: %s\n", i, ifname)
	}
}
```

## Author

👤 **Nicolas Lepage**

* Twitter: [@njblepage](https://twitter.com/njblepage)
* Github: [@nlepage](https://github.com/nlepage)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/nlepage/go-wpa/issues).

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2019 [Nicolas Lepage](https://github.com/nlepage).<br />
This project is [Apache 2.0](https://spdx.org/licenses/Apache-2.0.html) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_