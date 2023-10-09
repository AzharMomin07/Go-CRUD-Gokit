package main

import (
	config "AdminPortal/internal/config"
	"AdminPortal/internal/dao"
	"AdminPortal/internal/database"
	"AdminPortal/internal/endpoints"
	"AdminPortal/internal/service"
	"AdminPortal/internal/transport"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

func main() {

	logrus.Info("Starting Employee Management Service....")
	config.LoadConfig(config.GetAppEnvLocation())
	//y variables are present
	if err := config.CheckRequiredVariables(); err != nil {
		logrus.Fatalf("Failed to fulfill CheckRequiredVariables Condition :%v", err)
	}

	//setup to Db
	dbConfig := database.DbConfig{
		Host:   config.GetPostgresHost(),
		Port:   config.GetPostgresPort(),
		User:   config.GetPostgresUser(),
		Pass:   config.GetPostgresPass(),
		DbName: config.GetPostgresDb(),
	}

	dbCon, err := database.InitDatabase(dbConfig)

	if err != nil {
		log.Fatal("error occured while connecting to db")
	}
	router := mux.NewRouter()

	orgDao := dao.OrgDao(dbCon)
	orgService := service.NewOrgService(orgDao)
	orgEndpoint := endpoints.MakeOrgEndpoints(orgService)
	transport.CreatesOrgHttpHandler(orgEndpoint, router)
	startServer(router)

}

func startServer(router *mux.Router) {

	serverPort := fmt.Sprintf(":%s", config.GetServerPort())
	logrus.Infof("Starting server on %s.........", serverPort)
	readTimeout := time.Duration(config.GetReadTimeout())
	writeTimeout := time.Duration(config.GetWriteTimeout())

	server := &http.Server{
		Addr:         serverPort,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		Handler:      router,
	}

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatalf("failed to start the server on %s %v", config.GetServerPort(), err)
	}
}
