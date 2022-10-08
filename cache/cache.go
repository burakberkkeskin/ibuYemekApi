package cache

import (
	"fmt"
	"ibu-yemek-api/models"
	"ibu-yemek-api/services"

	"github.com/robfig/cron/v3"
)

var lunchToday models.Lunch
var lunchTomorrow models.Lunch

func GetLunch(day *string) *models.Lunch {

	if *day == "today" {
		return &lunchToday
	} else if *day == "tomorrow" {
		return &lunchTomorrow
	} else {
		return nil
	}

}

func CheckDate(day *string) bool {
	if *day == "today" {
		return true
	} else if *day == "tomorrow" {
		return true
	} else {
		return false
	}
}

func IsEmpty(day string) bool {
	if GetLunch(&day).AnaYemek == "" {
		return true
	} else {
		return false
	}
}

func init() {
	fmt.Println("Cache initialized")

	lunchToday = services.Scrapper("today")
	fmt.Println("Today's lunch cached")

	lunchTomorrow = services.Scrapper("tomorrow")
	fmt.Println("Tomorrow's lunch cached")

	c := cron.New()
	c.AddFunc("10 03 * * *", func() {
		lunchToday = services.Scrapper("today")
		lunchTomorrow = services.Scrapper("tomorrow")
	})

	c.Start()
}
