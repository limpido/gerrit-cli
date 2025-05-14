package util

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
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

func LogFatal(msg string) {
	log.Fatalln("\033[91;1m", msg, "\033[0m")
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

func GitServer() string {
	repoUrl := GitRepoUrl()
	re := regexp.MustCompile("(.+://.+?):.+")
	match := re.FindSubmatch([]byte(repoUrl))
	if match == nil {
		LogFatal("error: unable to parse repo url")
	}

	return string(match[1])
}

func GitUpstream() string {
	s := Execute("git rev-parse --abbrev-ref @{u}")
	re := regexp.MustCompile("(.+?)/(.+)")
	match := re.FindSubmatch([]byte(s))
	if match == nil {
		LogFatal("error: unable to parse upstream")
	}

	upstream := match[2]
	return string(upstream)
}

func Query(commit string) string {
	server := GitServer()
	cmd := fmt.Sprintf("ssh %s gerrit query %s --current-patch-set --format JSON", server, commit)
	res := Execute(cmd)
	strs := strings.Split(res, "\n")
	if len(strs) < 2 {
		LogFatal(fmt.Sprintf("error: unable to find commit %s", commit))
	}
	return strings.Split(res, "\n")[0]
}
