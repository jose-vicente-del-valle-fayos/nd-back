package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"nd-back/bbdd"
	"nd-back/rutas"
)

func main() {
	/*
		err = db.AutoMigrate(&Usuario{})
		if err != nil {
			fmt.Println(err)
		}
		err = db.AutoMigrate(&Comentario{})
		if err != nil {
			fmt.Println(err)
		}
		err = db.AutoMigrate(&Entrada{})
		if err != nil {
			fmt.Println(err)
		}
	*/

	// db.Updates(&usuario)
	// db.Delete(&usuario)

	/*
		comentarios := []Comentario{
			{
				Usuario:    "Sara",
				Correo:     "sara@gmail.com",
				Fecha:      "2018-04-02",
				Comentario: "Hola Chevi. Eres el mejor.",
			},
			{
				Usuario:    "Susana",
				Correo:     "susana@gmail.com",
				Fecha:      "2019-05-01",
				Comentario: "Hola Chevi. Te quiero.",
			},
			{
				Usuario:    "Reme",
				Correo:     "reme@gmail.com",
				Fecha:      "2021-01-28",
				Comentario: "Hola Chevi. Eres mi osito.",
			},
		}

		usuario := Usuario{
			Sobrenombre: "Chevi",
			Nombre:      "José Vicente",
			Apellidos:   "del Valle Fayos",
			Correo:      "hola@chevi.soy",
		}

		entrada := Entrada{
			Usuario:     usuario,
			Especial:    false,
			Titulo:      "Este es un título fantástico",
			Fecha:       "2024-02-19",
			Contenido:   "Hola pato.",
			Imagen:      "https://www.google.com/url?sa=i&url=https%3A%2F%2Fes.wikipedia.org%2Fwiki%2FAnas_platyrhynchos_domesticus&psig=AOvVaw16M0C4Pe7qEKlJtoZDDoS7&ust=1708445611787000&source=images&cd=vfe&opi=89978449&ved=0CBAQjRxqFwoTCJionIDmt4QDFQAAAAAdAAAAABAF",
			Comentarios: comentarios,
		}
		db.Delete(&entrada)
		db.Create(&entrada)

		e := Entrada{}
		db.Preload("Comentarios").Preload("Usuario").Where("especial", false).Find(&e)
		fmt.Println(e)
	*/
	bbdd.Conectar()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://www.nuestrodiario.es, http://localhost:3000, http://localhost:8000,",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	rutas.Configuracion(app)
	app.Listen(":8000")

}
