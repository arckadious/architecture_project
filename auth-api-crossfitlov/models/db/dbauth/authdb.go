package dbauth

import (
	"database/sql"
)

func GetPasswdAndID(login string, db *sql.DB) (ID, encryptedPassword string, userExist bool, err error) {

	sqlStatement := `SELECT crossfitlovID, encrypted_passwd FROM dataCL WHERE user_login=?`

	row := db.QueryRow(sqlStatement, login)
	switch err := row.Scan(&ID, &encryptedPassword); err {
	case sql.ErrNoRows:
		return "", "", false, nil
	case nil:
		return ID, encryptedPassword, true, nil
	default:
		return "", "", false, err
	}
}

func GetLastInsertID(db *sql.DB) (ID int, err error) {

	sqlStatement := `SELECT LAST_INSERT_ID()`

	err = db.QueryRow(sqlStatement).Scan(&ID)
	if err != nil {
		return 0, err
	}

	return ID, nil
}

func CreatePasswdAndID(db *sql.DB, login, encryptedPassword string) error {
	sql := "INSERT INTO dataCL" +
		"(user_login," +
		"encrypted_passwd) " +
		"VALUES (?,?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(login,
		encryptedPassword)

	if err != nil {
		return err
	}

	return nil

}

func DeletePasswdAndID(login string, db *sql.DB) (userExist bool, err error) {

	sqlStatement := `DELETE FROM dataCL WHERE user_login=?`

	res, err := db.Exec(sqlStatement, login)
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil

}
