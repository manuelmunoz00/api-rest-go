package models

import (
	"github.com/jinzhu/gorm"
)

// Oferta struct
type Oferta struct {
	gorm.Model
	ID        uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Nombre    string
	Precio    int
	Servicios []Servicio
}
