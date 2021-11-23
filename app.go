package main

import (
	"fmt"
	"gek_exec"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	defaultOSList   = []string{"android", "darwin", "dragonfly", "freebsd", "illumos", "ios", "js", "linux", "netbsd", "openbsd", "plan9", "solaris", "windows"}
	defaultArchList = []string{"amd64", "386", "arm", "arm64", "ppc64le", "mips64le", "mips64", "mipsle", "mips", "s390x", "wasm"}
	mainOSList      = []string{"darwin", "linux", "windows"}
	mainArchList    = []string{"amd64", "386", "arm", "arm64"}
	customOSList    []string
	customArchList  []string
)

// 基本功能函数
// 用于实现基本功能
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
func getStaticName(targetOS string, targetARCH string, customName ...string) (name string, err error) {
	// 如果使用自定义名称
	if len(customName) > 0 && customName[0] != "" {
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
func build(name string, location string, targetOS string, targetARCH string) (err error) {
	// 指定编译目标操作系统
	OS := fmt.Sprintf("GOOS=%s", targetOS)
	// 指定编译目标系统架构
	ARCH := fmt.Sprintf("GOARCH=%s", targetARCH)
	// 编译命令
	cmd := exec.Command("go", "build", "-v", "-trimpath", "-ldflags", "-s -w", "-o", filepath.Join(location, name))
	// 指定环境
	cmd.Env = append(os.Environ(), OS, ARCH)

	// 运行程序
	err = gek_exec.Run(cmd)
	if err != nil {
		return err
	}
	return nil
}

// buildList 按给定的 操作系统/架构列表对 编译,osList为操作系统列表,archList为系统架构,location为存储静态文件的地址,customName为自定义静态文件名(可选)
func buildList(osList []string, archList []string, location string, customName ...string) (err error) {
	// 双重循环找出操作系统/架构列表对
	for _, OS := range osList {
		for _, ARCH := range archList {
			var name = ""
			if len(customName) > 0 && customName[0] != "" {
				// 有自定义静态文件名,直接使用自定义静态文件名
				name, err = getStaticName(OS, ARCH, customName...)
				if err != nil {
					return err
				}
			} else {
				// 无自定义静态文件名,使用应用模块名
				name, err = getStaticName(OS, ARCH)
				if err != nil {
					return err
				}
			}

			// 编译
			err = build(name, location, OS, ARCH)
			if err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

// 流程函数
// 用于实现命令行工具的处理流程
// buildAll 编译所有支持的操作系统和系统架构对的静态文件,location为存储静态文件的地址,customName为自定义静态文件名(可选)
func buildAll(location string, customName ...string) (err error) {
	err = buildList(defaultOSList, defaultArchList, location, customName...)
	if err != nil {
		return err
	}
	return nil
}

// buildMain 编译主要的操作系统和系统架构对的静态文件,location为存储静态文件的地址,customName为自定义静态文件名(可选)
func buildMain(location string, customName ...string) (err error) {
	err = buildList(mainOSList, mainArchList, location, customName...)
	if err != nil {
		return err
	}
	return nil
}

// buildCustom 编译自定义的操作系统和系统架构对的静态文件,location为存储静态文件的地址,customName为自定义静态文件名(可选),如果控制台未指定编译目标,则默使用当前系统的操作系统与系统架构
func buildCustom(location string, customName ...string) (err error) {
	// 控制台未指定编译目标,填充为当前系统的操作系统与系统架构
	if cliOS == "" && cliArch == "" {
		cliOS = runtime.GOOS
		cliArch = runtime.GOARCH
	}
	// 控制台os参数不为空,则添加进自定义os列表,为空则添加默认列表
	if cliOS != "" {
		customOSList = append(customArchList, cliOS)
	} else {
		customOSList = append(customOSList, defaultOSList...)
	}
	// 控制台arch参数不为空,则添加进自定义arch列表,为空则添加默认列表
	if cliArch != "" {
		customArchList = append(customArchList, cliArch)
	} else {
		customArchList = append(customArchList, defaultArchList...)
	}

	// 自定义编译
	err = buildList(customOSList, customArchList, location, customName...)
	if err != nil {
		return err
	}
	return nil
}
