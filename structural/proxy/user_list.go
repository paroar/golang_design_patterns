package proxy

import "errors"

type UserList []User

func (u *UserList) FindUser(id int32) (User, error) {
	for _, v := range *u {
		if v.ID == id {
			return v, nil
		}
	}
	return User{}, errors.New("User not found")
}

func (u *UserList) AddUser(user User) {
	*u = append(*u, user)
}
