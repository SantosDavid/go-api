package presenter

type LoginSucesseful struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
