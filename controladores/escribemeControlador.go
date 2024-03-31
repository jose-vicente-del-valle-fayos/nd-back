package controladores

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
	"nd-back/modelos"
	"os"
	"strconv"
	"sync"
	"time"
)

var variableEstado = true
var tiemposLlamada = make(map[int]time.Time)
var mu sync.Mutex

func ComprobarLlamadas(tramo int, maxLlamadas int, intervalo time.Duration) bool {
	mu.Lock()
	defer mu.Unlock()
	now := time.Now()
	for k, v := range tiemposLlamada {
		if now.Sub(v) > intervalo {
			delete(tiemposLlamada, k)
		}
	}
	tiemposLlamada[tramo] = now
	return len(tiemposLlamada) > maxLlamadas
}

func CambiarEstado(tramo int) {
	switch tramo {
	case 1:
		variableEstado = false
	case 2:
		variableEstado = false
	case 3:
		variableEstado = false
	case 4:
		variableEstado = false
	case 5:
		variableEstado = false
	default:
		variableEstado = true
	}
}

func Escribeme(c *fiber.Ctx) error {
	/*
		{
			"nombre": "Chevi",
			"correo": "chevielmejor@gmail.com",
			"mensaje": "Hola chevi. Qué tal estás?"
		}
	*/

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(10 * time.Second)
			mllamadas, em := strconv.Atoi(os.Getenv("CORREO_MAX_LLAMADAS_POR_TRAMO"))
			if em != nil {
				fmt.Println(em)
			}
			t1, e1 := strconv.Atoi(os.Getenv("CORREO_TIMEOUT_TRAMO_1"))
			if e1 != nil {
				fmt.Println(e1)
			}
			t2, e2 := strconv.Atoi(os.Getenv("CORREO_TIMEOUT_TRAMO_2"))
			if e2 != nil {
				fmt.Println(e2)
			}
			t3, e3 := strconv.Atoi(os.Getenv("CORREO_TIMEOUT_TRAMO_3"))
			if e3 != nil {
				fmt.Println(e3)
			}
			t4, e4 := strconv.Atoi(os.Getenv("CORREO_TIMEOUT_TRAMO_4"))
			if e4 != nil {
				fmt.Println(e4)
			}
			t5, e5 := strconv.Atoi(os.Getenv("CORREO_TIMEOUT_TRAMO_5"))
			if e5 != nil {
				fmt.Println(e5)
			}
			if ComprobarLlamadas(1, mllamadas, time.Duration(t1)*time.Second) {
				CambiarEstado(1)
			}
			if ComprobarLlamadas(2, mllamadas, time.Duration(t2)*time.Second) {
				CambiarEstado(2)
			}
			if ComprobarLlamadas(3, mllamadas, time.Duration(t3)*time.Second) {
				CambiarEstado(3)
			}
			if ComprobarLlamadas(4, mllamadas, time.Duration(t4)*time.Second) {
				CambiarEstado(4)
			}
			if ComprobarLlamadas(5, mllamadas, time.Duration(t5)*time.Second) {
				CambiarEstado(5)
			}
		}
	}()

	var correo modelos.Correo
	if err := c.BodyParser(&correo); err != nil {
		return err
	}
	if (correo.Nombre != "") && (correo.Correo != "") && (correo.Mensaje != "") && (variableEstado) {
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
			return c.JSON(correo)
		}
	} else {
		return errors.New("error de validación")
	}
}
