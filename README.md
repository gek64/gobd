# gobuilder
GO Builder
- A tool can help you easily build golang static files for multiple operating systems and architectures 
- Written in golang 

## Usage
```
Usage:
    gobuilder [Options] [Commands]

Options:
    -n      name             : set static file name
    -d      location         : set static file output location
    -os     operatingSystem  : specify os
    -arch   architecture     : specify architecture
    -all                     : build all supported os and architecture
    -main                    : build all supported architecture for windows macos linux

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
    6) gobuilder -main                    : Build all supported architecture for windows, macos and linux
    7) gobuilder -v                       : Show version
    8) gobuilder -h                       : Show help

More Information:
    1) Visit https://golang.org/doc/install/source#environment for more information on supported operating system and architecture
    2) You may need to install gcc to build some special os and arch pairs,such as android/386 and android/amd64 and android/arm

```

## Build
### Example
```sh
# dependence
git clone https://github.com/gek64/gek.git

git clone https://github.com/gek64/gobuilder.git

cd gobuilder

go build -v -trimpath -ldflags "-s -w"
```
