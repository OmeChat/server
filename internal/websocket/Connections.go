package websocket

var WS_CONNECTIONS map[string][]ConnectionIdentifier

var WS_DATAFLOW_CHANNEL chan ConnectionPair

// The DataHandler is an extra goroutine running
// to store all the given connections into a
// map. If there is already an client connected
// with the same hash, the new connection will be
// added to the list of connected clients
func DataHandler(dataChannel chan ConnectionPair) {
	for {
		pair := <-dataChannel
		if WS_CONNECTIONS[pair.UserHash] == nil {
			var arr []ConnectionIdentifier
			arr = append(arr, ConnectionIdentifier{
				Connection: pair.Connection,
				ClientHash: pair.ClientHash,
			})
			WS_CONNECTIONS[pair.UserHash] = arr
		} else {
			clientExists := false
			for _, el := range WS_CONNECTIONS[pair.UserHash] {
				if el.ClientHash == pair.ClientHash {
					clientExists = true
				}
			}
			if !clientExists {
				WS_CONNECTIONS[pair.UserHash] = append(WS_CONNECTIONS[pair.UserHash], ConnectionIdentifier{
					Connection: pair.Connection,
					ClientHash: pair.ClientHash,
				})
			}
		}
	}
}
