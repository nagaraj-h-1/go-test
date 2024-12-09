package event_handler

import (
	"fmt"
	"testing"
)

func TestGetAllEvents(t *testing.T) {
	EvHandler, err := NewEventHandlerIntf()
	if err != nil {
		fmt.Print("error getting evHandler Intf")
	}
	EvHandler.GetAllEvents()
}
