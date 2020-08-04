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

package test

import (
	"github.com/arnumina/armen.core/pkg/jw"
	"github.com/arnumina/failure"
	"github.com/arnumina/uuid"

	"github.com/arnumina/armen.example/internal/pkg/job"
)

type (
	// Job AFAIRE.
	Job struct {
		*job.Job
	}
)

// New AFAIRE.
func New(job *job.Job) *Job {
	return &Job{
		Job: job,
	}
}

// Run AFAIRE.
func (j *Job) Run() *jw.Result {
	steps := map[string]*jw.Step{
		"alpha": {
			Plugin: "example",
			Type:   "nop",
			Config: nil,
			Next: map[string]interface{}{
				"default": "beta",
			},
		},
		"beta": {
			Plugin: "example",
			Type:   "nop",
			Config: map[string]interface{}{
				"day": 19,
			},
			Next: map[string]interface{}{
				"default": "gamma",
			},
		},
		"gamma": {
			Plugin: "example",
			Type:   "nop",
			Config: nil,
			Next:   nil,
		},
	}

	wf := jw.NewWorkflow(uuid.New(), "abg", "Workflow de test", "test", jw.None, "alpha", steps, nil)

	if err := j.Ctn.JWHandler().InsertWorkflow(wf); err != nil {
		return jw.NewResult(failure.New(err), 0)
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
