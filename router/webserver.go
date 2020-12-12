package router

import (
	"goTodo/controllers"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
)

type WebServer struct {
	Port string
}

var server WebServer

func init() {
	config, _ := ini.Load("config.ini")
	server = WebServer{
		Port: config.Section("web").Key("port").MustString(":8000"),
	}
}

func StartWebServer()  {
	http.HandleFunc("/save/", controllers.SaveHandler)
	http.HandleFunc("/", controllers.MainHandler)
	http.HandleFunc("/data/", controllers.DataHandler)
	http.HandleFunc("/show/", controllers.ShowHandler)
	http.HandleFunc("/detail", controllers.DetailHandler)
	http.HandleFunc("/edit", controllers.DetailHandler)
	http.HandleFunc("/confirm/", controllers.ConfirmHandler)
	http.HandleFunc("/delete", controllers.DeleteHandler)
	http.HandleFunc("/update", controllers.UpdateHandler)
	log.Fatalln(http.ListenAndServe(server.Port, nil))
}