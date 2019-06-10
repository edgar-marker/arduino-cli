/*
 * This file is part of arduino-cli.
 *
 * Copyright 2018 ARDUINO SA (http://www.arduino.cc/)
 *
 * This software is released under the GNU General Public License version 3,
 * which covers the main part of arduino-cli.
 * The terms of this license can be found at:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 * You can be released from the requirements of the above licenses by purchasing
 * a commercial license. Buying such a license is mandatory if you want to modify or
 * otherwise use the software for commercial activities involving the Arduino
 * software without disclosing the source code of your own applications. To purchase
 * a commercial license, send an email to license@arduino.cc.
 */

package compile

import (
	"context"
	"errors"
	"os"

	"github.com/arduino/arduino-cli/cli"
	"github.com/arduino/arduino-cli/commands/compile"
	"github.com/arduino/arduino-cli/common/formatter"
	"github.com/arduino/arduino-cli/proto-gen/commands/rpc_v1"
	"github.com/arduino/go-paths-helper"
	"github.com/spf13/cobra"
)

// InitCommand prepares the command.
func InitCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     "compile",
		Short:   "Compiles Arduino sketches.",
		Long:    "Compiles Arduino sketches.",
		Example: "  " + cli.AppName + " compile -b arduino:avr:uno /home/user/Arduino/MySketch",
		Args:    cobra.MaximumNArgs(1),
		Run:     run,
	}
	command.Flags().StringVarP(
		&flags.fqbn, "fqbn", "b", "",
		"Fully Qualified Board Name, e.g.: arduino:avr:uno")
	command.Flags().BoolVar(
		&flags.showProperties, "show-properties", false,
		"Show all build properties used instead of compiling.")
	command.Flags().BoolVar(
		&flags.preprocess, "preprocess", false,
		"Print preprocessed code to stdout instead of compiling.")
	command.Flags().StringVar(
		&flags.buildCachePath, "build-cache-path", "",
		"Builds of 'core.a' are saved into this path to be cached and reused.")
	command.Flags().StringVarP(
		&flags.exportFile, "output", "o", "",
		"Filename of the compile output.")
	command.Flags().StringVar(
		&flags.buildPath, "build-path", "",
		"Path where to save compiled files. If omitted, a directory will be created in the default temporary path of your OS.")
	command.Flags().StringSliceVar(
		&flags.buildProperties, "build-properties", []string{},
		"List of custom build properties separated by commas. Or can be used multiple times for multiple properties.")
	command.Flags().StringVar(
		&flags.warnings, "warnings", "none",
		`Optional, can be "none", "default", "more" and "all". Defaults to "none". Used to tell gcc which warning level to use (-W flag).`)
	command.Flags().BoolVarP(
		&flags.verbose, "verbose", "v", false,
		"Optional, turns on verbose mode.")
	command.Flags().BoolVar(
		&flags.quiet, "quiet", false,
		"Optional, supresses almost every output.")
	command.Flags().StringVar(
		&flags.vidPid, "vid-pid", "",
		"When specified, VID/PID specific build properties are used, if boards supports them.")
	return command
}

var flags struct {
	fqbn            string   // Fully Qualified Board Name, e.g.: arduino:avr:uno.
	showProperties  bool     // Show all build preferences used instead of compiling.
	preprocess      bool     // Print preprocessed code to stdout.
	buildCachePath  string   // Builds of 'core.a' are saved into this path to be cached and reused.
	buildPath       string   // Path where to save compiled files.
	buildProperties []string // List of custom build properties separated by commas. Or can be used multiple times for multiple properties.
	warnings        string   // Used to tell gcc which warning level to use.
	verbose         bool     // Turns on verbose mode.
	quiet           bool     // Suppresses almost every output.
	vidPid          string   // VID/PID specific build properties.
	exportFile      string   // The compiled binary is written to this file
}

// FIXME: temporary fix
func createInstance() *rpc_v1.Instance {
	resp := cli.InitInstance()
	if resp.GetPlatformsIndexErrors() != nil {
		for _, err := range resp.GetPlatformsIndexErrors() {
			formatter.PrintError(errors.New(err), "Error loading index")
		}
		formatter.PrintErrorMessage("Launch 'core update-index' to fix or download indexes.")
		os.Exit(cli.ErrGeneric)
	}
	return &rpc_v1.Instance{Id: resp.GetInstance().Id}
}

func run(cmd *cobra.Command, args []string) {
	instance := createInstance()

	var path *paths.Path
	if len(args) > 0 {
		path = paths.New(args[0])
	}
	sketchPath := cli.InitSketchPath(path)
	_, err := compile.Compile(context.Background(), &rpc_v1.CompileReq{
		Instance:        instance,
		Fqbn:            flags.fqbn,
		SketchPath:      sketchPath.String(),
		ShowProperties:  flags.showProperties,
		Preprocess:      flags.preprocess,
		BuildCachePath:  flags.buildCachePath,
		BuildPath:       flags.buildPath,
		BuildProperties: flags.buildProperties,
		Warnings:        flags.warnings,
		Verbose:         flags.verbose,
		Quiet:           flags.quiet,
		VidPid:          flags.vidPid,
		ExportFile:      flags.exportFile,
	}, os.Stdout, os.Stderr)
	if err == nil {
		// FIXME: re-enable
		//outputCompileResp(compRes)
	} else {
		formatter.PrintError(err, "Error during build")
		os.Exit(cli.ErrGeneric)
	}
}

func outputCompileResp(details *rpc_v1.ExecResp) {

}
