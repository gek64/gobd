package main

import (
	"fmt"
	"gek_exec"
	"os"
	"os/exec"
	"strings"
)

var (
	osList   = []string{"android", "darwin", "dragonfly", "freebsd", "illumos", "ios", "js", "linux", "netbsd", "openbsd", "plan9", "solaris", "windows"}
	archList = []string{"amd64", "386", "arm", "arm64", "ppc64le", "mips64le", "mips64", "mipsle", "mips", "s390x", "wasm"}
)

// getModuleName 获取模块名称
func getModuleName() (name string, err error) {
	// 使用go mod 列出所有的依赖,依赖列表中会包含当前的包名
	output, err := gek_exec.Output("go mod graph")
	if err != nil {
		return "", err
	}
	// 按空格分词
	modInfo := strings.Fields(output)
	// 检查包名称中是否含有错误的字符,同时也能检查包名是否为链接
	if strings.ContainsAny(modInfo[0], "\\/:*?\"<>|") {
		return "", fmt.Errorf("%s is an invalid name", modInfo[0])
	}

	return modInfo[0], nil
}

// getStaticName 获取编译后的静态文件名,customName为自定义名称(可选)
func getStaticName(targetOS string, targetARCH string, customName ...interface{}) (name string, err error) {
	// 如果使用自定义名称
	if len(customName) > 0 {
		name = fmt.Sprintf("%s-%s-%s", customName[0], targetOS, targetARCH)
	} else {
		// 不使用自定义名称,则获取模块名称
		packageName, err := getModuleName()
		if err != nil {
			return "", err
		}
		name = fmt.Sprintf("%s-%s-%s", packageName, targetOS, targetARCH)
	}

	// windows 系统则在静态文件名中加入 .exe 后缀
	if targetOS == "windows" {
		name = name + ".exe"
	}
	return name, nil
}

// build 编译指定的名称,操作系统,系统架构的应用
func build(name string, targetOS string, targetARCH string) (err error) {
	// 指定编译目标操作系统
	OS := fmt.Sprintf("GOOS=%s", targetOS)
	// 指定编译目标系统架构
	ARCH := fmt.Sprintf("GOARCH=%s", targetARCH)
	// 编译命令
	cmd := exec.Command("go", "build", "-v", "-trimpath", "-ldflags", "-s -w", "-o", name)
	// 指定环境
	cmd.Env = append(os.Environ(), OS, ARCH)

	// 运行程序
	err = gek_exec.Run(cmd)
	if err != nil {
		return err
	}
	return nil
}
