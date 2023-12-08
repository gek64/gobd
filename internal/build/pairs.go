package build

import (
	"runtime"
)

type Pair struct {
	OS   string
	ARCH string
}

var Pairs = []Pair{
	{OS: "aix", ARCH: "ppc64"},
	{OS: "android", ARCH: "386"},
	{OS: "android", ARCH: "amd64"},
	{OS: "android", ARCH: "arm"},
	{OS: "android", ARCH: "arm64"},
	{OS: "darwin", ARCH: "amd64"},
	{OS: "darwin", ARCH: "arm64"},
	{OS: "dragonfly", ARCH: "amd64"},
	{OS: "freebsd", ARCH: "386"},
	{OS: "freebsd", ARCH: "amd64"},
	{OS: "freebsd", ARCH: "arm"},
	{OS: "illumos", ARCH: "amd64"},
	{OS: "ios", ARCH: "arm64"},
	{OS: "js", ARCH: "wasm"},
	{OS: "linux", ARCH: "386"},
	{OS: "linux", ARCH: "amd64"},
	{OS: "linux", ARCH: "arm"},
	{OS: "linux", ARCH: "arm64"},
	{OS: "linux", ARCH: "loong64"},
	{OS: "linux", ARCH: "mips"},
	{OS: "linux", ARCH: "mipsle"},
	{OS: "linux", ARCH: "mips64"},
	{OS: "linux", ARCH: "mips64le"},
	{OS: "linux", ARCH: "ppc64"},
	{OS: "linux", ARCH: "ppc64le"},
	{OS: "linux", ARCH: "riscv64"},
	{OS: "linux", ARCH: "s390x"},
	{OS: "netbsd", ARCH: "386"},
	{OS: "netbsd", ARCH: "amd64"},
	{OS: "netbsd", ARCH: "arm"},
	{OS: "openbsd", ARCH: "386"},
	{OS: "openbsd", ARCH: "amd64"},
	{OS: "openbsd", ARCH: "arm"},
	{OS: "openbsd", ARCH: "arm64"},
	{OS: "plan9", ARCH: "386"},
	{OS: "plan9", ARCH: "amd64"},
	{OS: "plan9", ARCH: "arm"},
	{OS: "solaris", ARCH: "amd64"},
	{OS: "wasip1", ARCH: "wasm"},
	{OS: "windows", ARCH: "386"},
	{OS: "windows", ARCH: "amd64"},
	{OS: "windows", ARCH: "arm"},
	{OS: "windows", ARCH: "arm64"},
}

func GetAllPairs() []Pair {
	return Pairs
}

func GetMainPairs() (p []Pair) {
	var goos = []string{"darwin", "freebsd", "linux", "windows"}
	var goarch = []string{"amd64", "arm64", "386", "arm"}

	for _, pair := range Pairs {
		if in(goos, pair.OS) && in(goarch, pair.ARCH) {
			p = append(p, pair)
		}
	}
	return p
}

func GetSelectedPairs(OS string, ARCH string) (p []Pair) {
	if OS != "" && ARCH != "" {
		for _, pair := range Pairs {
			if in([]string{OS}, pair.OS) && in([]string{ARCH}, pair.ARCH) {
				p = append(p, pair)
			}
		}
		return p
	}

	if OS == "" && ARCH != "" {
		for _, pair := range Pairs {
			if in([]string{ARCH}, pair.ARCH) {
				p = append(p, pair)
			}
		}
		return p
	}

	if OS != "" && ARCH == "" {
		for _, pair := range Pairs {
			if in([]string{OS}, pair.OS) {
				p = append(p, pair)
			}
		}
		return p
	}

	// OS == "" && ARCH == ""
	for _, pair := range Pairs {
		if in([]string{runtime.GOOS}, pair.OS) && in([]string{runtime.GOARCH}, pair.ARCH) {
			p = append(p, pair)
		}
	}
	return p
}

func in(ss []string, t string) bool {
	for _, s := range ss {
		if t == s {
			return true
		}
	}
	return false
}
