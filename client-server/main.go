package main

import (
	"bytes"
	"context"
	"dashboard/client-server/devicepb"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

const (
	protocol = "unix"
	sockAddr = "/tmp/echo.sock"
)

func main() {
	cleanup := func() {
		if _, err := os.Stat(sockAddr); err == nil {
			if err := os.RemoveAll(sockAddr); err != nil {
				log.Fatal(err)
			}
		}
	}
	cleanup()
	listener, err := net.Listen(protocol, sockAddr)
	if err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		fmt.Println("ctrl-c pressed..")
		close(quit)
		cleanup()
		os.Exit(0)
	}()

	fmt.Println("server launched...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(">>> accepted")
		go echo(conn)
	}
}

func echo(conn net.Conn) {
	defer conn.Close()
	log.Printf("Connected: %s\n", conn.RemoteAddr().Network())

	buf := &bytes.Buffer{}
	_, err := io.Copy(buf, conn)
	if err != nil {
		log.Println(err)
		return
	}

	s := buf.String()
	bt := []byte(s)
	var d devicepb.Device
	error := json.Unmarshal(bt, &d)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(d)

	buf.Reset()
	buf.WriteString(s)

	_, err = io.Copy(conn, buf)
	if err != nil {
		log.Println(err)
		return
	}

	connection, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := devicepb.NewDeviceServiceClient(connection)

	_, errorG := c.DeviceUpdate(context.Background(), &devicepb.DeviceUpdateRequest{Device: &d})
	if errorG != nil {
		log.Fatal(errorG)
	}
}
