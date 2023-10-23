package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	cliName     string
	cliLocation string
	cliOS       string
	cliArch     string
	cliCommand  string
	cliAll      bool
	cliMain     bool
	cliHelp     bool
	cliVersion  bool
	toolbox     = []string{"go"}
)

func init() {
	flag.StringVar(&cliName, "n", "", "set static file name")
	flag.StringVar(&cliLocation, "d", "bin", "set static file output location")
	flag.StringVar(&cliOS, "os", "", "specify os")
	flag.StringVar(&cliArch, "arch", "", "specify architecture")
	flag.StringVar(&cliCommand, "cmd", "", "use custom build commands")
	flag.BoolVar(&cliAll, "all", false, "build all supported os and architecture")
	flag.BoolVar(&cliMain, "main", false, "build all supported architecture for windows, macos, linux and freebsd")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
    gobuilder [Options] [Commands]

Options:
    -n      name             : set static file name
    -d      location         : set static file output location
    -os     operatingSystem  : specify os
    -arch   architecture     : specify architecture
    -cmd    commands         : use custom build commands
    -all                     : build all supported os and architecture
    -main                    : build all supported architecture for windows, macos, linux and freebsd

Command:
    -h                       : Show help
    -v                       : Show version

Example:
    0) gobuilder                          : Built using the operating system and architecture of the current system
    1) gobuilder -n myapp -d bin          : Build with the name myapp and put the output files into the bin folder
    2) gobuilder -os windows              : Build all supported architectures for Windows
    3) gobuilder -arch amd64              : Build all supported operating systems for amd64
    4) gobuilder -os windows -arch amd64  : Build use Windows and amd64
    5) gobuilder -all                     : Build all supported os and architecture
    6) gobuilder -main                    : Build all supported architecture for windows, macos, linux and freebsd
    7) gobuilder -v                       : Show version
    8) gobuilder -h                       : Show help

More Information:
    1) Visit https://golang.org/doc/install/source#environment for more information on supported operating system and architecture
    2) You may need to install gcc to build some special os and arch pairs,such as android/386 and android/amd64 and android/arm`
		fmt.Println(helpInfo)
	}

	// 打印用法
	if cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println("v1.05")
		os.Exit(0)
	}

	// 检查依赖
	for _, tool := range toolbox {
		_, err := exec.LookPath(tool)
		if err != nil {
			log.Panicf("install %s before running", tool)
		}
	}
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release
  1.01:
    - Update os/arch list
    - Support ARCHs add riscv64 
    - Support OSs add aix
  1.02:
    - Add FreeBSD to the main building support list
  1.03:
    - Use "go mod edit -json" to get the package name instead of "go mod graph" to avoid errors when the package has no modules
  1.04:
    - Removed software dependencies for easier use
  1.05:
    - Support custom build command`
	fmt.Println(versionInfo)
}

func main() {
	// 指定编译全部
	if cliAll {
		err := buildAll(cliLocation, cliName)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
	// 指定编译主要的
	if cliMain {
		err := buildMain(cliLocation, cliName)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
	// 未指定或者指定os/arch自定义编译
	err := buildCustom(cliLocation, cliName)
	if err != nil {
		log.Fatal(err)
	}
}
