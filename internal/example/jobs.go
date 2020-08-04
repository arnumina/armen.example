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

package example

import (
	"github.com/arnumina/armen.core/pkg/jw"
	"github.com/arnumina/failure"
	"github.com/arnumina/logger"

	"github.com/arnumina/armen.example/internal/jobs/nop"
	"github.com/arnumina/armen.example/internal/jobs/test"
	_job "github.com/arnumina/armen.example/internal/pkg/job"
)

// RunJob AFAIRE.
func (e *Example) RunJob(job *jw.Job, logger *logger.Logger) *jw.Result {
	var jwr jw.Runner

	ej := _job.New(job, logger, e.cr)

	switch job.Type {
	case "nop":
		jwr = nop.New(ej)
	case "test":
		jwr = test.New(ej)
	default:
		return jw.NewResult(
			failure.New(nil).Msg("the type of this job is not valid"), /////////////////////////////////////////////////
			0,
		)
	}

	return jwr.Run()
}

/*
######################################################################################################## @(°_°)@ #######
*/
