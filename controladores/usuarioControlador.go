package controladores

import (
	"github.com/gofiber/fiber/v2"
	"nd-back/bbdd"
	"nd-back/modelos"
	"strconv"
)

func TodosUsuarios(c *fiber.Ctx) error {
	var usuarios []modelos.Usuario
	bbdd.DB.Preload("Entradas").Preload("Entradas.Comentarios").Find(&usuarios)
	return c.JSON(usuarios)
}

func CrearUsuario(c *fiber.Ctx) error {
	/*
		{
			"sobrenombre": "Chevi",
			"nombre":      "José Vicente",
			"apellidos":   "del Valle Fayos",
			"correo":      "hola@chevi.soy"
		}
	*/
	var usuario modelos.Usuario
	if err := c.BodyParser(&usuario); err != nil {
		return err
	}
	usuario.PonContrasena("1234")
	bbdd.DB.Create(&usuario)
	return c.JSON(usuario)
}

func LeerUsuario(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	usuario := modelos.Usuario{
		Id: uint(id),
	}
	bbdd.DB.Preload("Entradas").Find(&usuario)
	usuario.CalcularTotalEntradas()
	return c.JSON(usuario)
}

func ActualizarUsuario(c *fiber.Ctx) error {
	/*
		{
			"sobrenombre": "Ramón",
			"nombre":      "Supermones rascajeribrother",
			"apellidos":   "Bronson",
			"correo":      "pato@gmail.com"
		}
	*/
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	usuario := modelos.Usuario{
		Id: uint(id),
	}
	if err := c.BodyParser(&usuario); err != nil {
		return err
	}
	bbdd.DB.Model(&usuario).Updates(usuario)
	return c.JSON(usuario)
}

/*
func BorrarUsuario(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	usuario := modelos.Usuario{
		Id: uint(id),
	}
	bbdd.DB.Delete(&usuario)
	return nil
}
*/
