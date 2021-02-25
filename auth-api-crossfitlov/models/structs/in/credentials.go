package in

// Create a struct to read the username and password from the request body
//login is a username or email
type CredentialsData struct {
	Password string `json:"password"`
	Login    string `json:"login"`
	GetInfos bool   `json:"getInfos"`
}
