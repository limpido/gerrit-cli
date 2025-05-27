package util

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"os"
)

func Execute(cmd string) {
	cmd_arr := strings.Split(cmd, " ")
	name := cmd_arr[0]
	args := cmd_arr[1:]
	c := exec.Command(name, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}

func ExecuteAndReturnOutput(cmd string) string {
	cmd_arr := strings.Split(cmd, " ")
	name := cmd_arr[0]
	args := cmd_arr[1:]

	out, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		LogFatal(string(out[:]))
	}
	s := string(out[:])
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\n")
	return s
}

func LogFatal(msg string) {
	log.SetFlags(0)
	log.Fatal("\033[91;1m", msg, "\033[0m")
}

func GitHead() string {
	return ExecuteAndReturnOutput("git rev-parse HEAD")
}

func GitBranch() string {
	return ExecuteAndReturnOutput("git rev-parse --abbrev-ref HEAD")
}

func GitOrigin() string {
	return ExecuteAndReturnOutput("git remote show")
}

func GitRepoUrl() string {
	origin := GitOrigin()
	s := fmt.Sprint("git remote get-url ", origin)
	return ExecuteAndReturnOutput(s)
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
	s := ExecuteAndReturnOutput("git rev-parse --abbrev-ref @{u}")
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
	res := ExecuteAndReturnOutput(cmd)
	strs := strings.Split(res, "\n")
	if len(strs) < 2 {
		LogFatal(fmt.Sprintf("error: unable to find commit %s", commit))
	}
	return strings.Split(res, "\n")[0]
}
