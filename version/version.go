package version

import (
	"fmt"

	"github.com/github/hub/git"
)

var Version = "2.12.6"

func FullVersion() (string, error) {
	gitVersion, err := git.Version()
	if err != nil {
		gitVersion = "git version (unavailable)"
	}
	return fmt.Sprintf("%s\nhub version %s", gitVersion, Version), err
}
