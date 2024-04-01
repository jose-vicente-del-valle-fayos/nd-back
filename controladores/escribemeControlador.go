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

var envioPermitido = true
var tiemposLlamada = make(map[int]time.Time)
var mu sync.Mutex

func getEnvMaxLlam(key string) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		fmt.Println(err)
	}
	return val
}

func getEnvDuracion(key string) time.Duration {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		fmt.Println(err)
	}
	return time.Duration(val) * time.Second
}

func ComprobarLlamada(tramo int, maxLlamadas int, intervalo time.Duration) bool {
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
	case 1, 2, 3, 4, 5:
		envioPermitido = false
	default:
		envioPermitido = true
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
		t, e := strconv.Atoi(os.Getenv("CORREO_TIMEOUT_TRAMO_5"))
		if e != nil {
			fmt.Println(e)
		}
		for i := 0; i < (t / 10); i++ {
			time.Sleep(10 * time.Second)
			tramos := []struct {
				maxLlamadas int
				timeout     time.Duration
			}{
				{getEnvMaxLlam("CORREO_MAX_LLAMADAS_TRAMO_1"), getEnvDuracion("CORREO_TIMEOUT_TRAMO_1")},
				{getEnvMaxLlam("CORREO_MAX_LLAMADAS_TRAMO_2"), getEnvDuracion("CORREO_TIMEOUT_TRAMO_2")},
				{getEnvMaxLlam("CORREO_MAX_LLAMADAS_TRAMO_3"), getEnvDuracion("CORREO_TIMEOUT_TRAMO_3")},
				{getEnvMaxLlam("CORREO_MAX_LLAMADAS_TRAMO_4"), getEnvDuracion("CORREO_TIMEOUT_TRAMO_4")},
				{getEnvMaxLlam("CORREO_MAX_LLAMADAS_TRAMO_5"), getEnvDuracion("CORREO_TIMEOUT_TRAMO_5")},
			}

			for i, tramo := range tramos {
				if ComprobarLlamada(i+1, tramo.maxLlamadas, tramo.timeout) {
					CambiarEstado(i + 1)
				}
			}
		}
	}()

	var correo modelos.Correo
	if err := c.BodyParser(&correo); err != nil {
		return err
	}
	if (correo.Nombre != "") && (correo.Correo != "") && (correo.Mensaje != "") && envioPermitido {
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
