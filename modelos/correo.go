package modelos

type Correo struct {
	Nombre  string `json:"nombre"`
	Correo  string `json:"correo"`
	Mensaje string `json:"mensaje"`
}

func (correo *Correo) ValidarNombre() bool { return correo.Nombre != "" }

func (correo *Correo) ValidarCorreo() bool { return correo.Correo != "" }

func (correo *Correo) ValidarMensaje() bool { return correo.Mensaje != "" }
