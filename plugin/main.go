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

package main

import (
	"github.com/arnumina/armen.core/pkg/plugin"
	"github.com/arnumina/armen.core/pkg/resources"

	"github.com/arnumina/armen.example/internal/example"
)

var (
	_version string
	_builtAt string
)

// Export AFAIRE.
func Export(resources resources.Container) plugin.Plugin {
	return example.New(_version, _builtAt, resources)
}

func main() {
	_ = Export(nil) // avoid linter errors /////////////////////////////////////////////////////////////////////////////
}

/*
######################################################################################################## @(°_°)@ #######
*/
