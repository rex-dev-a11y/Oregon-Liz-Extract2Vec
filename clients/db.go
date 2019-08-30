package clients

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
	"log"
	"os"
)

type dbURLParts struct {
	Host string
	Port int
	User string
	DBName string
	Password string
}

type DBClientConfig struct {
	DBURLParts dbURLParts
	args       interface{}
	driverName string
}

type DBConnection struct {
	DB *gorm.DB
	Error error
	Logger *log.Logger
}

type option func(client *DBConnection) *DBConnection

func(d *DBConnection) Set(options ...option) {
	for _, option := range options {
		d = option(d)
	}
}

func postgresTemplateString(parts dbURLParts) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", parts.Host, parts.Port, parts.User, parts.DBName, parts.Password)
}


func sqlite3TemplateString(parts dbURLParts) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", parts.Host, parts.Port, parts.User, parts.DBName, parts.Password)
}

var PostgresDB = func(dbc *DBConnection) {

	logger := log.New(os.Stdout, "PostgresDB > ", log.LstdFlags)

	viper.SetConfigType("json")

	viper.SetConfigName("db.config")

	viper.AddConfigPath("./local")

	err := viper.ReadInConfig()

	if err != nil {
		logger.Panic(err)
	}

	logger.Println(viper.Get("user"))

	url  := viper.Get("url").(string)

	port := int(viper.Get("port").(float64))

	user := viper.Get("user").(string)

	dbName := viper.Get("db_name").(string)

	password := viper.Get("password").(string)

	parts := dbURLParts{
		url,
		port,
		user,
		dbName,
		password,
	}

	db, err := gorm.Open("postgres", postgresTemplateString(parts))

	if err != nil {
		logger.Panic("Failed to open PostgresDB", err)
	}

	dbc.DB = db
}

var SQLite3 = func(dbc *DBConnection) {
	logger := log.New(os.Stdout, "DBConn > ", log.LstdFlags)

	cwd, err := os.Getwd()
	if err != nil {
		logger.Panicln("cannot find current working dir")
	}
	db, err := gorm.Open("sqlite3", cwd + "/tmp/a2j.sql3")

	if err != nil {
		dbc.Error = err
		logger.Println(err)
		return
	}

	dbc.DB = db
}

func DefaultDBConnection() *DBConnection {

	logger := log.New(os.Stdout, "DBConn > ", log.LstdFlags)

	dbc := &DBConnection{
		Logger: logger,
	}


	if dbc.Error != nil {
		dbc.Logger.Panicln("Failed to initialize default DBConnection.")
	}

	return dbc
}