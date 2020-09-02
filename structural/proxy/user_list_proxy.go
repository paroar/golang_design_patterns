package proxy

import "errors"

type UserListProxy struct {
	SomeDatabase           UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		u.DidLastSearchUsedCache = true
		return user, nil
	}

	user, err = u.SomeDatabase.FindUser(id)
	if err == nil {
		u.AddUserToCache(user)
		u.DidLastSearchUsedCache = false
		return user, nil
	}

	u.DidLastSearchUsedCache = false
	return User{}, errors.New("User not found")
}

func (u *UserListProxy) AddUserToCache(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.AddUser(user)
	}
}
