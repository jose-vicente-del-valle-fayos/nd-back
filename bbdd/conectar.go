package bbdd

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nd-back/modelos"
	"os"
)

var DB *gorm.DB

func Conectar() {
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic("No se puede conectar a la base de datos.")
	}
	err = db.AutoMigrate(&modelos.Usuario{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(&modelos.Entrada{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(&modelos.Comentario{})
	if err != nil {
		fmt.Println(err)
	}
	DB = db
}
