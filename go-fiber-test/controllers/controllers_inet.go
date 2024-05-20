package controllers

import (
	"go-fiber-test/database"
	m "go-fiber-test/models"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

	if err := c.BodyParser(empData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("error")
	}
	// Validate Username
	//		boolean			fuc					regular expression 		ข้อมูลจาก Models
	validateUsername, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, empData.Username) // ใช้ Fuc MatchString ในการหาค้นหาค่าใน regular expression  ว่า ใน empData.Username มีอักษรที่มีอยู่ใน REGEXP หรือไม่
	/*
		regex คือ การกำหนดรูปแบบหรือกลุ่มคำ
		เพื่อเอาไว้ใช้ค้นหาข้อความต่างๆตามที่เราต้องการ
		สามารถค้นหาได้ทั้งอักขระธรรมดา หรือค้นหาความข้อที่กำหนดไว้
		หรือจะเป็นอักขระพิเศษก็สามารถค้นหาได้
	*/
	if !validateUsername {
		return c.Status(fiber.StatusBadRequest).SendString(`ใช้ตัวอักษรภาษาอังกฤษ (a-z), (A-Z), ตัวเลข (0-9), หรือเครื่องหมาย ( _ )หรือ ( - ) เท่านั้น`)
	}
	// Validate Website
	validateWebsite, _ := regexp.MatchString(`^[a-z0-9-]+$`, empData.Website)
	if !validateWebsite {
		return c.Status(fiber.StatusBadRequest).SendString(`2-30 ตัวอักษรต้องเป็นตัวอักษรภาษาอังกฤษตัวเล็ก (a-z) ตัวเลข (0-9) ห้ามใช้เครื่องหมายอักขระพิเศษยกเว้นเครื่องหมายขีด (-) ห้ามเว้นวรรคและห้ามใช้ภาษาไทย`)
	}
	//
	if errors := validate.Struct(empData); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	return c.JSON(empData)
}

func GetDogIDBETWEEN50AND100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ?", 50).Where("dog_id < ?", 100)
}
func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs // ดึงข้อมูลเป็นแบบ Arrays

	db.Scopes(GetDogIDBETWEEN50AND100).Find(&dogs)
	// db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetALLDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs // ดึงข้อมูลเป็นแบบ Arrays

	db.Find(&dogs)
	// db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	// WHERE DATA
	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

// ADD DATA
func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn //*  สร้างตัวแปร DB เก็บการ Connect DB
	var dog m.Dogs        // * สร้างตัวแปร Dog เพื่อเรียก Dogs จาก Models

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog) // * คำสั่งสร้างข้อมูล
	return c.Status(201).JSON(dog)
}

// UPDATE DATA
func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id") // iD จาก parameter

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

// DELETE DATA
func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// GET DATA
func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	// Create ตัวรับค่า การนับ Sum Color
	countColor := map[string]int{
		"red":      0,
		"green":    0,
		"pink":     0,
		"no color": 0,
	}

	// สร้าง dataReults เพื่อใช้ DogsRes
	var dataResults []m.DogsRes
	countData := len(dogs)
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		countColor[typeStr]++ //นับค่า สีของ dog_id

		d := m.DogsRes{
			Name:  v.Name,  //inet1
			DogID: v.DogID, //113
			Type:  typeStr, //green
		}

		dataResults = append(dataResults, d)
	}
	r := m.ResultData{
		Count:       countData, //หาผลรวม,
		Data:        dataResults,
		Name:        "golang-test",
		Sum_Red:     countColor["red"],
		Sum_Green:   countColor["green"],
		Sum_Pink:    countColor["pink"],
		Sum_Nocolor: countColor["no color"],
	}

	return c.Status(200).JSON(r)
}

func GetDataDelete(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs)

	return c.Status(200).JSON(dogs)
}

// * Company API Func Group
func AddCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&company)
	return c.Status(201).JSON(company)
}

func GetCompony(c *fiber.Ctx) error {
	db := database.DBConn
	var company []m.Company

	db.Find(&company)
	return c.Status(200).JSON(company)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)

	return c.Status(200).JSON(company)
}

func DeleteCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Company

	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).SendString("Delete Company Successfully")
}

func GetComponyByID(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var company []m.Company

	// Serach ID
	result := db.Find(&company, "company_id", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&company)
}
