package storage

// CheckIfUserExists checks weather there is already an user identified
// by the given username
func (s Service) CheckIfUserExists(username string) bool {
	for _, usr := range (UserModel{}).getAllUser() {
		if usr.Username == username {
			return true
		}
	}
	return false
}

// CheckUserAccess checks if the user SECRET is correct and returns
// this state as a boolean value
func (s Service) CheckUserAccess(username string, secret string) bool {
	if !s.CheckIfUserExists(username) {
		return false
	}
	usr := (UserModel{}).GetUserByUsername(username)
	return usr.Secret == secret
}
