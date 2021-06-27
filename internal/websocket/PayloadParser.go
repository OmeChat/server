package websocket

// The PayloadParser parses the payload of the given websocket request
// and parses it into a map. After that it checks, if all required parameters
// are given
func PayloadParser(action string, payload interface{}) (map[string]interface{}, bool) {
	data := payload.(map[string]interface{})
	switch action {
	case "exchange-key":
		return data, data["target_hash"] != nil && data["key"] != nil
	}
	return nil, false
}
