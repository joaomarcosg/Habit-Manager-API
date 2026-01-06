package user

type CreateUserReq struct {
	Name     string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
