package controladores

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"nd-back/bbdd"
	"nd-back/modelos"
	"nd-back/utilidades"
	"strconv"
	"time"
)

func Registrar(c *fiber.Ctx) error {
	/*
		{
			"sobrenombre": "Chevi",
			"nombre":      "José Vicente",
			"apellidos":   "del Valle Fayos",
			"correo":      "hola@chevi.soy",
			"contrasena":  "almasera",
		    "contrasenaconf": "almasera"
		}
	*/
	var datos map[string]string
	if err := c.BodyParser(&datos); err != nil {
		return err
	}
	if datos["contrasena"] != datos["contrasenaconf"] {
		c.Status(400)
		return c.JSON(fiber.Map{"mensaje": "las contraseñas no coinciden"})
	}
	usuario := modelos.Usuario{
		Sobrenombre: datos["sobrenombre"],
		Nombre:      datos["nombre"],
		Apellidos:   datos["apellidos"],
		Correo:      datos["correo"],
		Contrasena:  []byte(""),
	}
	usuario.PonContrasena(datos["contrasena"])
	bbdd.DB.Create(&usuario)
	return c.JSON(usuario)
}

func Ingresar(c *fiber.Ctx) error {
	/*
		{
			"correo":      "hola@chevi.soy",
			"contrasena":  "1234"
		}
	*/
	var datos map[string]string
	if err := c.BodyParser(&datos); err != nil {
		return err
	}
	usuario := modelos.Usuario{}
	bbdd.DB.Where("correo", datos["correo"]).First(&usuario)
	if usuario.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{"mensaje": "usuario no encontrado"})
	}
	if err := usuario.ComparaContrasenas(datos["contrasena"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{"mensaje": "el nombre de usuario o la contraseña no son correctos"})
	}
	token, err := utilidades.GenerarJWT(strconv.Itoa(int(usuario.Id)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	cookie := fiber.Cookie{
		Name:     "nd-jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 12), // Expira en medio día
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"mensaje": "autenticado",
	})
}

func Usuario(c *fiber.Ctx) error {
	// cookie := c.Cookies("nd-jwt")
	cookie := fiber.Cookie{
		Name:     "nd-jwt",
		HTTPOnly: true,  // La cookie solo es accesible mediante HTTP(S)
		Secure:   true,  // Solo se enviará si la conexión es segura (HTTPS)
		SameSite: "Lax", // Configuración de SameSite (puede ser "None", "Lax", o "Strict")
	}
	// Configurar la cookie en la respuesta
	c.Cookie(&cookie)
	id, err := utilidades.ParsearJWT(cookie.Value)
	if err != nil {
		fmt.Println(err)
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"mensaje": "sin autenticar",
		})
	}
	usuario := modelos.Usuario{}
	bbdd.DB.Where("id = ?", id).First(&usuario)
	fmt.Println(usuario)
	return c.JSON(usuario)
}

func Salir(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "nd-jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Expira en medio día
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"mensaje": "salido",
	})
}
