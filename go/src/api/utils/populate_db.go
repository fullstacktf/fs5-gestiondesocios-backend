package utils

import (
	"fmt"
	"fs5-gestiondesocios-backend/src/api/models"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	xj "github.com/basgys/goxml2json"
	"github.com/tidwall/gjson"
)

func Populate_db() {

	resp, err := http.Get("https://www.boardgamegeek.com/xmlapi/collection/El%20Club%20Dante")
	if err != nil {
		fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Read body: %v", err)
	}

	xml := strings.NewReader(string(data))
	json1, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	value := gjson.GetMany(json1.String(), "items.item.#.-objectid", "items.item.#.name.#content", "items.item.#.stats.rating.average.-value", "items.item.#.status.-lastmodified")
	if err != nil {
		log.Fatal(err)
	}

	arrIds := strings.Split(strings.ReplaceAll(value[0].String(), "\"", ""), ",")
	arrNames := strings.Split(strings.ReplaceAll(value[1].String(), "\"", ""), ",")
	arrRating := strings.Split(strings.ReplaceAll(value[2].String(), "\"", ""), ",")
	arrDate := strings.Split(strings.ReplaceAll(value[3].String(), "\"", ""), ",")

	game := models.Game{}

	games := []models.Game{}
	for i := 0; i < len(arrIds); i++ {
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		processedID := reg.ReplaceAllString(arrIds[i], "")
		game.ID, _ = strconv.ParseUint(processedID, 10, 0)
		game.GameName = arrNames[i]
		game.Rating, _ = strconv.ParseFloat(arrRating[i], 64)
		game.EntryDate = arrDate[i]
		game.Disponibility = true
		game.Comments = ""
		game.IDOwner = 1
		games = append(games, game)
	}

	db := GetConnection()
	db.FirstOrCreate(&games)
}
