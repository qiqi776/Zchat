package request

type Loginrequest struct {
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}