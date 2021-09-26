package version

import "fmt"

var (
	Version = "v0.2"
	BuildTime = "2021-09-26"
	GitHash = ""
	Author = "Gourds"
)

func GetVersion() string{
	return fmt.Sprintf("Version is %s", Version)
}