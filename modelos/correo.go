package modelos

type Correo struct {
	Nombre  string `json:"nombre"`
	Correo  string `json:"correo"`
	Mensaje string `json:"mensaje"`
}
