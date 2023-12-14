package main

import (
	"net/http"

	"github.com/hashicorp/go-hclog"
	"github.com/reeveci/reeve-lib/plugin"
	"github.com/reeveci/reeve-lib/schema"
)

const PLUGIN_NAME = "consul"

func main() {
	log := hclog.New(&hclog.LoggerOptions{})

	plugin.Serve(&plugin.PluginConfig{
		Plugin: &ConsulPlugin{
			Log: log,

			http: &http.Client{},
		},

		Logger: log,
	})
}

type ConsulPlugin struct {
	Url       string
	Token     string
	KeyPrefix string
	Priority  uint32
	Secret    bool

	Log hclog.Logger

	http *http.Client
}

func (p *ConsulPlugin) Name() (string, error) {
	return PLUGIN_NAME, nil
}

func (p *ConsulPlugin) Register(settings map[string]string, api plugin.ReeveAPI) (capabilities plugin.Capabilities, err error) {
	api.Close()

	var enabled bool
	if enabled, err = boolSetting(settings, "ENABLED"); !enabled || err != nil {
		return
	}
	if p.Url, err = requireSetting(settings, "URL"); err != nil {
		return
	}
	if p.Token, err = requireSetting(settings, "TOKEN"); err != nil {
		return
	}
	p.KeyPrefix = settings["KEY_PREFIX"]
	var priority int
	if priority, err = intSetting(settings, "PRIORITY", 1); err != nil {
		return
	} else {
		p.Priority = uint32(priority)
	}
	if p.Secret, err = boolSetting(settings, "SECRET"); err != nil {
		return
	}

	capabilities.Resolve = true
	return
}

func (p *ConsulPlugin) Unregister() error {
	return nil
}

func (p *ConsulPlugin) Message(source string, message schema.Message) error {
	return nil
}

func (p *ConsulPlugin) Discover(trigger schema.Trigger) ([]schema.Pipeline, error) {
	return nil, nil
}

func (p *ConsulPlugin) Notify(status schema.PipelineStatus) error {
	return nil
}

func (p *ConsulPlugin) CLIMethod(method string, args []string) (string, error) {
	return "", nil
}
