package utils

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/mouday/cron-runner-shell/src/config"
)

// https://www.cnblogs.com/-wenli/p/12026413.html
func GetScriptPath(scriptName string) string {
	return filepath.Join(config.GetScriptDir(), scriptName+".sh")
}

func CheckScriptExists(scriptName string) bool {
	scriptPath := GetScriptPath(scriptName)
	log.Printf("Check scriptPath: %v", scriptPath)
	return FileExists(scriptPath)
}

func RunShellScript(scriptName string) {
	// check file exiests

	scriptPath := GetScriptPath(scriptName)

	log.Printf("scriptPath: %v", scriptPath)

	// 二次检查
	if !CheckScriptExists(scriptName) {
		log.Printf("file not found: %v", scriptPath)
		return
	}

	// 使用exec.Command执行命令
	// run shell script
	cmd := exec.Command("bash", scriptPath)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		log.Printf("Error Start: %s", err.Error())
		return
	}

	go asyncLog(stdout)
	go asyncLog(stderr)

	if err := cmd.Wait(); err != nil {
		log.Printf("Error waiting: %s", err.Error())
		return
	}
}

// ref: https://blog.csdn.net/xuezhangjun0121/article/details/135284214
// https://blog.csdn.net/flyfreelyit/article/details/103697013
func asyncLog(std io.ReadCloser) {
	reader := bufio.NewReader(std)

	for {
		line, err := reader.ReadString('\n')

		if err != nil || io.EOF == err {
			break
		}

		log.Printf("out: %s", line)
	}
}
