package compose

import (
	"fmt"
)

type Compose struct {
	ConfigFile string
	Format     string
	Target     string
	Workspace  string
	Config     ComposeConfig
	Verbose    bool
}

func NewComposeContext(config string, workspace string, verbose bool) *Compose {
	return &Compose{
		ConfigFile: config,
		Workspace:  workspace,
		Format:     "tar",
		Verbose:    verbose,
	}
}

func (c *Compose) Bootstrap() error {
	err := BootstrapDebian(c)
	if err != nil {
		return fmt.Errorf("Bootstrap failed: %v", err)
	}
	return nil
}
