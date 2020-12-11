package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

type Person struct {
	ID   int
	Name string
}

type ConfigList struct {
	DbName string
	Driver string
	DbUser string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Driver: cfg.Section("db").Key("driver").MustString("example.sql"),
		DbUser: cfg.Section("db").Key("user").MustString("root@/gosample"),
	}
}

//
//func Connect(){
//	db, err := sql.Open(Config.Driver, Config.DbUser)
//	log.Println("Connected to mysql.")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//データベースへクエリを送信。引っ張ってきたデータがrowsに入る。
//	rows, err := db.Query("SELECT * FROM users")
//	if err != nil {
//		panic(err.Error())
//	}
//
//	for rows.Next() {
//		var person Person //構造体Person型の変数personを定義
//		err := rows.Scan(&person.ID, &person.Name)
//
//		if err != nil {
//			panic(err.Error())
//		}
//		fmt.Println(person.ID, person.Name) //結果　1 yamada 2 suzuki
//	}
//	defer rows.Close()
//
//}

func Save(name string) (sql.Result, error) {
	db, err := sql.Open(Config.Driver, Config.DbUser)

	if err != nil {
		println(err)
	}

	stmt, err := db.Prepare("INSERT INTO users SET name=?")

	if err != nil {
		println(err)
	}

	result, err := stmt.Exec(name)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf( "%v",)

	defer db.Close()
	return result , nil
}

func Select(id int64)  {
	db, err := sql.Open(Config.Driver, Config.DbUser)
	p := Person{}
	if err != nil {
		fmt.Println(err)
	}

	if err := db.QueryRow("SELECT * FROM users WHERE id=? LIMIT 1", id).Scan(&p.ID,&p.Name);
	err != nil {
		fmt.Printf(err.Error())
	}


	fmt.Println(p)

}