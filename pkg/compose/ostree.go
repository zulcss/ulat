package compose

import (
	"log"
)

type Ostree struct {
	RootfsDir string
}

func OstreeCreate(c *Compose) error {
	log.Println("Creating ostree repo...")

	return nil
}

func (o *Ostree) Unpack(c *Compose) error {
	log.Println("Unpacking tarball...")

	return nil
}

func (o *Ostree) BuildBranch(c *Compose) error {
	log.Println("Building ostree branch...")

	return nil
}
