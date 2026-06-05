package events

import "NickBrisebois/BadMovieSpinnerGo/pkg/models"

type UIEvent int

const (
	EventTypeSpinButtonClicked UIEvent = iota
	EventTypeSuggestedByChanged
)

type EventCallbackData struct {
	EventType      UIEvent
	SuggestedUsers *[]models.PersonName
}

type EventCallback func(data *EventCallbackData)
