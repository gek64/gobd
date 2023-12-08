package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gobd/internal/build"
	"log"
	"os"
)

func main() {
	var build_all bool
	var build_main bool
	var build_no_debug bool
	var build_no_cgo bool
	var build_os string
	var build_arch string
	var build_output_directory string
	var build_output_name string
	var build_opts cli.StringSlice
	var build_envs cli.StringSlice

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:        "all",
			Usage:       "set build all supported os and architecture",
			Destination: &build_all,
		},
		&cli.BoolFlag{
			Name:        "main",
			Usage:       "set build all supported architecture for windows, macos, linux and freebsd",
			Destination: &build_main,
		},
		&cli.StringFlag{
			Name:        "os",
			Usage:       "set build operating system",
			Destination: &build_os,
		},
		&cli.StringFlag{
			Name:        "arch",
			Usage:       "set build architecture",
			Destination: &build_arch,
		},
		&cli.StringFlag{
			Name:        "name",
			Aliases:     []string{"n"},
			Usage:       "set build output name",
			Destination: &build_output_name,
		},
		&cli.StringFlag{
			Name:        "dir",
			Aliases:     []string{"d"},
			Usage:       "set build output directory",
			Destination: &build_output_directory,
		},
		&cli.BoolFlag{
			Name:        "no_debug",
			Usage:       "set build use -trimpath -ldflags '-s -w'",
			Destination: &build_no_debug,
		},
		&cli.BoolFlag{
			Name:        "no_cgo",
			Usage:       "set build not use cgo",
			Destination: &build_no_cgo,
		},
		&cli.StringSliceFlag{
			Name:        "opts",
			Usage:       "set build opts",
			Destination: &build_opts,
		},
		&cli.StringSliceFlag{
			Name:        "envs",
			Usage:       "set build envs",
			Destination: &build_envs,
		},
	}

	// 打印版本函数
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("%s", cCtx.App.Version)
	}

	app := &cli.App{
		Usage:   "Golang Build Tool",
		Version: "v2.00",
		Flags:   flags,
		Action: func(ctx *cli.Context) error {
			var ps []build.Pair

			if build_main {
				ps = build.GetMainPairs()
			} else if build_all {
				ps = build.GetAllPairs()
			} else {
				ps = build.GetSelectedPairs(build_os, build_arch)
			}

			for _, p := range ps {
				err := build.Build(p.OS, p.ARCH, build_output_name, build_output_directory, build_no_debug, build_no_cgo, build_opts.Value(), build_envs.Value())
				if err != nil {
					log.Println(err)
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
