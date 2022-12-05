package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

const (
	protocol = "unix"
	sockAddr = "/tmp/echo.sock"
)

type Device struct {
	IdVendor   string `json:"id_vendor"`
	IdProducet string `json:"id_product"`
	Serial     string `json:"serial"`
	Action     string `json:"action"`
}

func main() {
	conn, err := net.Dial(protocol, sockAddr)
	if err != nil {
		log.Fatal(err)
	}
	dev := Device{os.Args[1], os.Args[2], os.Args[3], os.Args[4]}
	b, _ := json.Marshal(dev)

	_, err = conn.Write([]byte(b))
	if err != nil {
		log.Fatal(err)
	}

	err = conn.(*net.UnixConn).CloseWrite()
	if err != nil {
		log.Fatal(err)
	}

	b, error := ioutil.ReadAll(conn)
	if error != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
