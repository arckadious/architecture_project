package out

import (
	"auth-api-crossfitlov/models/structs/in"
)

type TokenInfos struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	ExpiresAt string `json:"expiresAt"`
}
type TokenData struct {
	TokenInfos TokenInfos  `json:"tokenInfos"`
	UserInfos  interface{} `json:"userInfos"`
}

type UserInfos struct {
	CrossfitlovID int    `json:"crossfitlovID"`
	Firstname     string `json:"firstname"`
	Gender        string `json:"gender"`
	Age           int    `json:"age"`
	Email         string `json:"email"`
	BoxCity       string `json:"boxCity"`
	Biography     string `json:"biography"`
	Job           string `json:"job"`
	CreatedAt     string `json:"createdAt"`
}

type UserRegisterInfos struct {
	Credentials in.CredentialsData `json:"credentialsData"`
	UserInfos   UserInfos          `json:"userInfos"`
}
