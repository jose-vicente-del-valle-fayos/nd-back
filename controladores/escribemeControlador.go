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

var formEmpezarTiempo = make(map[string]time.Time)

func Escribeme(c *fiber.Ctx) error {
	/*
		{
			"nombre": "Chevi",
			"correo": "chevielmejor@gmail.com",
			"mensaje": "Hola chevi. Qué tal estás?"
		}
	*/
	if _, ok := formEmpezarTiempo[c.IP()]; !ok {
		// Si formEmpezarTiempo[c.IP()] no está definido
		formEmpezarTiempo[c.IP()] = time.Now().Add(-1800 * time.Second)
	}
	var correo modelos.Correo
	if err := c.BodyParser(&correo); err != nil {
		return err
	}
	if (correo.Nombre != "") && (correo.Correo != "") && (correo.Mensaje != "") && (time.Since(formEmpezarTiempo[c.IP()]) > 1800*time.Second) {
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
			formEmpezarTiempo[c.IP()] = time.Now()
			return c.JSON(correo)
		}
	} else {
		return errors.New("error de validación")
	}
}
