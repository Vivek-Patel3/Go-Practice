package producer

import (
	"time"
	"fmt"
	"github.com/Vivek-Patel3/Producer-Consumer/internal/model"
	"github.com/Vivek-Patel3/Producer-Consumer/internal/topic"
)
// tasks of producer: 
/*
1. produce events
2. push those events into streaming medium - thus promoting loose coupling
*/

func Producer(id int64, t *topic.Topic) {
	event := model.Event{
		Id:        id,
		Payload:   "event_payload",
		CreatedAt: time.Now(),
	}

	t.Send(event)
	fmt.Println("Produced event", id)
}