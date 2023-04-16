package compose

import (
	"github.com/spf13/viper"
)

type ComposeConfig struct {
	// Debian release code name (eg: bullseye, bookworm), or a symbolic name (unstable, testing, etc)
	Suite string
	// Name of the tarball artifact
	Target string
	// Package set to install
	Variant string
	// Mirror to use
	Mirror string
	// Extra packages to install
	Packages []string
	// Pass options to apt
	AptOpt []string
	// Pass options to dpkg
	DpkgOpt []string
	// Debian components (eg: main, contrib, non-free)
	Components []string
	// Execute commands right after initial setup (directory creation, configuration of apt, etc) but
	// before any packages are downloaded or installed.
	SetupHook []string
	// Extract commands after the "Essential:yes" apckages have been extracted but before installing them"
	ExtractHook []string
	// Execute arbitrary commands after the "Essential:yes" packages have been installed but before installing
	// the remaining packages.
	EssentialHook []string
	// Execute arbitrary commands after the chroot is setup and all the packages are installed
	CustomizeHook []string
	// Ensure scripts in directory with filenames starting with "setup", "extract", "essential", or "customize"
	HookDirectory []string
	// OStree Repo
	Repo string
	// Ostree archive type
	Mode string
	// Ostree branch name
	Branch string
}

func LoadConfig(c *Compose) ComposeConfig {
	viper.SetConfigFile(c.ConfigFile)

	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	cfg := ComposeConfig{
		// common configuration
		Target: viper.GetString("target"),
		// mmdebstrap configuration
		Suite:         viper.GetString("rootfs.suite"),
		Variant:       viper.GetString("rootfs.variant"),
		Mirror:        viper.GetString("rootfs.mirror"),
		Packages:      viper.GetStringSlice("rootfs.packages"),
		AptOpt:        viper.GetStringSlice("rootfs.aptOpt"),
		DpkgOpt:       viper.GetStringSlice("rootfs.dpkgOpt"),
		Components:    viper.GetStringSlice("rootfs.components"),
		SetupHook:     viper.GetStringSlice("rootfs.setup-hook"),
		ExtractHook:   viper.GetStringSlice("rootfs.extract-hook"),
		EssentialHook: viper.GetStringSlice("rootfs.essential-hook"),
		CustomizeHook: viper.GetStringSlice("rootfs.customize-hook"),
		HookDirectory: viper.GetStringSlice("rootfs.hook-directory"),

		// Ostree
		Repo:   viper.GetString("ostree.repo"),
		Mode:   viper.GetString("ostree.mode"),
		Branch: viper.GetString("ostree.branch"),
	}
	return cfg
}
