package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	gf "github.com/shareed2k/goth_fiber"
)

func GoogleCallback(c *fiber.Ctx) error {
	user, err := gf.CompleteUserAuth(c)
	log.Fatal(user, "user")
	log.Fatal(err, "error")
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
		log.Fatal(gothUser)
		c.JSON(gothUser)
	} else {
		log.Fatal("did not work")
		gf.BeginAuthHandler(c)
	}
	return nil
}
