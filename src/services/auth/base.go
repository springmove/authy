package auth

type IAuthProvider interface {
	Auth(reqData interface{}) (interface{}, error)
}
