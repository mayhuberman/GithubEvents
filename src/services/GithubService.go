package services

import (
	"encoding/json"
	"fmt"
	"githubEventsListener/models"
	"githubEventsListener/utils"
	"io"
	"log"
	"net/http"
	"sort"
)

var EventTypes = make(map[string]int)

var ActorsList models.LinkedList
var actorsMap = make(map[string]int)
var ActorsSize = 50

var UrlsList models.LinkedList
var UrlsSize = 20

var Emails = make(map[string]bool)

func ReturnData(w http.ResponseWriter) {
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"EventTypes": EventTypes,
		"actors":     ActorsList.ConvertToSlice(),
		"urls":       UrlsList.ConvertToSlice(),
		"Emails":     utils.GetMapKeys(Emails),
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalln("An error occurred while trying to encode json: ", data, err)
		return
	}
}

func ReturnReposStars(w http.ResponseWriter) {
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	repoStars := getRepoStars(UrlsList)

	data := map[string]interface{}{
		"repoStars": repoStars.ConvertUrlStarsToStrings(),
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatalln("An error occurred while trying to encode json: ", data, err)
		return
	}
}

func ExtractData(events []models.Event) {
	log.Println("Extracting data from Github events")

	index := 0

	for _, event := range events {

		EventTypes[event.Type] = EventTypes[event.Type] + 1

		var commits = event.Payload.Commits
		for _, commit := range commits {
			if commit.Author.Email != "" {
				Emails[commit.Author.Email] = true
			}
		}

		actorName := event.Actor.Login
		if ActorsList.Size() < ActorsSize {
			if !ActorsList.Search(actorName) {
				ActorsList.Append(actorName)
				actorsMap[actorName] = index
				index = index + 1
			} else {
				currIndex := actorsMap[actorName]
				actorsMap[actorName] = index - 1
				for key, value := range actorsMap {
					if value >= currIndex && key != actorName {
						actorsMap[key] = value - 1
					}
				}
			}
		} else {
			if !ActorsList.Search(actorName) {
				for key, value := range actorsMap {
					if value == 0 {
						ActorsList.Detach(key)
						delete(actorsMap, key)
						break
					}
				}
				for key, value := range actorsMap {
					actorsMap[key] = value - 1
				}
				ActorsList.Append(actorName)
				actorsMap[actorName] = ActorsSize - 1
			} else {
				currIndex := actorsMap[actorName]
				actorsMap[actorName] = ActorsSize - 1
				for key, value := range actorsMap {
					if value > currIndex && key != actorName {
						actorsMap[key] = value - 1
					}
				}
			}
		}

		url := event.Repo.URL
		if UrlsList.Size() >= UrlsSize {
			UrlsList.DetachHead()
		}
		UrlsList.Append(url)
	}
}

func getRepoStars(urlsList models.LinkedList) models.UrlStarsSlice {
	log.Println("Getting repo stars from Github API")

	var urlsStars []models.UrlStars

	for urlsList.Size() > 0 {
		url := urlsList.GetHeadDataAndMoveNext()
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln("An error occurred while trying to get the repo stars from Github. url: ", url, err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("An error occurred while trying to read the response from Github.: ", err)
			return nil
		}
		var repo models.Repo
		err = json.Unmarshal(body, &repo)
		if err != nil {
			fmt.Println("An error occurred while trying to unmarshal the response from Github.: ", err)
			return nil
		}
		newPair := models.UrlStars{Url: url, Stars: repo.StarsCount}
		urlsStars = append(urlsStars, newPair)
	}
	sort.Slice(urlsStars, func(i, j int) bool {
		return urlsStars[i].Stars < urlsStars[j].Stars
	})
	return urlsStars
}
