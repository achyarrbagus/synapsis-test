package userdto

type UsersResponse struct {
	Name  string
	Email string
}

type LoginResponse struct {
	Email string
	Token string
}
