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
	"github.com/arnumina/armen.core/pkg/resources"

	"github.com/arnumina/armen.example/internal/pkg/container"
	"github.com/arnumina/armen.example/internal/pkg/jw"
	"github.com/arnumina/armen.example/internal/pkg/messages"
	"github.com/arnumina/armen.example/internal/pkg/plugin"
)

type (
	// Example AFAIRE.
	Example struct {
		*plugin.Plugin
		cr *container.Resources
		mh *messages.Handler
	}
)

// New AFAIRE.
func New(version, builtAt string, resources resources.Container) *Example {
	util := resources.Util()
	plugin := plugin.New("example", version, util.UnixToTime(builtAt))
	logger := resources.Logger().Clone(util.LoggerPrefix(plugin.Name(), plugin.ID()))

	cr := container.New()
	cr.SetUtil(util)
	cr.SetPlugin(plugin)
	cr.SetLogger(logger)
	cr.SetModel(resources.Model())
	cr.SetBus(resources.Bus())
	cr.SetJWHandler(jw.NewHandler(cr))

	return &Example{
		Plugin: plugin,
		cr:     cr,
		mh:     messages.NewHandler(cr),
	}
}

// Build AFAIRE.
func (e *Example) Build() error {
	return e.mh.Subscribe(e.cr)
}

// Close AFAIRE.
func (e *Example) Close() {}

/*
######################################################################################################## @(°_°)@ #######
*/
