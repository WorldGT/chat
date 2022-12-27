package types

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
