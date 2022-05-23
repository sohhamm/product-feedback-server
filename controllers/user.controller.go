package controllers

import (
	"github.com/gofiber/fiber/v2"
	gf "github.com/shareed2k/goth_fiber"
)

func GoogleCallback(c *fiber.Ctx) error {
	user, err := gf.CompleteUserAuth(c)
	if err != nil {
		return err
	}
	c.JSON(user)
	return nil
}

func GoogleLogout(c *fiber.Ctx) error {
	gf.Logout(c)
	c.Redirect("/")
	return nil
}

func Login(c *fiber.Ctx) error {
	if gothUser, err := gf.CompleteUserAuth(c); err == nil {
		c.JSON(gothUser)
	} else {
		gf.BeginAuthHandler(c)
	}
	return nil
}
