package controllers

import (
	"ibu-yemek-api/cache"
	"ibu-yemek-api/responses"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetLunch(c *fiber.Ctx) error {
	log.Println("GetLaunch Request For", c.Params("day"))

	lunch, err := cache.GetLunch(c.Params("day"))

	if err != nil {
		response := responses.LunchResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &fiber.Map{"data": err.Error()},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	} else if lunch.AnaYemek == "" {
		response := responses.LunchResponse{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Data:    &fiber.Map{"data": "Yemek Bulunamadı"},
		}
		return c.Status(http.StatusBadRequest).JSON(response)
	} else {
		response := responses.LunchResponse{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    lunch,
		}
		return c.Status(http.StatusOK).JSON(response)
	}

}
