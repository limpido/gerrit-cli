package utils

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Execute(cmd string) string {
	cmd_arr := strings.Split(cmd, " ")
	name := cmd_arr[0]
	args := cmd_arr[1:]
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out[:])
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\n")
	return s
}

func GitHead() string {
	return Execute("git rev-parse HEAD")
}

func GitBranch() string {
	return Execute("git rev-parse --abbrev-ref HEAD")
}

func GitOrigin() string {
	return Execute("git remote show")
}

func GitRepoUrl() string {
	origin := GitOrigin()
	s := fmt.Sprint("git remote get-url ", origin)
	return Execute(s)
}

// func GitUpstream() string {
// 	s := Execute("git rev-parse --abbrev-ref @{u}")

// }

func GitServer() string {
	return ""
}
