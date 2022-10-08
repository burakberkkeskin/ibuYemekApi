package services

import (
	"fmt"
	"ibu-yemek-api/models"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func getHtml() (string, error) {
	var url = "http://ibu.edu.tr/yemek-listesi"
	resp, err := http.Get(url)
	if err != nil {
		log.Default().Println("Error while getting html")
		return "", err
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// show the HTML code as a string %s
	return string(html), nil

}

func Scrapper(day string) models.Lunch {

	html, err := getHtml()
	if err != nil {
		log.Default().Println("Error while getting html")
		for i := 0; i < 3; i++ {
			html, err = getHtml()
			if err == nil {
				break
			}
			fmt.Println("Waiting 5 seconds to get food list again")
			time.Sleep(5 * time.Second)
		}
	}
	var unixDate int64
	if day == "today" {
		date := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, 21, 0, 0, 0, time.UTC)
		unixDate = date.Unix()
	} else if day == "tomorrow" {
		date := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 21, 0, 0, 0, time.UTC)
		unixDate = date.Unix()
	} else {
		log.Println("Invalid day")
	}

	searchString := strconv.FormatInt(unixDate, 10) + `]">([A-Za-zğüşöçıİĞÜŞÖÇ].*)</span>`
	re := regexp.MustCompile(searchString)
	res := re.FindAllStringSubmatch(html, -1)

	if res == nil {
		return models.Lunch{Corba: "", AnaYemek: "", YardimciAnaYemek: "", YanYemek1: "", YanYemek2: ""}
	} else {
		lunch := models.Lunch{
			Corba:            res[0][1],
			AnaYemek:         res[1][1],
			YardimciAnaYemek: res[2][1],
			YanYemek1:        res[3][1],
			YanYemek2:        res[4][1],
		}

		return lunch
	}

}
