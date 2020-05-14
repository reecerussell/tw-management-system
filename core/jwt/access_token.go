package jwt

type AccessToken struct {
	AccessToken string  `json:"accessToken"`
	Expires     float64 `json:"expires"`
}
