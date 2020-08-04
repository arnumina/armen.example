/*
#######
##                                                             __
##       ___ _______ _  ___ ___       _____ _____ ___ _  ___  / /__
##      / _ `/ __/  ' \/ -_) _ \  _  / -_) \ / _ `/  ' \/ _ \/ / -_)
##      \_,_/_/ /_/_/_/\__/_//_/ (_) \__/_\_\\_,_/_/_/_/ .__/_/\__/
##                                                    /_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package messages

import (
	"github.com/arnumina/armen.core/pkg/message"
	"github.com/arnumina/failure"
	"github.com/arnumina/logger"

	"github.com/arnumina/armen.example/internal/pkg/container"
)

type (
	// Handler AFAIRE.
	Handler struct {
		jwHandler container.JWHandler
		logger    *logger.Logger
	}
)

// NewHandler AFAIRE.
func NewHandler(ctn container.Container) *Handler {
	return &Handler{
		jwHandler: ctn.JWHandler(),
		logger:    ctn.Logger(),
	}
}

func (h *Handler) msgHandler(msg *message.Message) {
	h.logger.Info( /////////////////////////////////////////////////////////////////////////////////////////////////////
		"Consume message",
		"id", msg.ID,
		"topic", msg.Topic,
		"publisher", msg.Publisher,
	)

	switch msg.Topic {
	case "create.job.example.test":
		h.jwHandler.CreateJobTest(msg)
	default:
		h.logger.Error( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"The topic of this message is not valid",
			"id", msg.ID,
			"topic", msg.Topic,
			"publisher", msg.Publisher,
		)
	}
}

// Subscribe AFAIRE.
func (h *Handler) Subscribe(ctn container.Container) error {
	if err := ctn.Bus().Subscribe(
		h.msgHandler,
		"create[.]job[.]example[.].+",
	); err != nil {
		return failure.New(err).Msg("messages") ////////////////////////////////////////////////////////////////////////
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
