package main

import (
	"sync"

	"github.com/mattermost/mattermost-server/v5/plugin"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

// OnActivate initialize the cowsay command
func (p *Plugin) OnActivate() error {
	if err := p.API.RegisterCommand(createCowsayCommand()); err != nil {
		return err
	}
	return nil
}
