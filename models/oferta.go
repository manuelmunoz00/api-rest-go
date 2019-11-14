package models

import (
	"github.com/jinzhu/gorm"
	otro "github.com/manuelmunoz00/api-rest-go/otros"
)

// Oferta struct
type Oferta struct {
	gorm.Model
	ID            uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Nombre        string
	Precio        int
	ListadoMetros []otro.Metro
}
