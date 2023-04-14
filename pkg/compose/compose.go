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
}

func NewComposeContext(config string, workspace string) *Compose {
	return &Compose{
		ConfigFile: config,
		Workspace:  workspace,
		Format:     "tar",
	}
}

func (c *Compose) Bootstrap() error {
	err := BootstrapDebian(c)
	if err != nil {
		return fmt.Errorf("Bootstrap failed: %v", err)
	}
	return nil
}
