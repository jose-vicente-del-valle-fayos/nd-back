package controladores

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
	"nd-back/modelos"
	"os"
	"strconv"
	"time"
)

var formEmpezarTiempo time.Time

func Escribeme(c *fiber.Ctx) error {
	/*
		{
			"nombre": "Chevi",
			"correo": "chevielmejor@gmail.com",
			"mensaje": "Hola chevi. Qué tal estás?"
		}
	*/
	c.Response().Header.Set("Access-Control-Allow-Origin", os.Getenv("CORS_DOMINIO_PERMITIDO"))
	var correo modelos.Correo
	if err := c.BodyParser(&correo); err != nil {
		return err
	}
	to, err := strconv.Atoi(os.Getenv("CORREO_TIMEOUT"))
	if err != nil {
		return err
	}
	tiempoTranscurrido := time.Since(formEmpezarTiempo)
	if (correo.Nombre != "") && (correo.Correo != "") && (correo.Mensaje != "") && (tiempoTranscurrido > time.Duration(to)*time.Second) {
		m := gomail.NewMessage()
		m.SetHeader("From", m.FormatAddress(os.Getenv("CORREO_FROM"), "Nuestro Diario") /* email */)
		m.SetHeader("To", os.Getenv("CORREO_TO"))
		m.SetAddressHeader("reply-to", correo.Correo, correo.Nombre)
		m.SetHeader("Subject", "Correo desde Nuestro Diario")
		m.SetBody("text/html", correo.Mensaje)
		port, _ := strconv.Atoi(os.Getenv("CORREO_PORT"))
		d := gomail.NewDialer(os.Getenv("CORREO_SERVER"), port, os.Getenv("CORREO_FROM"), os.Getenv("CORREO_PASS"))
		err := d.DialAndSend(m)
		if err != nil {
			return err
		} else {
			formEmpezarTiempo = time.Now()
			return c.JSON(correo)
		}
	} else {
		return errors.New("error de validación")
	}
}
