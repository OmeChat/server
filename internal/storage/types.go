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
