package utils

import (
	"fmt"
	"log"
	"os/exec"
)

func RunShellScript(scriptsDir string, scriptName string) (string, error) {
	// check file exiests
	scriptPath := fmt.Sprintf("%s/%s.sh", scriptsDir, scriptName)
	fmt.Println(scriptPath) // 输出: Hello world!

	if !FileExists(scriptPath) {
		return "", fmt.Errorf("文件不存在 %v", scriptPath)
	}

	// 使用exec.Command执行命令
	// run shell script
	out, err := exec.Command("bash", scriptPath).CombinedOutput()

	log.Printf("output: %v", out)

	return string(out), err
}
