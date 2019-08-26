/*
Copyright © 2019 M-Way Solutions GmbH,
Author: Marcel Pohland <m.pohland@mwaysolutions.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-gaia-updater/cmd"
	"os"
)

var cfgFile string
var ProjectBase string
var userLicense string

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "go-gaiad-updater",
	Short: "update cosmos-sdk-gaia repository",
	Long: `Update and compile cosmos-sdk-gaia repository:

You have to specify the git version/tag to checkout and compile the right version`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	// rootCmd.AddCommand(rootCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rootCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func main() {
	fmt.Println(`
                   _         _
                  | |       | |
 _   _  _ __    __| |  __ _ | |_   ___  _ __
| | | || '_ \  / _| |` + ` / _` + `| |__|_| / _ \| '__|
| |_| || |_) || (_| || (_| || |_ |  __/| |
 \__,_|| .__/  \__,_| \__,_| \__| \___||_|
       | |
       |_|`)

	rootCmd.AddCommand(
		cmd.VersionCmd,
		cmd.StartCmd)
	rootCmd.Execute()
}

