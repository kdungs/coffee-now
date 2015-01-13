package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type CoffeePerson struct {
	Name     string
	Lat, Lng float64
}

type CoffeeConnection struct {
	Hub    *CoffeeHub
	Ws     *websocket.Conn
	Send   chan []byte
	Person CoffeePerson
}

func (cc *CoffeeConnection) Reader() {
	for {
		_, message, err := cc.Ws.ReadMessage()
		if err != nil {
			break
		}
		log.Println(message)
		cc.Hub.Broadcast <- message
	}
}

func (cc *CoffeeConnection) Writer() {
	for message := range cc.Send {
		err := cc.Ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	cc.Ws.Close()
}

type CoffeeHub struct {
	Connections map[*CoffeeConnection]bool
	Broadcast   chan []byte
	Register    chan *CoffeeConnection
	Unregister  chan *CoffeeConnection
}

func NewCoffeeHub() *CoffeeHub {
	return &CoffeeHub{
		make(map[*CoffeeConnection]bool),
		make(chan []byte),
		make(chan *CoffeeConnection),
		make(chan *CoffeeConnection)}
}

func (ch *CoffeeHub) Run() {
	for {
		select {
		case c := <-ch.Register:
			ch.Connections[c] = true
		case c := <-ch.Unregister:
			if _, ok := ch.Connections[c]; ok {
				delete(ch.Connections, c)
				close(c.Send)
			}
		case m := <-ch.Broadcast:
			for c := range ch.Connections {
				select {
				case c.Send <- m:
				default:
					delete(ch.Connections, c)
					close(c.Send)
				}
			}
		}
	}
}

func main() {
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	hub := NewCoffeeHub()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		var person CoffeePerson
		err = ws.ReadJSON(&person)
		if err != nil {
			log.Println(err)
			return
			// Actually send error message and close connection...
		} else {
			log.Printf("%s registered.\n", person.Name)
			log.Printf("%s - %s\n", ws.LocalAddr(), ws.RemoteAddr())
			suc := struct {
				Status string
			}{"Success"}
			ws.WriteJSON(suc)
		}
		c := &CoffeeConnection{hub, ws, make(chan []byte, 256), person}
		hub.Register <- c
		defer func() { hub.Unregister <- c }()
		go c.Writer()
		c.Reader()
	})
	go hub.Run()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
