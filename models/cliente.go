package models

import (
	"github.com/jinzhu/gorm"
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
func NewCliente(c Cliente) {

}

// GetCliente función para obtener cliente
func GetCliente() {

}
