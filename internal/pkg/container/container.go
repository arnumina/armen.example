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

package container

import (
	"time"

	"github.com/arnumina/armen.core/pkg/jw"
	"github.com/arnumina/armen.core/pkg/message"
	"github.com/arnumina/armen.core/pkg/resources"
	"github.com/arnumina/logger"
)

type (
	// Bus AFAIRE.
	Bus interface {
		resources.Bus
	}

	// JWHandler AFAIRE.
	JWHandler interface {
		CreateJobTest(msg *message.Message)
		InsertWorkflow(wf *jw.Workflow) error
	}

	// Model AFAIRE.
	Model interface {
		resources.Model
	}

	// Plugin AFAIRE.
	Plugin interface {
		ID() string
		Name() string
		Version() string
		BuiltAt() time.Time
	}

	// Util AFAIRE.
	Util interface {
		resources.Util
	}

	// Container AFAIRE.
	Container interface {
		Bus() Bus
		JWHandler() JWHandler
		Logger() *logger.Logger
		Model() Model
		Plugin() Plugin
		Util() Util
	}

	// Resources AFAIRE.
	Resources struct {
		bus       Bus
		jwHandler JWHandler
		logger    *logger.Logger
		model     Model
		plugin    Plugin
		util      Util
	}
)

// New AFAIRE.
func New() *Resources {
	return &Resources{}
}

// Bus AFAIRE.
func (r *Resources) Bus() Bus {
	return r.bus
}

// SetBus AFAIRE.
func (r *Resources) SetBus(bus Bus) {
	r.bus = bus
}

// JWHandler AFAIRE.
func (r *Resources) JWHandler() JWHandler {
	return r.jwHandler
}

// SetJWHandler AFAIRE.
func (r *Resources) SetJWHandler(jwHandler JWHandler) {
	r.jwHandler = jwHandler
}

// Logger AFAIRE.
func (r *Resources) Logger() *logger.Logger {
	return r.logger
}

// SetLogger AFAIRE.
func (r *Resources) SetLogger(logger *logger.Logger) {
	r.logger = logger
}

// Model AFAIRE.
func (r *Resources) Model() Model {
	return r.model
}

// SetModel AFAIRE.
func (r *Resources) SetModel(model Model) {
	r.model = model
}

// Plugin AFAIRE.
func (r *Resources) Plugin() Plugin {
	return r.plugin
}

// SetPlugin AFAIRE.
func (r *Resources) SetPlugin(plugin Plugin) {
	r.plugin = plugin
}

// Util AFAIRE.
func (r *Resources) Util() Util {
	return r.util
}

// SetUtil AFAIRE.
func (r *Resources) SetUtil(util Util) {
	r.util = util
}

/*
######################################################################################################## @(°_°)@ #######
*/
