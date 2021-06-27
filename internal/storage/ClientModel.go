package storage

import (
	"encoding/json"
	"github.com/OmeChat/server/pkg/hashing"
	"github.com/OmeChat/server/pkg/random"
	"io/ioutil"
)

// getAllClients returns a map of all clients.
// The key represents the hash of the clientID
// and the value is an instance of the ClientModel
func (s ClientModel) getAllClients() map[string]ClientModel {
	data, err := ioutil.ReadFile("./data/client.json")
	if err != nil {
		panic(err.Error())
	}
	var clients map[string]ClientModel
	if err := json.Unmarshal(data, &clients); err != nil {
		panic(err.Error())
	}
	return clients
}

// hashExists checks if there is already an client
// identified by the given hash
func (s ClientModel) hashExists(hash string) bool {
	for h := range s.getAllClients() {
		if h == hash {
			return true
		}
	}
	return false
}

// generateClientHash generates an unique client
// hash and returns it
func (s ClientModel) generateClientHash() string {
	hash := hashing.SHA512(random.GenerateUID(64))
	for {
		if !s.hashExists(hash) {
			break
		}
		hash = hashing.SHA512(random.GenerateUID(64))
	}
	return hash
}

// AddClient adds an new client for the user and returns its hash
// and ClientModel
func (s ClientModel) AddClient(userHash string) (string, ClientModel) {
	hash := s.generateClientHash()
	accessToken := random.GenerateUID(64)
	client := ClientModel{
		Owner:       userHash,
		AccessToken: accessToken,
	}
	clients := s.getAllClients()
	clients[hash] = client
	jsonString, _ := json.Marshal(clients)
	_ = ioutil.WriteFile("./data/client.json", jsonString, 0644)
	return hash, client
}
