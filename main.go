package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/paudev/landing-page-be/models"
	"github.com/paudev/landing-page-be/storage"
	"gorm.io/gorm"
)

type Characters struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetCharacters(c *fiber.Ctx) error {
	characterModels := &[]models.Characters{}
	roleCharacterModels := &[]models.Characters{}
	limitQuery := c.Query("limit", "10")
	limit, err := strconv.ParseInt(limitQuery, 10, 64)
	if limit < 1 || limit > 100 {
		limit = 10
	}

	if err != nil {
		return c.Status(500).JSON("Invalid limit")
	}

	err = r.DB.Order("id asc").Limit(int(limit)).Find(characterModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get character"})
		return err
	}

	err = r.DB.Order("id desc").Limit(int(limit)).Find(roleCharacterModels).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get role character"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"characters": characterModels, "roleCharacters": roleCharacterModels})
	return nil
}

func (r *Repository) GetCharacterByID(c *fiber.Ctx) error {
	id := c.Params("id")
	characterModel := &models.Characters{}

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "id cannot be empty"})
	}

	fmt.Println("this ID is ", id)

	err := r.DB.Where("id = ?", id).First(characterModel).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get character"})
		return err
	}

	c.Status(http.StatusOK).JSON(&fiber.Map{"data": characterModel})
	return nil
}

func (r *Repository) GetFantasies(c *fiber.Ctx) error {
	fantasyModels := &[]models.Fantasies{}
	err := r.DB.Find(fantasyModels).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get fantasies"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"fantasies": fantasyModels})
	return nil
}

func (r *Repository) GetFantasyByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fantasyModel := &models.Fantasies{}

	if id == "" {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "id cannot be empty"})
	}

	fmt.Println("this ID is ", id)

	err := r.DB.Where("id = ?", id).First(fantasyModel).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get character"})
		return err
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"data": fantasyModel})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/character/:id", r.GetCharacterByID)
	api.Get("/characters", r.GetCharacters)
	api.Get("/fantasy/:id", r.GetFantasyByID)
	api.Get("/fantasies", r.GetFantasies)
}

func main() {

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load database")
	}

	err = models.MigrateCharacters(db)

	if err != nil {
		log.Fatal("could not migrate characters db")
	}

	err = models.MigrateFantasies(db)

	if err != nil {
		log.Fatal("could not migrate fantasies db")
	}
	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":" + os.Getenv("APP_PORT"))
}
