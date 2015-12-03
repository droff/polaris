package models

// Users type
type Users []User

// User model
type User struct {
	ID             int
	Username       string
	Password       string
	HasVideoAccess bool
}

// Userser describe whole logic of User model
type Userser interface {
	Find(int) *User
	FindByLogin(string) *User
}

// Find user by ID
func (users *Users) Find(id int) *User {
	for _, user := range *users {
		if user.ID == id {
			return &user
		}
	}

	return nil
}

// FindByLogin search user by login
func (users *Users) FindByLogin(value string) *User {
	for _, user := range *users {
		if user.Username == value {
			return &user
		}
	}

	return nil
}
