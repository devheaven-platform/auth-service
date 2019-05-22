package domain

// Token represents an auth token in the system. The
// object contains the token and de expire date of the
// token.
type Token struct {
	Token   string  `json:"token"`
	Expires float64 `json:"expires"`
}
