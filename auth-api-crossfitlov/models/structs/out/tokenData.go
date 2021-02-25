package out

import (
	"auth-api-crossfitlov/models/structs/in"
	"net/http"
)

type TokenData struct {
	CookieData http.Cookie
	UserInfos  interface{}
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
