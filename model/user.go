package model

type User struct {
	Email string `json:"email"`
}

func NewUser(email string) *User {
	return &User{
		Email: email,
	}
}

type UserStore interface {
	FindAll() []*User
	Create(user *User) (*User, error)
}

type LocalUserStore struct {
	users []*User
}

func NewLocalUserStore() *LocalUserStore {
	return &LocalUserStore{
		users: make([]*User, 0),
	}
}

func (s *LocalUserStore) FindAll() []*User {
	return s.users
}

func (s *LocalUserStore) Create(user *User) (*User, error) {
	s.users = append(s.users, user)
	return user, nil
}
