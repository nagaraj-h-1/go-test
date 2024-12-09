package event_handler

import (
	"encoding/json"
	"fmt"
	"gocode/constants"
	structs "gocode/structs"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func NewEventHandlerIntf() (EventHandler, error) {
	evHandler := &eventHandler{}
	evHandler.eventIDToEventMap = make(map[int]structs.Event)
	return evHandler, nil
}

func (event *eventHandler) FetchEventByID(eventID int) (structs.Event, error) {
	eventFetchUrl := constants.EventPrefixURL + "/" + strconv.Itoa(eventID) + ".json"

	response, err := http.Get(eventFetchUrl)
	if err != nil {
		fmt.Println("Error occured while calling url ", eventFetchUrl, err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error occured while reading response body ", response, err)
	}

	var eventData structs.Event
	if err = json.Unmarshal(data, &eventData) ; err != nil {
		log.Fatalf("Error parsing JSON for event id %d : %s", eventID, err)
	}

	return eventData, nil
}

func (event *eventHandler) GetEventByID(eventID int) (structs.Event, error){
	if _, ok := event.eventIDToEventMap[eventID]; !ok {
		return structs.Event{}, fmt.Errorf("event ID doesn't  exist, please fetch show first")
	}
	return event.eventIDToEventMap[eventID], nil
}


func (event *eventHandler) GetAllEvents() {

	eventFetchUrl := constants.EventPrefixURL

	response, err := http.Get(eventFetchUrl)
	if err != nil {
		fmt.Println("Error occured while calling url ", eventFetchUrl, err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error occured while reading response body ", response, err)
	}

	var contents []structs.GitHubContent
	if err := json.Unmarshal(data, &contents); err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}

	for _, content := range contents {
		eventID, err := strconv.Atoi(strings.Split(content.Name, ".")[0])
		if err != nil {
			log.Fatal("error while converting event id to int, please check event filename..", content.Name)
		}
		eventData, err := event.FetchEventByID(eventID)
		if err != nil {
			log.Fatal("Error while fetching event by id", eventID, err)
		}
		event.eventIDToEventMap[eventID] = eventData
		fmt.Println(eventID, eventData)
	}

}
