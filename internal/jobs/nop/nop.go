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

package nop

import (
	"time"

	"github.com/arnumina/armen.core/pkg/jw"

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
	j.Data[j.ID] = time.Now()
	j.Data["_"+j.Name] = j.ID

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
