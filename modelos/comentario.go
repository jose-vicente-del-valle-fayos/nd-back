package modelos

import "time"

type Comentario struct {
	Id         uint   `json:"id" gorm:"unique"`
	IdEnt      uint   `json:"id_ent" gorm:"not null"`
	Usuario    string `json:"usuario" gorm:"type:VARCHAR(50); not null"`
	Correo     string `json:"correo" gorm:"type:VARCHAR(50); not null"`
	Fecha      string `json:"fecha" gorm:"type:VARCHAR(10); not null"`
	Comentario string `json:"comentario" gorm:"not null"`
}

func (comentario *Comentario) ValidarFecha() bool {
	const layoutFechaNumerica string = "2006-01-02"
	_, e := time.Parse(layoutFechaNumerica, comentario.Fecha)
	if e != nil {
		return false
	} else {
		return true
	}
}

func (comentario *Comentario) ValidarIdEnt() bool {
	return comentario.IdEnt != 0
}

func (comentario *Comentario) ValidarUsuario() bool {
	return comentario.Usuario != ""
}

func (comentario *Comentario) ValidarCorreo() bool {
	return comentario.Correo != ""
}

func (comentario *Comentario) ValidarComentario() bool {
	return comentario.Comentario != ""
}
