package websocket

import "github.com/gofiber/websocket/v2"

var WS_CONNECTIONS map[string][]*websocket.Conn

var WS_DATAFLOW_CHANNEL chan ConnectionPair

// The DataHandler is an extra goroutine running
// to store all the given connections into a
// map. If there is already an client connected
// with the same hash, the new connection will be
// added to the list of connected clients
func DataHandler() {
	for {
		pair := <-WS_DATAFLOW_CHANNEL
		if WS_CONNECTIONS[pair.Hash] == nil {
			WS_CONNECTIONS[pair.Hash] = []*websocket.Conn{pair.Connection}
		} else {
			WS_CONNECTIONS[pair.Hash] = append(WS_CONNECTIONS[pair.Hash], pair.Connection)
		}
	}
}
