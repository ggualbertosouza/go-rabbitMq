package userRbMq

type UserCreatedMessage struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdatingMessage struct {
	Email string `json:"email"`
}
