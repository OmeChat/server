package storage

type Service struct{}

// UserModel is identified by
// random generated hash
type UserModel struct {
	Clients  []string `json:"clients"`
	Age      int      `json:"age"`
	Secret   string   `json:"secret"`
	Username string   `json:"username"`
}

// ClientModel is identified by a unique
// hash generated for the client system
type ClientModel struct {
	Owner       string `json:"owner"`
	AccessToken string `json:"access_token"`
}

// The ExposedUser is the UserModel, but
// with the personal data removed and the
// userHash added
type ExposedUser struct {
	Age      int    `json:"age"`
	Username string `json:"username"`
	UserHash string `json:"user_hash"`
}
