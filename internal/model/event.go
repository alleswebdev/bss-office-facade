package model

import (
	"database/sql"
	pb "github.com/ozonmp/bss-office-facade/pkg/bss-office-facade"
	"time"
)

type EventType string

// EventStatus enum for event status
type EventStatus uint8

// Created - событие создано
// Updated - событие обновлено
// Removed - событие удалено
const (
	Created = "created"
	Updated = "updated"
	Removed = "removed"
)

// OfficeEvent - office event model
type OfficeEvent struct {
	ID       uint64        `db:"id"`
	OfficeID uint64        `db:"office_id"`
	Type     EventType     `db:"type"`
	Status   EventStatus   `db:"status"`
	Created  time.Time     `db:"created_at"`
	Updated  sql.NullTime  `db:"updated_at"`
	Payload  OfficePayload `db:"payload"`
}

//OfficePayload Сктура для записи информации о изменениях в сущности office
type OfficePayload struct {
	ID          uint64 `json:"id,string"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Removed     bool   `json:"removed,omitempty"`
}

func ConvertPbToBssOfficeEvent(pb *pb.OfficeEvent) *OfficeEvent {
	officeEvent := &OfficeEvent{
		ID:       pb.GetId(),
		OfficeID: pb.GetOfficeId(),
		Status:   EventStatus(pb.GetStatus()),
		Type:     EventType(pb.GetType()),
		Created:  pb.GetCreated().AsTime(),
		Updated: sql.NullTime{
			Time: pb.GetUpdated().AsTime(),
		},
		Payload: ConvertPbToBssOfficePayload(pb.GetPayload()),
	}

	return officeEvent
}

func ConvertPbToBssOfficePayload(pb *pb.OfficePayload) OfficePayload {
	payload := OfficePayload{
		ID:          pb.GetId(),
		Name:        pb.GetName(),
		Description: pb.GetDescription(),
	}

	return payload
}
