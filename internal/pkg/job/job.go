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

package job

import (
	"github.com/arnumina/armen.core/pkg/jw"
	"github.com/arnumina/logger"

	"github.com/arnumina/armen.example/internal/pkg/container"
)

type (
	// Job AFAIRE.
	Job struct {
		*jw.Job
		Logger *logger.Logger
		Ctn    container.Container
	}
)

// New AFAIRE.
func New(job *jw.Job, logger *logger.Logger, ctn container.Container) *Job {
	return &Job{
		Job:    job,
		Logger: logger,
		Ctn:    ctn,
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
