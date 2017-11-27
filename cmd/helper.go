package cmd

import (
	"fmt"
	"strings"
)

func splitRepoTag(args ...string) (string, string, error) {
	if len(args) != 1 {
		return "", "", fmt.Errorf("Missing Repo:Tag")
	}
	splitparts := strings.Split(args[0], ":")
	if len(splitparts) < 2 {
		return "", "", fmt.Errorf("Wrong Format: Repo:Tag")
	}
	tag := splitparts[len(splitparts)-1]
	reposlice := splitparts[0 : len(splitparts)-1]
	repo := strings.Join(reposlice, ":")
	return repo, tag, nil
}
