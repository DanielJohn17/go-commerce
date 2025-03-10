package types

type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"Email"`
	Password  string `json:"password"`
}
