package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // default upgrader options

func handleIncoming(w http.ResponseWriter, r *http.Request) {
	// upgrade incoming http connection to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("err in upgrading connection")
		fmt.Println(err)
	}
	defer conn.Close()

	for {
		// receive message
		msgType, msgBuf, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		} else if msgType != websocket.TextMessage {
			fmt.Println("server only accepts text messages")
		}

		fmt.Println("received msg:", string(msgBuf))
		fmt.Println("sending back acknowledgement")

		// send ack
		if err := conn.WriteMessage(websocket.TextMessage, []byte("recvd your msg")); err != nil {
			fmt.Println("error sending acknowledgement")
		}
	}

}

func main() {
	http.HandleFunc("/", handleIncoming)

	http.ListenAndServe(":5050", nil)
}
