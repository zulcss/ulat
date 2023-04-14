package compose

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/zulcss/ulat/pkg/constants"
	"github.com/zulcss/ulat/pkg/utils"
)

type Bootstrap struct {
	Opts []string
}

func NewBootstrapContext() *Bootstrap {
	return &Bootstrap{}
}

func BootstrapDebian(c *Compose) error {
	log.Println("Creating rootfs...")
	b := NewBootstrapContext()

	// Dont overwrite an existing target
	if exists := utils.Exists(c.Config.Target); exists {
		log.Fatalf("Target file exists, please remove before creating a new rootfs.")
	}

	// Check for a valid suite
	if _, ok := constants.GetSuite()[c.Config.Suite]; ok {
		log.Printf("Found valid suite: %s", c.Config.Suite)
	} else {
		log.Fatalf("Invalid suite: %s", c.Config.Suite)
	}

	// Check for a valid package set
	if constants.GetPackageSet(c.Config.Variant) {
		log.Printf("Valid package set found: %s", c.Config.Variant)
		b.Opts = append(b.Opts, fmt.Sprintf("--variant=%s", c.Config.Variant))
	}

	b.Opts = append(b.Opts, fmt.Sprintf("--format=%s", c.Format))

	// Add additional packages
	if len(c.Config.Packages) != 0 {
		for _, pkg := range c.Config.Packages {
			b.Opts = append(b.Opts, "--include", pkg)
		}
	}

	// Add additional apt options
	if len(c.Config.AptOpt) != 0 {
		for _, apt := range c.Config.AptOpt {
			b.Opts = append(b.Opts, "--aptopt", apt)
		}
	}

	// Add additonal dpkg options
	if len(c.Config.DpkgOpt) != 0 {
		for _, dpkg := range c.Config.DpkgOpt {
			b.Opts = append(b.Opts, "--dpkgopt", dpkg)
		}
	}
	// Additinal Components
	if len(c.Config.Components) != 0 {
		for _, component := range c.Config.Components {
			b.Opts = append(b.Opts, "--components", component)
		}
	}
	// Add setup-hook
	if len(c.Config.SetupHook) != 0 {
		for _, setupHook := range c.Config.SetupHook {
			b.Opts = append(b.Opts, "--setup-hook", setupHook)
		}
	}

	// Add extract-hook
	if len(c.Config.ExtractHook) != 0 {
		for _, extractHook := range c.Config.ExtractHook {
			b.Opts = append(b.Opts, "--extract-hook", extractHook)
		}
	}

	// Add essential-hook
	if len(c.Config.EssentialHook) != 0 {
		for _, essentialHook := range c.Config.EssentialHook {
			b.Opts = append(b.Opts, "--extract-hook", essentialHook)
		}
	}

	// Add customize-hook
	if len(c.Config.CustomizeHook) != 0 {
		for _, customizeHook := range c.Config.CustomizeHook {
			b.Opts = append(b.Opts, "--extract-hook", customizeHook)
		}
	}

	// Add hook-directory
	if c.Config.HookDirectory != "" {
		b.Opts = append(b.Opts, "--hook-directory", c.Config.HookDirectory)
	}

	b.Opts = append(b.Opts, c.Config.Suite,
		c.Config.Target, c.Config.Mirror)

	//Run the mmdebstrap command
	err := b.RunBootstrap()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (b *Bootstrap) RunBootstrap() error {
	log.Println("Running mmdebstrap...")

	cmd := exec.Command("mmdebstrap", b.Opts...)
	err := utils.Run(cmd)
	if err != nil {
		return fmt.Errorf("Failed to run mmdebstrap: %v", err)
	}
	return nil
}
