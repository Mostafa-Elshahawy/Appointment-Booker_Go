package handlers

import (
	"mostafaba29/intialization"
	"mostafaba29/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	var userData models.User
	if err := c.BodyParser(&userData); err != nil {
		return c.Status(400).JSON(err.Error())

	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 14)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to hash password",
		})
	}

	user := models.User{Username: userData.Username, Password: hashedPass, Position: userData.Position}
	result := intialization.DB.Create(&user)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"massege": "failed to create user",
		})
	}

	return c.Status(200).JSON(user, "account created successfully")
}

func Login(c *fiber.Ctx) error {

	var userInfo struct {
		Username string
		Password string
	}

	if err := c.BodyParser(&userInfo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "couldn't read user data",
		})
	}

	var user models.User
	intialization.DB.Where("username=?", userInfo.Username).First(&userInfo)

	if user.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"massege": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(userInfo.Password)); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"massege": "incorrect password",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"massege": "logged in",
	})
}

func Logout(c *fiber.Ctx) error {

}
