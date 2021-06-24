package storage

// CheckAuthStatus checks, if the client has the permission to access
// the private data of the given userHash. It returns this
// state as a boolean value
func CheckAuthStatus(userHash string, clientHash string, accessToken string) bool {
	users := (UserModel{}).getAllUser()
	if _, ok := users[userHash]; !ok {
		return false
	}
	userClients := users[userHash].Clients

	clientExists := false

	for _, client := range userClients {
		if client == clientHash {
			clientExists = true
			break
		}
	}
	if !clientExists {
		return false
	}
	allClients := ClientModel{}.getAllClients()
	return allClients[clientHash].AccessToken == accessToken
}
