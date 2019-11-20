package models

import (
	"github.com/jinzhu/gorm"
)

// Servicio struct
type Servicio struct {
	gorm.Model
	// ID     uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Nombre  string
	Precio  int
	Ofertas []*Oferta `gorm:"many2many:oferta_servicios;"`
}
