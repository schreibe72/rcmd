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

func getServerCredential(connectionstring string) (server string, username string, password string) {
	parts := strings.SplitN(connectionstring, "/", 2)
	server = parts[0]
	username = Servers[server].Username
	password = Servers[server].Password
	return
}

func getRepo(connectionstring string) (repo string) {
	parts := strings.SplitN(connectionstring, "/", 2)
	repo = parts[1]
	return
}
