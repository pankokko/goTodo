package main

import (
	_ "github.com/go-sql-driver/mysql"
	"goTodo/controllers"
)

////TemperatureDataElem 気温データの一つのデータセット
//type TemperatureDataElem struct {
//	Label string
//	Data  []float64
//}
//
//func mainHandler(w http.ResponseWriter, r *http.Request) {
//	t, err := template.ParseFiles("index.html")
//	if err != nil {
//		panic(err.Error())
//	}
//	if err := t.Execute(w, nil); err != nil {
//		panic(err.Error())
//	}
//}
//
//func dataHandler(w http.ResponseWriter, r *http.Request) {
//	var temperatureData []TemperatureDataElem
//	temperatureData = append(temperatureData, TemperatureDataElem{
//		Label: "沖縄県",
//		Data:  []float64{17.0, 17.1, 18.9, 21.4, 24.0, 26.8, 28.9, 28.7, 27.6, 25.2, 22.1, 18.7},
//	})
//	temperatureData = append(temperatureData, TemperatureDataElem{
//		Label: "東京都",
//		Data:  []float64{5.2, 5.7, 8.7, 13.9, 18.2, 21.4, 25.0, 26.4, 22.8, 17.5, 12.1, 7.6},
//	})
//	t, err := template.ParseFiles("data.html")
//	if err != nil {
//		panic(err.Error())
//	}
//	if err := t.Execute(w, temperatureData); err != nil {
//		panic(err.Error())
//	}
//}
//
//func showHandler(w http.ResponseWriter, r *http.Request) {
//	t, err := template.ParseFiles("show.html")
//
//	if err != nil {
//		panic(err.Error())
//	}
//
//	if err := t.Execute(w, nil); err != nil {
//		err.Error()
//	}
//
//}
//
//func confirmHandler(w http.ResponseWriter, r *http.Request) {
//
//	values := map[string]string{
//		"account": r.FormValue("account"),
//		"name":    r.FormValue("name"),
//		"passwd":  r.FormValue("passwd"),
//	}
//
//	t, err := template.ParseFiles("confirm.html")
//
//	if err != nil {
//		panic(err.Error())
//	}
//
//	if err := t.Execute(w, values); err != nil {
//		panic(err.Error())
//	}
//
//}
//
//func saveHandler(w http.ResponseWriter, h *http.Request) {
//	println("saving...............")
//	result, _ := db.Save(h.FormValue("name"))
//	id, err := result.LastInsertId()
//	if err == nil {
//		db.Select(id)
//
//		t, _ := template.ParseFiles("index.html")
//		t.Execute(w, nil)
//	}
//
//}

func main() {
	//http.HandleFunc("/save/", saveHandler)
	//http.HandleFunc("/", mainHandler)
	//http.HandleFunc("/data/", dataHandler)
	//http.HandleFunc("/show/", showHandler)
	//http.HandleFunc("/confirm/", confirmHandler)
	//log.Fatalln(http.ListenAndServe(":8000", nil))
	controllers.StartWebServer()
}
