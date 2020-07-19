package utils

import (
	//"encoding/json"
	//"dt-logger/logger"
	//"io/ioutil"
	//"sync"
	//"time"
	//"net/http"
	kitutils "go-microservices-platform/utils"
)
/*
type (
	configuration struct {
		AppConfig appConfiguration
		DBConfig  dbConfiguration
	}

	applicationContext struct {
		Configuration configuration
		User          string
	}

	appConfiguration struct {
		ApiPort      string `json:"api_port"`
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		UaaHost      string `json:"uaa_host"`
		AuthEnabled  bool   `json:"auth_enabled"`
	}

	dbConfiguration struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Database string `json:"database"`
		Schema   string `json:"schema"`
		Username string `json:"username"`
		Password string `json:"password"`
		SSLMode  string `json:"sslmode"`
	}
)

const (
	DB_CONFIG  = "config/dbconfig.json"
	APP_CONFIG = "config/appconfig.json"
)
*/
/*
//A singleton context
var ApplicationContext *applicationContext
var once sync.Once
*/
// Initialize AppConfig
func Init() {
	//loadResources()
	kitutils.LoadDBConfigurationFromFile("config")
}
/*
// Load application configurations from config files and decode into context
func loadResources() {
	GetApplicationContext()
	//load database configurations
	loadConfig(DB_CONFIG, &ApplicationContext.Configuration.DBConfig)
	//load application configurations
	loadConfig(APP_CONFIG, &ApplicationContext.Configuration.AppConfig)
}

func loadConfig(fileName string, conf interface{}) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		logger.Fatal(0, "[loadConfig]: %s\n", err)
	}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		logger.Fatal(0, "[loadConfig]: %s\n", err)
	}
}

// GetInstance returns a singleton instance of application context.
func GetApplicationContext() *applicationContext {
	if ApplicationContext == nil {
		once.Do(func() {
			ApplicationContext = &applicationContext{}
		})
	}
	return ApplicationContext

}

//Logs HTTP request with it's parameters i.e HTTP Method,Request URI, Handler Name and time.
func LogHttp(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		logger.Info(
			"HttpMethod,URI, HandlerMethod, Start Time",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
*/