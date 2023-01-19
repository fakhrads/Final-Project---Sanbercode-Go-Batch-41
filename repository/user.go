package repository

import (
	"database/sql"
	"final-project/structs"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func GetAllUser(db *sql.DB) (err error, results []structs.User) {
	sql := "SELECT * FROM tbl_users"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}

		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Created_at, &user.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, user)

	}
	return

}

func GetAllUserById(db *sql.DB, user structs.User) (results []structs.User, err error) {
	sql := "SELECT * FROM tbl_user WHERE id= $1"

	rows, err := db.Query(sql, user.ID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}

		err = rows.Scan(&user.ID, &user.Created_at, &user.Updated_at)
		if err != nil {
			panic(err)
		}

		results = append(results, user)

	}
	return

}

func InsertUser(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO tbl_users (username, password, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	hashed, _ := hashPassword(user.Password)
	fmt.Println(hashed)
	errs := db.QueryRow(sql, user.Username, hashed, time.Now(), time.Now())

	return errs.Err()
}

func UpdateUser(db *sql.DB, user structs.User) (err error) {
	sql := "UPDATE tbl_users SET username= $2, password= $3, updated_at= $4 WHERE id= $1"
	hashed, _ := hashPassword(user.Password)
	errs := db.QueryRow(sql, user.ID, user.Username, hashed, time.Now())

	return errs.Err()
}

func DeleteUser(db *sql.DB, User structs.User) (err error) {
	sql := "DELETE FROM tbl_users WHERE id= $1"

	errs := db.QueryRow(sql, User.ID)

	return errs.Err()
}
