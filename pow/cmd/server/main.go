package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"time"

	"pow/api"
	"pow/pow/hashcash"
)

var (
	addr = os.Getenv("ADDR")
)

type App struct {
	api api.API
}

var app = App{
	api: *api.NewAPI(hashcash.NewHashcash(time.Now, hashcash.RandInt32)),
}

func (a *App) handleRequest(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		conn.SetReadDeadline(time.Now().Add(time.Minute))
		conn.SetWriteDeadline(time.Now().Add(time.Minute))

		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error occured:", err)
			return
		}

		log.Printf("received a challenge: %q", string(message))

		resp, err := a.api.HandleVerify(message)
		if err != nil {
			log.Println("error occured:", err)
			conn.Write([]byte(err.Error() + "\n"))
		} else {
			conn.Write([]byte(resp + "\n"))
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":1313")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	log.Println("starting a tcp server")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go app.handleRequest(conn)
	}
}
