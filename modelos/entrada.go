package modelos

import (
	"fmt"
	"time"
)

const layoutFechaNumericae string = "2006-01-02"

var dayse = [...]string{
	"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado",
}
var monthse = [...]string{
	"enero", "febrero", "marzo", "abril", "mayo", "junio",
	"julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre",
}

type Entrada struct {
	Id        uint   `json:"id" gorm:"unique"`
	IdUs      uint   `json:"id_us" gorm:"not null"`
	Usuario   string `json:"usuario" gorm:"type:VARCHAR(50); not null"`
	Especial  *bool  `json:"especial" gorm:"not null"`
	Titulo    string `json:"titulo" gorm:"type:VARCHAR(50); not null"`
	Fecha     string `json:"fecha" gorm:"type:VARCHAR(10); not null"`
	Contenido string `json:"contenido" gorm:"not null"`
	// Imagen      string       `json:"imagen" gorm:"null"`
	Comentarios []Comentario `json:"comentarios" gorm:"foreignKey:IdEnt"`
	TotalCom    uint         `json:"total_com" gorm:"-"`
	Visitas     uint         `json:"visitas" gorm:"default:0"`
}

func (entrada *Entrada) FormatearFecha(fecha string) {
	t, e := time.Parse(layoutFechaNumericae, fecha)
	if e != nil {
		panic(e)
	}
	entrada.Fecha = fmt.Sprintf("%s %d, %d",
		/*dayse[t.Weekday()],*/ monthse[t.Month()-1], t.Day(), t.Year())
}

func (entrada *Entrada) CalcularTotalComentarios() {
	entrada.TotalCom = uint(len(entrada.Comentarios))
}
