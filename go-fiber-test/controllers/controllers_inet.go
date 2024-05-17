package controllers

import (
	m "go-fiber-test/models"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!") //* แสดง Hello World !
}

func InputFactorial(c *fiber.Ctx) error {
	numbers := c.Params("number")
	this_number, err := strconv.Atoi(numbers)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Invalid number")
	}
	x := 1 // * กำหนดค่าเริ่มต้น
	// เริ่มต้นด้วย 5 , 4, 3, 2, 1
	for i := this_number; i > 0; i-- {
		x *= i
		// 5(x) = 1 * 5(i)  // *i
		// 20 =  4(x) * 4(i)
		// 60 =  20(x) * 3(i)
		// 120 =  60(x) * 2(i)
		// 120 =  120(x) * 1(i)
	}
	result := c.Params("number") + "!=" + strconv.Itoa(x) //แปลงจาก int เป็น string
	return c.JSON(result)
}

func CornvertAscii(c *fiber.Ctx) error {
	// names := c.Params("name")
	queryparam := c.Query("tax_id") // ถ้า Search มาจาก หน้าบ้าน จะเก็บใน a

	var result string
	var text string

	for _, v := range queryparam {
		result = strconv.Itoa(int(v))
		/*
		* เริ่มต้นคือ นำค่า v ที่ได้มาจากการ Query มาแปลงเป็น int32 ก่อน
		* แปลงเสร็จแล้ว แปลงค่าเป็น string
		* แล้วนำค่าที่แปลงเป็น string มาต่อกัน เพื่อสร้าง string ใหม่
		 */
		text += " " + result
	}

	return c.JSON(text)
}

func RegisterEmployee(c *fiber.Ctx) error {
	empData := new(m.Employees)
	validate := validator.New()
	// checkErrors := validate.Struct(empData)

	if err := c.BodyParser(empData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("error")
	}
	if errors := validate.Struct(empData); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	return c.JSON(empData)
}
