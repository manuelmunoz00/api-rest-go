package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// DBcon ...
var DBcon *gorm.DB

// Conectar ...
func Conectar() error {
	errenv := godotenv.Load()
	if errenv != nil {
		fmt.Println("Error loading env file")
	}
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	var er error
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	// fmt.Println(dbURL)
	DBcon, er := gorm.Open("mysql", dbURL)
	if er != nil {
		fmt.Println(DBcon)
		panic(er)
		// fmt.Println("error de conexion")

	}
	// defer DBcon.Close()
	return er
}

// Cerrar ...
func Cerrar() error {
	return DBcon.Close()
}
