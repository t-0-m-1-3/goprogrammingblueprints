package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/gorilla/websocket"
)

// client represents a single chatting user
type client struct {
	//socket is the web socket for this client.
	socket *websocket.Conn
	// send is a channel on which messages are sent
	send chan []byte
	// room is the room this client is chatting in.
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}

//templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP requests
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	// root
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
