package main

type room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan []byte
	// join is a channel for clients to enter a room
	join chan *client
	// leave is a channel for clients wanting to leave
	leave chan *client
	// clients holds all current clients in this room
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
		case client := <-r.leave:
			//leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward a message to all clients
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}
