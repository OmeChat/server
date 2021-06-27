package storage

import (
	"encoding/json"
	"github.com/OmeChat/server/pkg/hashing"
	"github.com/OmeChat/server/pkg/random"
	"io/ioutil"
)

// getAllUser returns an map of all users.
// the key is the hash of the userID and the
// value is the general UserModel struct of
// the specific user identified by the hash
func (s UserModel) getAllUser() map[string]UserModel {
	data, err := ioutil.ReadFile("./data/user.json")
	if err != nil {
		panic(err.Error())
	}
	var user map[string]UserModel
	if err := json.Unmarshal(data, &user); err != nil {
		panic(err.Error())
	}
	return user
}

// hashExists checks, if there is already an
// user existing identified by the given hash.
// It returns this state as a boolean value
func (s UserModel) hashExists(hash string) bool {
	for h, _ := range s.getAllUser() {
		if h == hash {
			return true
		}
	}
	return false
}

// CreateUserAccount creates a new user account trough the
// given values and inserts it into the user.json in the
// data folder. After that it returns the complete struct
// of the new user instance and the hash of the user
func (s UserModel) CreateUserAccount(username string, age int) (UserModel, string) {
	hash := hashing.SHA512(random.GenerateUID(64))
	for {
		if !s.hashExists(hash) {
			break
		}
		hash = hashing.SHA512(random.GenerateUID(64))
	}
	secret := random.GenerateUID(32)
	users := s.getAllUser()
	user := UserModel{
		Clients:  []string{},
		Age:      age,
		Secret:   secret,
		Username: username,
	}
	users[hash] = user
	jsonString, _ := json.Marshal(users)
	_ = ioutil.WriteFile("./data/user.json", jsonString, 0644)
	return user, hash
}

// GetUserByUsername requests the user identified by
// the given username. If there is no existing user
// an empty UserModel will be returned
func (s UserModel) GetUserByUsername(username string) UserModel {
	for _, usr := range s.getAllUser() {
		if usr.Username == username {
			return usr
		}
	}
	return UserModel{}
}

// GetHashByUsername tries to get the hash of the user
// identified by the given username.
func (s UserModel) GetHashByUsername(username string) string {
	for hash, usr := range s.getAllUser() {
		if usr.Username == username {
			return hash
		}
	}
	return ""
}

// GetUserByHash returns an instance of the UserModel,
// which contains information about the requested user
// identified by the given hash
func (s UserModel) GetUserByHash(userHash string) UserModel {
	for hash, usr := range s.getAllUser() {
		if hash == userHash {
			return usr
		}
	}
	return UserModel{}
}

// AddClientToUser adds a new clientHash to he array of the user
// identified by the given userHash. After that all the data is saved
// to the user.json file in the local storage.
func (s UserModel) AddClientToUser(userHash string, clientHash string) {
	users := s.getAllUser()
	userModel := users[userHash]
	userModel.Clients = append(userModel.Clients, clientHash)
	users[userHash] = userModel
	jsonString, _ := json.Marshal(users)
	_ = ioutil.WriteFile("./data/user.json", jsonString, 0644)
}

// GetUserAtAgeWithTolerance searches for user at the same age. There are also
// an tolerance given. It is the tolerance for the age. If the tolerance is 2,
// users at an age 2 years below or upon the given age are also fetched
func (s UserModel) GetUserAtAgeWithTolerance(age int, tolerance int, userHash string) []ExposedUser {
	users := s.getAllUser()
	var retunedUser []ExposedUser
	for hash, usr := range users {
		if (usr.Age < age+tolerance || usr.Age > age-tolerance) && hash != userHash {
			retunedUser = append(retunedUser, ExposedUser{
				Age:      usr.Age,
				Username: usr.Username,
				UserHash: hash,
			})
		}
	}
	return retunedUser
}
