package conn

import (
	"database/sql"
	"errors"
	"auth-api-crossfitlov/parameters"
	"time"

	// Driver mysql
	_ "github.com/go-sql-driver/mysql"
)

// GetConn obtenir de connecteur bdd
func GetConn() (*sql.DB, error) {

	c := parameters.Config

	if c != nil {
		address := c.Database.Username + ":" + c.Database.Password +
			"@tcp(" + c.Database.Host + ":" + c.Database.Port + ")/" + c.Database.Name + "?charset=utf8"

		db, err := sql.Open(c.Database.Adapter, address)

		if err == nil {
			db.SetMaxOpenConns(c.Database.MaxOpenConns)
			db.SetMaxIdleConns(c.Database.MaxIdleConns)
			db.SetConnMaxLifetime(time.Second * 60)
		}

		if err == nil {
			err := db.Ping()
			return db, err
		}

	}

	return nil, errors.New("Access denied")

}
