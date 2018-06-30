package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()
	var data []byte
	var err error
	switch flag.NArg() {
	case 0:
		data, err = ioutil.ReadAll(os.Stdin)
		check(err)
		headers := make(http.Header)
		headers.Add("Origin", "http://localhost")
		c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8546", headers)
		if err != nil {
			log.Fatal("dial:", err)
		}
		err = c.WriteMessage(websocket.TextMessage, []byte(string(data)))
		if err != nil {
			log.Println("write:", err)
			return
		}
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		fmt.Printf("%s\n", message)
		defer c.Close()
		break
	default:
		fmt.Printf("input must be from stdin or file\n")
		os.Exit(1)
	}
}
