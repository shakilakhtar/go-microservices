package utils


/*
import (
	"bytes"
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"dt-logger/logger"
	"go-microservice-seed-2/types"
	"os"
	"strings"
)

//Method to get local postgres database connection with GORM handle
func GetConnection() *gorm.DB {
	conn, err := gorm.Open(ApplicationContext.Configuration.DBConfig.Database, GetDBURI())
	if err != nil {
		panic("failed to connect database")
	}
	return conn
}

//Close an open gorm database connection
func CloseConnection(db *gorm.DB) {
	db.Close()
}

//Prepare Database connection URI from config details
func GetDBURI() string {
	var buffer bytes.Buffer
	buffer.WriteString("host=")
	buffer.WriteString(ApplicationContext.Configuration.DBConfig.Host)
	buffer.WriteString(" port=")
	buffer.WriteString(ApplicationContext.Configuration.DBConfig.Port)
	buffer.WriteString(" user=")
	buffer.WriteString(ApplicationContext.Configuration.DBConfig.Username)
	buffer.WriteString(" dbname=")
	buffer.WriteString(ApplicationContext.Configuration.DBConfig.Schema)
	buffer.WriteString(" sslmode=")
	buffer.WriteString(ApplicationContext.Configuration.DBConfig.SSLMode)
	buffer.WriteString(" password=")
	buffer.WriteString(ApplicationContext.Configuration.DBConfig.Password)

	logger.Info("Database Connection Uri", "DBURI", buffer.String())

	return buffer.String()
}

//Gets Database URI from Predix if running local will consider local configuration
func GetCFDBConfig(path string) string {
	var dbservice types.CFDBService
	var uri string

	vcapservices := os.Getenv("VCAP_SERVICES")

	if vcapservices != "" {
		json.NewDecoder(strings.NewReader(vcapservices)).Decode(&dbservice)
		credentials := dbservice.Postgres[0].CFCredentials
		uri = credentials.Dsn
	} else {
		uri = GetDBURI()
	}

	return uri
}

//Gets port details from Predix VCAP details
func GetCFPort() string {
	port := os.Getenv("VCAP_APP_PORT")

	if port == "" {
		port = "3030"
	}

	return port
}
*/