package build

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type Mod struct {
	Module struct {
		Path string `json:"Path"`
	} `json:"Module"`
}

// GetModuleName 获取模块名称
func GetModuleName() (name string, err error) {
	var mod Mod

	// 使用 go mod edit -json 列出模块信息
	output, err := exec.Command("go", "mod", "edit", "-json").Output()
	if err != nil {
		return "", err
	}

	// json解码到mod
	err = json.Unmarshal([]byte(output), &mod)
	if err != nil {
		return "", err
	}

	// mod中提取出模块名称
	modName := mod.Module.Path
	// 检查模块名称中是否含有错误的字符,同时也能检查包名是否为链接
	if strings.ContainsAny(modName, "\\/:*?\"<>|") {
		return "", fmt.Errorf("%s is an invalid name", modName)
	}

	return modName, nil
}
