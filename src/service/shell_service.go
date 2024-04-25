package service

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/mouday/cron-runner-shell/src/config"
	"github.com/mouday/cron-runner-shell/src/utils"
)

// 等待队列
var TaskWaitChannel = make(chan string, 16)

// 去重
var TaskMap sync.Map

// https://www.cnblogs.com/-wenli/p/12026413.html
func GetScriptPath(scriptName string) string {
	return filepath.Join(config.GetScriptDir(), scriptName+".sh")
}

func CheckScriptExists(scriptName string) bool {
	scriptPath := GetScriptPath(scriptName)
	log.Printf("Check scriptPath: %v", scriptPath)
	return utils.FileExists(scriptPath)
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

	// 启动
	if err := cmd.Start(); err != nil {
		log.Printf("Error Start: %s", err.Error())

		return
	}

	go asyncLog(stdout)
	go asyncLog(stderr)

	// 结束
	if err := cmd.Wait(); err != nil {
		log.Printf("Error waiting: %s", err.Error())
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

		log.Print(line)
	}
}

func Consumer() {
	for {
		name, ok := <-TaskWaitChannel
		if ok {
			// 队列中没有了就可以
			TaskMap.Delete(name)
			RunShellScript(name)
		} else {
			break
		}
	}

	fmt.Println("consumer done")
}

func AppendTask(name string) {
	// 队列去重，合并等待任务
	_, loaded := TaskMap.LoadOrStore(name, true)

	if !loaded {
		TaskWaitChannel <- name
	} else {
		log.Printf("script already in queue: %v", name)
	}
}
