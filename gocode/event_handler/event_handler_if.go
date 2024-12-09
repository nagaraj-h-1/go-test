package event_handler

import structs "gocode/structs"

type eventHandler struct {
	eventIDToEventMap map[int]structs.Event
}

type EventHandler interface {
	GetAllEvents()
	GetEventByID(eventID int) (structs.Event, error)
	FetchEventByID(eventID int) (structs.Event, error)
}