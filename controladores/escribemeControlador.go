package controladores

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
	"nd-back/modelos"
)

func Escribeme(c *fiber.Ctx) error {
	/*
		{
			"nombre": "Chevi",
			"correo": "chevielmejor@gmail.com",
			"mensaje": "Hola chevi. Qué tal estás?"
		}
	*/
	var correo modelos.Correo
	if err := c.BodyParser(&correo); err != nil {
		return err
	}
	if (correo.Nombre != "") && (correo.Correo != "") && (correo.Mensaje != "") {
		m := gomail.NewMessage()
		m.SetHeader("From", m.FormatAddress("jvdvf@icloud.com", "Nuestro Diario") /* email */)
		m.SetHeader("To", "hola@nuestrodiario.es")
		m.SetAddressHeader("reply-to", correo.Correo, correo.Nombre)
		m.SetHeader("Subject", "Correo desde Nuestro Diario")
		m.SetBody("text/html", correo.Mensaje)
		d := gomail.NewDialer("smtp.mail.me.com", 587, "jvdvf@icloud.com", "twax-oueq-jwbu-ywmr")
		err := d.DialAndSend(m)
		if err != nil {
			return err
		} else {
			return c.JSON(correo)
		}
	} else {
		return errors.New("error de validación")
	}
}
