package controladores

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	bbdd.DB.Model(&entradas).Clauses(clause.Returning{}).Where("especial", especial).Order("fecha desc").Offset(offset).Limit(limite).Update("visitas", gorm.Expr("visitas + ?", 1))
	bbdd.DB.Model(&modelos.Entrada{}).Where("especial", especial).Count(&total)
	var entradasConTotalCom []modelos.Entrada
	for _, e1 := range entradas {
		e1.CalcularTotalComentarios()
		entradasConTotalCom = append(entradasConTotalCom, e1)
	}
	return c.JSON(fiber.Map{
		"datos": entradasConTotalCom,
		"meta": fiber.Map{
			"total":         total,
			"pagina":        pagina,
			"ultima_pagina": math.Ceil(float64(int(total) / limite)),
		},
	})
}

func EntradasSinPaginar(c *fiber.Ctx) error {
	var entradas []modelos.Entrada
	bbdd.DB.Select("Id", "Titulo", "Fecha").Order("fecha desc").Find(&entradas)
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
	entrada := modelos.Entrada{
		Id: uint(id),
	}
	bbdd.DB.Preload("Comentarios").Find(&entrada)
	entrada.Visitas = entrada.Visitas + 1
	bbdd.DB.Updates(&entrada)
	entrada.CalcularTotalComentarios()
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
