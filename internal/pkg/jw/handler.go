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

package jw

import (
	"fmt"

	"github.com/arnumina/armen.core/pkg/jw"
	"github.com/arnumina/armen.core/pkg/message"

	"github.com/arnumina/armen.example/internal/pkg/container"
)

const (
	_jobTest = "test"
)

type (
	// Handler AFAIRE.
	Handler struct {
		model  container.Model
		plugin container.Plugin
	}
)

// NewHandler AFAIRE.
func NewHandler(ctn container.Container) *Handler {
	return &Handler{
		model:  ctn.Model(),
		plugin: ctn.Plugin(),
	}
}

func (h *Handler) createJob(id, n, t, o string, pr jw.Priority, k, e *string) *jw.Job {
	return jw.NewJob(id, n, h.plugin.Name(), t, o, pr, k, e)
}

func (h *Handler) maybeInsertJob(job *jw.Job) (bool, error) {
	return h.model.MaybeInsertJob(job)
}

// CreateJobTest AFAIRE.
func (h *Handler) CreateJobTest(msg *message.Message) {
	k := fmt.Sprintf("%s.%s", h.plugin.Name(), _jobTest)
	job := h.createJob(msg.ID, _jobTest, _jobTest, msg.Publisher, jw.Low, &k, nil)
	_, _ = h.maybeInsertJob(job)
}

// InsertWorkflow AFAIRE.
func (h *Handler) InsertWorkflow(wf *jw.Workflow) error {
	return h.model.InsertWorkflow(wf)
}

/*
######################################################################################################## @(°_°)@ #######
*/
