package controllers

import (
	"fmt"
	"goTodo/Models"
	"goTodo/db"
	"html/template"
	"net/http"
	"strconv"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	users := Models.GetUsers()
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, users); err != nil {
		panic(err.Error())
	}
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	var temperature []Models.TemperatureDataElem
	temperature = append(temperature, Models.TemperatureDataElem{
		Label: "沖縄県",
		Data:  []float64{17.0, 17.1, 18.9, 21.4, 24.0, 26.8, 28.9, 28.7, 27.6, 25.2, 22.1, 18.7},
	})
	temperature = append(temperature, Models.TemperatureDataElem{
		Label: "東京都",
		Data:  []float64{5.2, 5.7, 8.7, 13.9, 18.2, 21.4, 25.0, 26.4, 22.8, 17.5, 12.1, 7.6},
	})
	t, err := template.ParseFiles("data.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, temperature); err != nil {
		panic(err.Error())
	}
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("show.html")

	if err != nil {
		panic(err.Error())
	}

	if err := t.Execute(w, nil); err != nil {
		err.Error()
	}
}

func DetailHandler(w http.ResponseWriter, r *http.Request) {
	query :=r.URL.Query().Get("id")
	i, err := strconv.Atoi(query)
	fmt.Printf("%T\n", i) // int
	if err != nil {
		panic(err.Error())
	}
	result := Models.SelectUser(i)

	if r.URL.Path == "/detail" {
		t, err := template.ParseFiles("detail.html")
		if err != nil {
			panic(err.Error())
		}
		t.Execute(w, result)
	} else if r.URL.Path == "/edit" {
		t, err := template.ParseFiles("edit.html")
		if err != nil {
			panic(err.Error())
		}
		t.Execute(w, result)
	}

}

func ConfirmHandler(w http.ResponseWriter, r *http.Request) {

	values := map[string]string{
		"account": r.FormValue("account"),
		"name":    r.FormValue("name"),
		"passwd":  r.FormValue("passwd"),
	}

	t, err := template.ParseFiles("confirm.html")

	if err != nil {
		panic(err.Error())
	}

	if err := t.Execute(w, values); err != nil {
		panic(err.Error())
	}

}

func SaveHandler(w http.ResponseWriter, h *http.Request) {
	println("saving...............")
	id := Models.SaveUser(h)
	db.Select(id)
	http.Redirect(w, h, "/", 301)
}

func DeleteHandler(w http.ResponseWriter, h *http.Request) {
	value := h.URL.Query()
	i, err := strconv.Atoi(value["id"][0])
	fmt.Printf("%T\n", i) // int
	if err != nil {
		panic(err.Error())
	}
	result := Models.DeleteUser(i)

	if result == nil {
		fmt.Printf("削除に失敗しました")
	}

	http.Redirect(w, h, "/", 301)
}

func UpdateHandler(w http.ResponseWriter, h *http.Request) {
	fmt.Println(h.FormValue("id"), "idをチェックします")
	result := Models.UpdateUser(h)
	if result == nil {
		fmt.Printf("更新に成功しました")
	}
	http.Redirect(w, h, "/detail?id="+h.FormValue("id"), 301)
}
