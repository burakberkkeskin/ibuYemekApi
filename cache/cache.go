package cache

import (
	"errors"
	"ibu-yemek-api/models"
	"ibu-yemek-api/services"

	"github.com/robfig/cron/v3"
)

var lunchToday models.Lunch
var lunchTomorrow models.Lunch

func GetLunch(day string) (*models.Lunch, error) {

	if day == "today" {
		return &lunchToday, nil
	} else if day == "tomorrow" {
		return &lunchTomorrow, nil
	} else {
		return nil, errors.New("False Day")
	}

}

func init() {
	lunchToday = services.Scrapper("today")
	lunchTomorrow = services.Scrapper("tomorrow")

	c := cron.New()
	c.AddFunc("10 03 * * *", func() {
		lunchToday = services.Scrapper("today")
		lunchTomorrow = services.Scrapper("tomorrow")
	})

	c.Start()
}
