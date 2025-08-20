package controladores

import (
	"fmt"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math"
	"nd-back/bbdd"
	"nd-back/modelos"
	"os"
	"strconv"
)

// TodasEntradas returns an array of entries inside datos, with some metadata. It can be called with some URL parameters like limite, pagina and especial
func TodasEntradas(c *fiber.Ctx) error {
	limite, err := strconv.Atoi(c.Query("limite", strconv.Itoa(math.MaxInt32)))
	if err != nil {
		return err
	}
	pagina, err := strconv.Atoi(c.Query("pagina", "1"))
	if err != nil {
		return err
	}
	especial, err := strconv.ParseBool(c.Query("especial", "false"))
	if err != nil {
		return err
	}
	offset := (pagina - 1) * limite
	var total int64
	var entradas []modelos.Entrada
	bbdd.DB.Preload("Comentarios", func(db *gorm.DB) *gorm.DB {
		return db.Order("fecha ASC")
	}).Where("especial = ?", especial).Order("fecha desc").Offset(offset).Limit(limite).Find(&entradas)
	for i := range entradas {
		entradas[i].CalcularTotalComentarios()
	}
	bbdd.DB.Model(&modelos.Entrada{}).Where("especial = ?", especial).Count(&total)
	return c.JSON(fiber.Map{
		"datos": entradas,
		"meta": fiber.Map{
			"total":         total,
			"pagina":        pagina,
			"ultima_pagina": math.Ceil(float64(int(total) / limite)),
		},
	})
}

// ExtractoTodas returns an array with entries extract data inside datos
func ExtractoTodas(c *fiber.Ctx) error {
	var entradas []modelos.Entrada
	bbdd.DB.Select("Id", "Titulo", "Fecha", "Contenido").Order("fecha desc").Find(&entradas)
	/*
		if len(entradas) > 0 {
			// Actualizar las visitas de todas las entradas a 0
			for i := range entradas {
				entradas[i].Visitas = 0
			}
			// Actualiza todas las entradas a la vez con las visitas a 0
			bbdd.DB.Model(&entradas).Updates(map[string]interface{}{"Visitas": 0})
		}
	*/
	return c.JSON(fiber.Map{
		"datos": entradas,
	})
}

// CrearEntrada creates an entry
//
//	{
//		"id_us": 1,
//		"usuario": "Chevi",
//		"especial": false,
//		"titulo": "Esta es una entrada fantástica",
//		"fecha": "2024-02-22",
//		"contenido": "Este es un contenido fantástico."
//	}
func CrearEntrada(c *fiber.Ctx) error {
	idUs, _ := strconv.ParseUint(c.FormValue("id_us"), 10, 32)
	usuario := c.FormValue("usuario")
	especial, _ := strconv.ParseBool(c.FormValue("especial"))
	titulo := c.FormValue("titulo")
	fecha := c.FormValue("fecha")
	contenido := c.FormValue("contenido")
	entrada := modelos.Entrada{
		IdUs:      uint(idUs),
		Usuario:   usuario,
		Especial:  &especial,
		Titulo:    titulo,
		Fecha:     fecha,
		Contenido: contenido,
	}
	// if err := c.BodyParser(&entrada); err != nil {
	// 	return err
	// }
	if entrada.ValidarFecha() && entrada.ValidarUsuario() && entrada.ValidarTitulo() && entrada.ValidarContenido() {
		bbdd.DB.Create(&entrada)
		fmt.Println("entrada: " + strconv.Itoa(int(entrada.Id)))
		entrada.Imagen = SubirImagen(c, entrada.Id)
		bbdd.DB.Model(&entrada).Where("id = ?", entrada.Id).Updates(entrada)
		return c.JSON(entrada)
	}
	return c.JSON(fiber.Map{"mensaje": "error de validación"})
}

// LeerEntrada reads an entry taking entry's id as a URL parameter
func LeerEntrada(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	var entrada modelos.Entrada
	bbdd.DB.Preload("Comentarios", func(db *gorm.DB) *gorm.DB {
		return db.Order("fecha ASC")
	}).Find(&entrada, id)
	entrada.CalcularTotalComentarios()
	return c.JSON(fiber.Map{
		"datos": entrada,
	})
}

