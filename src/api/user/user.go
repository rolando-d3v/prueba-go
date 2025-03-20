package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rolando-d3v/prueba/src/config"
)


type User struct {
	ID   int    `db:"id_respuesta_i" json:"id_respuesta_i"`
	Name string `db:"respuesta_t" json:"respuesta_t"`
}

func GetUsers(c *fiber.Ctx) error {
	var users []User
	query := `SELECT id_respuesta_i, respuesta_t FROM respuestas`
	// query := `SELECT id_role_i, desc_corta_v FROM role`
	err := config.DB.Select(&users, query)
	if err != nil {
	    return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
	        "error": "Failed to fetch users",
	    })
	}

	return c.JSON(users)
}

// return c.JSON(fiber.Map{
// 	"message": "Hello, World! 🌍"})