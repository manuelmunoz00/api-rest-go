package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/manuelmunoz00/api-rest-go/config"
)

// Cliente struct
type Cliente struct {
	gorm.Model
	// ID        uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Rut       string
	Nombres   string
	Apellidos string
	Correo    string
}

// NewCliente función para crear nuevo cliente
func (c *Cliente) NewCliente() bool {
	// fmt.Println(c.Rut)

	// Siguiente codigo antes de comentario revisando funciona bien
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading env file")
	// }
	// username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASS")
	// dbName := os.Getenv("DB_NAME")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")

	// dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	// db, err := gorm.Open("mysql", dbURL)
	// defer db.Close()
	// db.Create(c)
	// return true

	// Revisando
	dbb := config.Conectar()
	// fmt.Println("error: ", err)
	fmt.Println("NILLLLLLLL")
	fmt.Println("fuera de error")
	// defer config.Cerrar()
	// config.DBcon.Create(c)
	dbb.Create(c)
	dbb.Close()
	return true

}

// GetCliente función para obtener cliente
func GetCliente() {

}
