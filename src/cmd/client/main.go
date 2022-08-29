package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"pow/client"
	"pow/pow/hashcash"
	"time"
)

type App struct {
	client *client.Client
}

var app = App{
	client: client.NewClient(hashcash.NewHashcash(time.Now, hashcash.RandInt32)),
}

var (
	addr = os.Getenv("ADDR")
	id   = os.Getenv("ID")
)

func main() {
	if addr == "" || id == "" {
		log.Println("please set ADDR and ID")
		return
	}

	conn, err := net.DialTimeout("tcp", addr, time.Second)
	if err != nil {
		log.Fatal(err)
		return
	}

	conn.SetReadDeadline(time.Now().Add(time.Second * 10)) // a minute should be enough
	conn.SetWriteDeadline(time.Now().Add(time.Minute))

	defer conn.Close()

	token := app.client.Generate(id)

	log.Println("generated a token:", token)

	io.WriteString(conn, token+"\n")

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("failed to read from conn:", err)
		return
	}

	log.Println(message)
}
