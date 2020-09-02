package proxy

type IUserFinder interface {
	FindUser(id int32) (User, error)
}
