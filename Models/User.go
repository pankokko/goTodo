package Models

import (
	"database/sql"
	"fmt"
	"goTodo/db"
	"log"
	"net/http"
)

type User struct {
	Id   int
	Name string
}

func GetUsers() []User {
	db := db.Connect()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	var result []User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			log.Fatalln(err)
		}
		result = append(result, user)
	}
	defer db.Close()
	return result
}

func SaveUser(h *http.Request) int64 {
	result, _ := db.Save(h.FormValue("name"))
	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println(err.Error())
	}
	return id
}

func SelectUser(id int) User {
	db := db.Connect()
	u := User{}

	if err := db.QueryRow("SELECT * FROM users WHERE id=? LIMIT 1", id).Scan(&u.Id, &u.Name);
		err != nil {
		fmt.Printf(err.Error())
	}
	defer db.Close()
	return u
}

func DeleteUser(id int) sql.Result {
	db := db.Connect()

	stmtDelete, err := db.Prepare("DELETE FROM users WHERE id=?")

	if err != nil {
		panic(err.Error())
	}
	result, err := stmtDelete.Exec(id)

	if err != nil {
		panic(err.Error())
	}
	defer stmtDelete.Close()
	defer db.Close()
	return result
}

func UpdateUser(h *http.Request) sql.Result {
	db := db.Connect()
	stmtUpdate, err := db.Prepare("UPDATE users set name = ? WHERE id=?")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(h.FormValue("id"), "idを受け取りました")
	fmt.Println(h.FormValue("name"), "nameを受け取りました")
	result, err := stmtUpdate.Exec(h.FormValue("name"), h.FormValue("id"))
	if err != nil {
		panic(err.Error())
	}
	defer stmtUpdate.Close()
	defer db.Close()
	return result
}
