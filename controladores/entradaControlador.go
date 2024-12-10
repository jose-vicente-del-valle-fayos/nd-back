package controladores

import (
	"fmt"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"io"
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
	bbdd.DB.Preload("Comentarios").Where("especial = ?", especial).Order("fecha desc").Offset(offset).Limit(limite).Find(&entradas)
	for i := range entradas {
		bbdd.DB.Model(&entradas[i]).Update("visitas", gorm.Expr("visitas + ?", 1))
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
	var entrada modelos.Entrada
	if err := c.BodyParser(&entrada); err != nil {
		return err
	}
	if entrada.ValidarFecha() && entrada.ValidarIdUs() && entrada.ValidarUsuario() && entrada.ValidarTitulo() && entrada.ValidarContenido() {
		bbdd.DB.Create(&entrada)
		fmt.Println("entrada: " + strconv.Itoa(int(entrada.Id)))
		entrada.Imagen = SubirImagen(c, int(entrada.Id))
		bbdd.DB.Model(&entrada).Where("id = ?", entrada.Id).Updates(entrada)
		return c.JSON(entrada)
	}
	return c.JSON(fiber.Map{"mensaje": "error de validación"})
}

// LeerEntrada reads an entry taking entry's id as an URL parameter
func LeerEntrada(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	bbdd.DB.Model(&modelos.Entrada{}).Where("id = ?", id).Update("visitas", gorm.Expr("visitas + ?", 1))
	var entrada modelos.Entrada
	bbdd.DB.Preload("Comentarios").Find(&entrada, id)
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
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	entrada := modelos.Entrada{
		Id: uint(id),
	}
	if err := c.BodyParser(&entrada); err != nil {
		return err
	}
	if entrada.ValidarFecha() && entrada.ValidarIdUs() && entrada.ValidarUsuario() && entrada.ValidarTitulo() && entrada.ValidarContenido() {
		bbdd.DB.Model(&entrada).Updates(SubirImagen(c, int(entrada.Id)))
		return c.JSON(entrada)
	}
	return c.JSON(fiber.Map{"mensaje": "error de validación"})
}

// BorrarEntrada deletes an entry taking the entry's id as an URL parameter
func BorrarEntrada(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	entrada := modelos.Entrada{
		Id: uint(id),
	}
	bbdd.DB.Delete(&entrada)
	return nil
}

// SubirImagen uploads any file to Cloudinary
func SubirImagen(c *fiber.Ctx, id int) string {
	fmt.Println(id)
	idStr := strconv.Itoa(id)
	fileHeader, err := c.FormFile("imagen-entrada")
	if err != nil {
		fmt.Printf("no se pudo obtener el archivo:", err)
		return "sin-imagen"
	}
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Printf("error al abrir el archivo:", err)
		return "sin-imagen"
	}
	defer file.Close()
	fileContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error al leer el archivo:", err)
		return "sin-imagen"
	}
	cloudName := os.Getenv("CLOUD_NAME")
	apiKey := os.Getenv("CLOUD_API_KEY")
	apiSecret := os.Getenv("CLOUD_API_SECRET")
	if cloudName == "" || apiKey == "" || apiSecret == "" {
		fmt.Printf("faltan las credenciales de Cloudinary")
		return "sin-imagen"
	}
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		fmt.Printf("error al inicializar Cloudinary:", err)
		return "sin-imagen"
	}
	upload, err := cld.Upload.Upload(c.Context(), fileContent, uploader.UploadParams{
		PublicID:         "nd/" + idStr,
		ResourceType:     "auto",
		Overwrite:        true,
		FilenameOverride: idStr,
	})
	if err != nil {
		fmt.Printf("error al subir el archivo:", err)
		return "sin-imagen"
	}
	fmt.Printf("archivo subido con éxito:", upload.SecureURL)
	return upload.SecureURL
}

/*
func SubirImagen(c *fiber.Ctx, id int) string {
	fmt.Println(id)
	idStr := strconv.Itoa(id)
	fileHeader, err := c.FormFile("imagen-entrada")
	if err != nil {
		return "sin-imagen"
	}
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fileContent, _ := io.ReadAll(file)
	cloudName := os.Getenv("CLOUD_NAME")
	apiKey := os.Getenv("CLOUD_API_KEY")
	apiSecret := os.Getenv("CLOUD_API_SECRET")
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		fmt.Println(err)
	}
	upload, err := cld.Upload.Upload(c.Context(), fileContent, uploader.UploadParams{PublicID: "nd/" + strconv.Itoa(id), ResourceType: "auto", Overwrite: true, FilenameOverride: idStr})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(upload.SecureURL)
	return upload.SecureURL
}
*/
