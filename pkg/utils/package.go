package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetModuleName() (string, error) {
	cmd := exec.Command("go", "list", "-m")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get module name: %v", err)
	}
	return strings.TrimSpace(out.String()), nil
}

func GetModulePath() (string, error) {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	var out bytes.Buffer
	cmd.Stdout = &out

	// 执行命令并捕获错误
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get project directory: %v", err)
	}

	// 去除换行符，返回结果
	return strings.TrimSpace(out.String()), nil
}

func GetModuleFullPath(module string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	moduleName, err := GetModuleName()
	if err != nil {
		return "", err
	}
	modulePath, err := GetModulePath()
	if err != nil {
		return "", err
	}
	pathArr := strings.Split(currentDir, modulePath)
	var actualPath string
	if len(pathArr) < 2 {
		actualPath = ""
	} else {
		actualPath = pathArr[1]
	}
	res := moduleName
	if actualPath != "" {
		if strings.HasPrefix(actualPath, "/") {
			res += actualPath
		} else {
			res += "/" + actualPath
		}
	}
	if strings.HasPrefix(module, "/") {
		res += module
	} else {
		res += "/" + module
	}
	if strings.HasSuffix(res, "/") {
		res = strings.TrimSuffix(res, "/")
	}
	return res, nil
}
