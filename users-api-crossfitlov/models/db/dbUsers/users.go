package dbUsers

import (
	"database/sql"
	"users-api-crossfitlov/models/structs/in"
)

func InsertUser(db *sql.DB, user in.UserInfos) error {
	sql := "INSERT INTO dataCL" +
		"(crossfitlovID," +
		"firstname," +
		"gender," +
		"age," +
		"boxCity," +
		"email," +
		"biography," +
		"job," +
		"created_at) " +
		"VALUES (?,?,?,?,?,?,?,?,?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.CrossfitlovID,
		user.Firstname,
		user.Gender,
		user.Age,
		user.BoxCity,
		user.Email,
		user.Biography,
		user.Job,
		user.CreatedAt)

	if err != nil {
		return err
	}

	return nil

}

func UpdateUser(db *sql.DB, user in.UserInfos) (bool, error) {

	sql := "UPDATE dataCL SET " +
		"boxCity = ?, " +
		"email = ?, " +
		"biography = ?, " +
		"job = ? " +
		"WHERE crossfitlovID = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, err
	}

	res, err := stmt.Exec(user.BoxCity,
		user.Email,
		user.Biography,
		user.Job,
		user.CrossfitlovID)

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

func SelectUser(db *sql.DB, ID string) (userInfos in.UserInfos, userExist bool, err error) {

	sqlStatement := `SELECT crossfitlovID, firstname, gender, age, boxCity, email, biography, job, created_at FROM dataCL WHERE crossfitlovID=?`

	row := db.QueryRow(sqlStatement, ID)
	switch err := row.Scan(&userInfos.CrossfitlovID,
		&userInfos.Firstname,
		&userInfos.Gender,
		&userInfos.Age,
		&userInfos.BoxCity,
		&userInfos.Email,
		&userInfos.Biography,
		&userInfos.Job,
		&userInfos.CreatedAt); err {
	case sql.ErrNoRows:
		return userInfos, false, nil
	case nil:
		return userInfos, true, nil
	default:
		return userInfos, false, err

	}
}

func DeleteUser(db *sql.DB, ID string) (userExist bool, err error) {

	sqlStatement := `DELETE FROM dataCL WHERE crossfitlovID=?`

	res, err := db.Exec(sqlStatement, ID)
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
