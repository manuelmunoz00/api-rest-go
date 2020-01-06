package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// Register some standard stuff
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DBcon ...
var DBcon *gorm.DB

// Conectar ...
func Conectar() *gorm.DB {
	errenv := godotenv.Load()
	if errenv != nil {
		fmt.Println("Error loading env file")
	}
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// var er error
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", username, password, dbHost, dbPort, dbName)
	fmt.Println(dbURL)
	DBcon, errenv := gorm.Open("mysql", dbURL)
	// defer DBcon.Close()
	// fmt.Println(username)
	// fmt.Println(password)
	// fmt.Println(dbHost)
	// fmt.Println(dbPort)
	// fmt.Println(dbName)
	// fmt.Println(DBcon)
	// fmt.Println(errenv)
	// fmt.Println("fin conexion")

	if errenv != nil {
		fmt.Println("error de conexion")
		fmt.Println(DBcon)
		panic(errenv)
	}

	// for {
	// 	pingErr := DBcon.DB().Ping()
	// 	if pingErr != nil {
	// 		fmt.Println(pingErr)
	// 	} else {
	// 		fmt.Println("Connected")
	// 	}
	// 	time.Sleep(time.Duration(3) * time.Second)
	// }

	// defer DBcon.Close()
	return DBcon
}

// Cerrar ...
func Cerrar() error {
	return DBcon.Close()
}
