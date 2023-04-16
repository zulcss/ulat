package compose

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/zulcss/ulat/pkg/utils"
)

type Ostree struct {
	RootfsDir string
}

func NewOstreeContext(rootfs string) *Ostree {
	return &Ostree{
		RootfsDir: rootfs,
	}
}

func OstreeCreate(c *Compose) error {
	log.Println("Creating ostree repo...")

	rootfs, err := os.MkdirTemp(c.Workspace, "_rootfs")
	if err != nil {
		return fmt.Errorf("Failed to create rootfs directory: %v", err)
	}
	o := NewOstreeContext(rootfs)

	err = o.Unpack(c)
	if err != nil {
		return fmt.Errorf("Failed to unpack tarball: %v", err)
	}

	err = o.BuildBranch(c)
	if err != nil {
		return fmt.Errorf("Failed to build ostree branch: %v", err)
	}
	return nil
}

func (o *Ostree) Unpack(c *Compose) error {
	log.Println("Unpacking tarball...")
	log.Println(c.Config.Target)
	var opts []string

	if c.Verbose {
		opts = append(opts, "zxvf")
	} else {
		opts = append(opts, "zxf")
	}
	opts = append(
		opts,
		fmt.Sprintf("%s", c.Config.Target),
		"--numeric-owner",
		"--exclude=./dev",
		"-C",
		fmt.Sprintf("%s", o.RootfsDir),
	)
	cmd := exec.Command("tar", opts...)
	err := utils.Run(cmd)
	if err != nil {
		return fmt.Errorf("Failed to unpack rootfs: %v", err)
	}
	return nil
}

func (o *Ostree) BuildBranch(c *Compose) error {
	log.Println("Building ostree branch...")

	_, err := utils.SH(
		fmt.Sprintf("ostree --repo=%s commit --branch=%s %s", c.Config.Repo, c.Config.Branch, o.RootfsDir))
	if err != nil {
		return err
	}
	return nil
}
