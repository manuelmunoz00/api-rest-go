package otros

import (
	"github.com/jinzhu/gorm"
)

// Metro struct
type Metro struct {
	gorm.Model
	ID     uint `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Nombre string
	Precio int
}
