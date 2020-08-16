package adapters

type LoginAdapter struct {
	Token string `json:"token"`
}

func NewLoginAdapter(token string) LoginAdapter {
	return LoginAdapter{
		Token: token,
	}
}