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
	"gaia-updater/cmd"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gaiad-updater",
	Short: "update cosmos-sdk-gaia repository",
	Long: `Update and compile cosmos-sdk-gaia repository:

You have to specify the git version/tag to checkout and compile the right version`,
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
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
