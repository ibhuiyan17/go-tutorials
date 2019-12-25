// +build ignore

package main

import (
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	url := url.URL{
		Scheme: "ws",
		Host:   "localhost:5050",
		Path:   "/",
	}
	fmt.Println("connecting to", url.String())

	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		fmt.Println("error dialing", err)
	}
	defer conn.Close()

	// send message
	msg := "hello client 1. From client 2"
	if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		fmt.Println("error writing message", err)
	}
	fmt.Println("sent msg")

	// receive ack
	_, msgBuf, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("error receiving ack")
	}
	fmt.Println("received ack:", string(msgBuf))

	// close connection
	fmt.Println("closing connection")
	conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}
