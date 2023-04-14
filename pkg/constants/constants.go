package constants

import "log"

const (
	CacheDir = "/var/cache"
)

func GetSuite() map[string]string {
	return map[string]string{
		"bullseye": "11",
		"bookworm": "12",
	}
}

func GetPackageSet(PackageSet string) bool {
	PackageSets := []string{"extract", "custom", "essential", "apt", "required", "minbase", "buildd", "important", "deboostrap", "standard"}
	log.Printf("Checking %s for valid package set...", PackageSet)
	for _, pkg := range PackageSets {
		if PackageSet == pkg {
			return true
		}
	}
	return false
}
