package controladores

import (
	"github.com/gofiber/fiber/v2"
	"nd-back/bbdd"
	"nd-back/modelos"
	"os"
	"strconv"
)

func TodosComentarios(c *fiber.Ctx) error {
	c.Response().Header.Set("Access-Control-Allow-Origin", os.Getenv("CORS_DOMINIO_PERMITIDO"))
	var comentarios []modelos.Comentario
	bbdd.DB.Find(&comentarios)
	var total int64
	bbdd.DB.Model(&comentarios).Count(&total)
	return c.JSON(fiber.Map{
		"datos": comentarios,
		"meta": fiber.Map{
			"total": total,
		},
	})
}

func CrearComentario(c *fiber.Ctx) error {
	/*
		{
			"id_ent": 	1,
			"usuario":    "Sara",
			"correo":     "sara@gmail.com",
			"fecha":      "2018-04-02",
			"comentario": "Hola Chevi. Eres el mejor."
		}
	*/
	c.Response().Header.Set("Access-Control-Allow-Origin", os.Getenv("CORS_DOMINIO_PERMITIDO"))
	var comentario modelos.Comentario
	if err := c.BodyParser(&comentario); err != nil {
		return err
	}
	if comentario.ValidarFecha() && comentario.ValidarIdEnt() && comentario.ValidarUsuario() && comentario.ValidarCorreo() && comentario.ValidarComentario() {
		bbdd.DB.Create(&comentario)
		return c.JSON(comentario)
	}
	return c.JSON(fiber.Map{"mensaje": "error de validación"})
}

func LeerComentario(c *fiber.Ctx) error {
	c.Response().Header.Set("Access-Control-Allow-Origin", os.Getenv("CORS_DOMINIO_PERMITIDO"))
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	comentario := modelos.Comentario{
		Id: uint(id),
	}
	bbdd.DB.Find(&comentario)
	return c.JSON(comentario)
}

func ActualizarComentario(c *fiber.Ctx) error {
	/*
		{
			"id_ent": 	1,
			"usuario":    "Sara actualizada",
			"correo":     "saractualizada@gmail.com",
			"fecha":      "2018-04-03",
			"comentario": "Hola Chevi. Eres el mejor del mundo mundial."
		}
	*/
	c.Response().Header.Set("Access-Control-Allow-Origin", os.Getenv("CORS_DOMINIO_PERMITIDO"))
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	comentario := modelos.Comentario{
		Id: uint(id),
	}
	if err := c.BodyParser(&comentario); err != nil {
		return err
	}
	if comentario.ValidarFecha() && comentario.ValidarIdEnt() && comentario.ValidarUsuario() && comentario.ValidarCorreo() && comentario.ValidarComentario() {
		bbdd.DB.Model(&comentario).Updates(comentario)
		return c.JSON(comentario)
	}
	return c.JSON(fiber.Map{"mensaje": "error de validación"})
}

func BorrarComentario(c *fiber.Ctx) error {
	c.Response().Header.Set("Access-Control-Allow-Origin", os.Getenv("CORS_DOMINIO_PERMITIDO"))
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	comentario := modelos.Comentario{
		Id: uint(id),
	}
	bbdd.DB.Delete(&comentario)
	return nil
}
