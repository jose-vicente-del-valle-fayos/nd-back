package modelos

const layoutFechaNumericac string = "2006-01-02"

var daysc = [...]string{
	"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado",
}
var monthsc = [...]string{
	"enero", "febrero", "marzo", "abril", "mayo", "junio",
	"julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre",
}

type Comentario struct {
	Id         uint   `json:"id" gorm:"unique"`
	IdEnt      uint   `json:"id_ent" gorm:"not null"`
	Usuario    string `json:"usuario" gorm:"type:VARCHAR(50); not null"`
	Correo     string `json:"correo" gorm:"type:VARCHAR(50); not null"`
	Fecha      string `json:"fecha" gorm:"type:VARCHAR(10); not null"`
	Comentario string `json:"comentario" gorm:"not null"`
}
