package handlers

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/ozonmp/bss-office-facade/internal/metrics"
	"github.com/ozonmp/bss-office-facade/internal/model"
	"github.com/ozonmp/bss-office-facade/internal/repo"
	pb "github.com/ozonmp/bss-office-facade/pkg/bss-office-facade"
	"github.com/pkg/errors"
)

type eventHandler struct {
	repo repo.OfficeRepo
}

// NewEventHandler - создаёт экземпляр обработчика событий из кафки
func NewEventHandler(repo repo.OfficeRepo) *eventHandler {
	handler := &eventHandler{repo: repo}

	return handler
}

func (h *eventHandler) Handle(ctx context.Context, message *sarama.ConsumerMessage) error {
	var pbEvent pb.OfficeEvent
	err := proto.Unmarshal(message.Value, &pbEvent)

	if err != nil {
		return err
	}

	officeEvent := model.ConvertPbToBssOfficeEvent(&pbEvent)

	// в задании требовалось вывести в stdout
	//fmt.Printf("%#+v\n", officeEvent)

	metrics.IncTotalEvents()
	metrics.IncTotalCud(officeEvent.Type)

	if officeEvent.Payload.ID == 0 {
		return errors.Wrap(err, "EventHandler.Handle : officeID is nul")
	}

	switch officeEvent.Type {
	case model.Created:
		_, err := h.repo.CreateOffice(ctx, officeEvent.Payload)
		if err != nil {
			return err
		}
	case model.Updated:
		_, err := h.repo.UpdateOffice(ctx, officeEvent.Payload.ID, officeEvent.Payload)
		if err != nil {
			return err
		}
	case model.Removed:
		_, err := h.repo.RemoveOffice(ctx, officeEvent.Payload.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
