package parameters

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

var Config *conf

type conf struct {
	Database struct {
		Adapter      string `json:"adapter"`
		Host         string `json:"host"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		Port         string `json:"port"`
		Name         string `json:"name"`
		Charset      string `json:"charset"`
		MaxOpenConns int    `json:"maxOpenConns"`
		MaxIdleConns int    `json:"maxIdleConns"`
	} `json:"database"`
	Auth struct {
		Username string   `json:"username"`
		Password string   `json:"password"`
		Port     string   `json:"port"`
		Hosts    []string `json:"hosts"`
		Schemes  string   `json:"schemes"`
	}
	Proxy struct {
		Host string `json:"host"`
		Port string `json:"port"`
		Type string `json:"type"`
	}
	Env       string `json:"env"`
	JwtFields struct {
		Key            string `json:"key"`
		ExpirationTime int    `json:"expirationTime"` //en secondes
	} `json:"jwtFields"`
	UserAPI struct {
		URL      string `json:"url"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"userAPI"`
}

// LoadConfiguration charge le fichier de configuration
func LoadConfiguration(file string) (*conf, error) {

	var config conf
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err == nil {
		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&config)

		Config = &config
	} else {
		logrus.Fatal(err)
	}

	return &config, err
}
