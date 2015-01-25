package storage

type User struct{}

func CurrentUser() User {
	var ret User
	return ret
}