// ActualizarEntrada updates an entry
//
//	{
//		"id_us": 1,
//		"usuario": "Susanita",
//		"especial": false,
//		"titulo": "Esta es una entrada actualizada",
//		"fecha": "2024-02-22",
//		"contenido": "Este es un contenido actualizado.",
//		"comentarios": []
//	}
func ActualizarEntrada(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.FormValue("id"), 10, 32)
	BorrarImagen(c, uint(id))
	idUs, _ := strconv.ParseUint(c.FormValue("id_us"), 10, 32)
	usuario := c.FormValue("usuario")
	especial, _ := strconv.ParseBool(c.FormValue("especial"))
	titulo := c.FormValue("titulo")
	fecha := c.FormValue("fecha")
	contenido := c.FormValue("contenido")
	entrada := modelos.Entrada{
		Id:        uint(id),
		IdUs:      uint(idUs),
		Usuario:   usuario,
		Especial:  &especial,
		Titulo:    titulo,
		Fecha:     fecha,
		Contenido: contenido,
		Imagen:    SubirImagen(c, uint(id)),
	}
	if entrada.ValidarFecha() && entrada.ValidarUsuario() && entrada.ValidarTitulo() && entrada.ValidarContenido() {
		bbdd.DB.Model(&entrada).Where("id = ?", id).Updates(entrada)
		return c.JSON(entrada)
	}
	return c.JSON(fiber.Map{"mensaje": "error de validación"})
}

// BorrarEntrada deletes an entry taking the entry's id as a URL parameter
func BorrarEntrada(c *fiber.Ctx) error {
	idUs, err := strconv.Atoi(c.Params("id_us"))
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	var entrada modelos.Entrada
	bbdd.DB.Find(&entrada, id)
	if entrada.IdUs == uint(idUs) {
		BorrarImagen(c, uint(id))
		bbdd.DB.Delete(&entrada)
	}
	return nil
}

// SubirImagen uploads a file (jpg, jpeg or mp3) to Cloudinary
func SubirImagen(c *fiber.Ctx, id uint) string {
	idStr := strconv.FormatUint(uint64(id), 10)
	fileHeader, err := c.FormFile("imagen-entrada")
	if err != nil {
		fmt.Println("no se pudo obtener el archivo: ", err)
		return "sin-imagen"
	}
	if fileHeader.Size == 0 {
		fmt.Println("el archivo está vacío")
		return "sin-imagen"
	}
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("error al abrir el archivo: ", err)
		return "sin-imagen"
	}
	cloudName := os.Getenv("CLOUD_NAME")
	apiKey := os.Getenv("CLOUD_API_KEY")
	apiSecret := os.Getenv("CLOUD_API_SECRET")
	if cloudName == "" || apiKey == "" || apiSecret == "" {
		fmt.Println("faltan las credenciales de cloudinary.")
		return "sin-imagen"
	}
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		fmt.Println("error al inicializar cloudinary: ", err)
		return "sin-imagen"
	}
	upload, err := cld.Upload.Upload(c.Context(), file, uploader.UploadParams{
		PublicID:     "nd/" + idStr,
		ResourceType: "auto",
		Overwrite:    true,
	})
	if err != nil {
		fmt.Println("error al subir el archivo: ", err)
		return "sin-imagen"
	}
	err = file.Close()
	if err != nil {
		fmt.Println("error al cerrar el archivo: ", err)
		return "sin-imagen"
	}
	fmt.Println("archivo subido: ", upload.SecureURL)
	return upload.SecureURL
}

// BorrarImagen destroys a file (jpg, jpeg or mp3) from Cloudinary
func BorrarImagen(c *fiber.Ctx, id uint) bool {
	idStr := strconv.FormatUint(uint64(id), 10)
	cloudName := os.Getenv("CLOUD_NAME")
	apiKey := os.Getenv("CLOUD_API_KEY")
	apiSecret := os.Getenv("CLOUD_API_SECRET")
	if cloudName == "" || apiKey == "" || apiSecret == "" {
		fmt.Println("faltan las credenciales de cloudinary.")
		return false
	}
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		fmt.Println("error al inicializar cloudinary: ", err)
		return false
	}
	_, err = cld.Upload.Destroy(c.Context(), uploader.DestroyParams{
		PublicID: "nd/" + idStr,
	})
	if err != nil {
		fmt.Println("error al destruir el archivo: ", err)
		return false
	}
	fmt.Println("archivo eliminado.")
	return true
}
