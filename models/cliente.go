package models

// Cliente struct
type Cliente struct {
	ID        uint   `json:"id"`
	Rut       string `json:"rut"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Correo    string `json:"correo"`
}

// NewCliente función para crear nuevo cliente
func NewCliente(c Cliente) {

}

// GetCliente función para obtener cliente
func GetCliente() {

}
