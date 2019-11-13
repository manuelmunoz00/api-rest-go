package models

// Cliente struct
type Cliente struct {
	ID        int    `json:"id"`
	Rut       string `json:"rut"`
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Correo    string `json:"correo"`
}
