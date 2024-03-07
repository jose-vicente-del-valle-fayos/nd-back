package controladores

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math"
	"nd-back/bbdd"
	"nd-back/modelos"
	"strconv"
)

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
	bbdd.DB.Preload("Comentarios").Where("especial", especial).Order("fecha desc").Offset(offset).Limit(limite).Find(&entradas)
	for i := range entradas {
		bbdd.DB.Model(&entradas[i]).Update("visitas", gorm.Expr("visitas + ?", 1))
	}
	bbdd.DB.Model(&modelos.Entrada{}).Where("especial", especial).Count(&total)
	return c.JSON(fiber.Map{
		"datos": entradas,
		"meta": fiber.Map{
			"total":         total,
			"pagina":        pagina,
			"ultima_pagina": math.Ceil(float64(int(total) / limite)),
		},
	})
}

func EntradasSinPaginar(c *fiber.Ctx) error {
	var entradas []modelos.Entrada
	bbdd.DB.Select("Id", "Titulo", "Fecha", "Contenido").Order("fecha desc").Find(&entradas)
	return c.JSON(fiber.Map{
		"datos": entradas,
	})
}

func CrearEntrada(c *fiber.Ctx) error {
	/*
		{
			"id_us": 1,
			"usuario": "Chevi",
			"especial": false,
			"titulo": "Esta es una entrada fantástica",
			"fecha": "2024-02-22",
			"contenido": "Este es un contenido fantástico."
		}
	*/
	var entrada modelos.Entrada
	if err := c.BodyParser(&entrada); err != nil {
		return err
	}
	bbdd.DB.Create(&entrada)
	return c.JSON(entrada)
}

func LeerEntrada(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	bbdd.DB.Model(&modelos.Entrada{}).Where("id", id).Update("visitas", gorm.Expr("visitas + ?", 1))
	var entrada modelos.Entrada
	bbdd.DB.Preload("Comentarios").Find(&entrada, id)
	return c.JSON(fiber.Map{
		"datos": entrada,
	})
}

func ActualizarEntrada(c *fiber.Ctx) error {
	/*
		{
			"id_us": 1,
			"usuario": "Susanita",
			"especial": false,
			"titulo": "Esta es una entrada actualizada",
			"fecha": "2024-02-22",
			"contenido": "Este es un contenido actualizado.",
			"comentarios": []
		}
	*/
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
	bbdd.DB.Model(&entrada).Updates(entrada)
	return c.JSON(entrada)
}

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
