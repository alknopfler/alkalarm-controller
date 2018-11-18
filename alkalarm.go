package main

import (
	_ "github.com/mattn/go-sqlite3"
	"os"
	"github.com/alknopfler/alkalarm-controller/database"
	cfg "github.com/alknopfler/alkalarm-controller/config"
	"github.com/alknopfler/alkalarm-controller/api"
	"github.com/gorilla/handlers"
	"net/http"
	"log/syslog"
	"log"
	"github.com/gorilla/mux"
)


func init(){
	needCreation:=true
 	logging, e := syslog.New(syslog.LOG_NOTICE, "alkalarm")
	if e == nil {
		log.SetOutput(logging)
	}

	log.Print("Validating the database, and other params...Could take some minutes...")
	//First Time to execute needs create database and scheme
	if _, err := os.Stat(cfg.DB_NAME); err == nil {
		log.Println("ya existe la base de datos")
		needCreation= false
	}
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB First Time (init)")
		os.Exit(2)
	}
	defer db.Close()
	err=database.CreateSchemas(db)
	if err!=nil{
		log.Println("Error creating the schema the first time (init)")
		os.Exit(3)
	}
	if needCreation {
		err=database.Operate(db,cfg.GLOBAL_STATE_INSERT,cfg.STATE_INAC)
		if err!=nil{
			log.Println("Error activating the first time (init)")
			os.Exit(3)
		}
	}
	log.Println("Success...Starting the program")
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/activate/full",api.HandlerActivateFull).Methods("POST")
	r.HandleFunc("/activate/partial",api.HandlerActivatePartial).Methods("POST")
	r.HandleFunc("/deactivate",api.HandlerDeactivate).Methods("POST")
	r.HandleFunc("/status",api.HandlerAlarmStatus).Methods("GET")

	corsObj:=handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS"})
	err := http.ListenAndServe(cfg.SERVER_API_PORT, handlers.CORS(corsObj,headersOk,methodsOk)(r))
	if err != nil {
		log.Println("Error listening api server...")
	}
}
