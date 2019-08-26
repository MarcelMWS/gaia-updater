/*
Copyright © 2019 Marcel Pohland <m.pohland@mwaysolutions.com>

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
package cmd

import (
	"bytes"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var gaiaBuildPath string
var configPath string
var link string
var version string

// startCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start update",
	Long:  `start update and specify version`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		home = home + "/"
		configPath = home + configPath
		log.Println("YOUR GAIA-SRC-HOME address: " + home + gaiaBuildPath)
		log.Println("YOUR GAIA-CONFIG-HOME address: " + configPath)
		log.Println("YOUR GAIA-GENESIS-LINK address: " + link)
		log.Println("YOUR GAIA-VERSION-TO-INSTALL: " + version)
		GitFetchCommand(home + gaiaBuildPath)
		GitCheckoutCleanFDCommand(home + gaiaBuildPath)
		GitCheckoutCleanFXCommand(home + gaiaBuildPath)
		GitCheckoutCommand(home + gaiaBuildPath)
		GitCheckoutVersionCommand(home+gaiaBuildPath, version)
		// StopGaia(home)
		GoVersionCheck(home + gaiaBuildPath)
		CheckGOPATH()
		MakeGoModCache(home + gaiaBuildPath)
		MakeInstall(home + gaiaBuildPath)
		CheckVersion(home + gaiaBuildPath)
		GaiaUnsafeResetAll(home)
		RemoveGenesis(configPath)
		GetGenesis(configPath, link)
		ChecksumGenesis(configPath)
		// StartGaia(home)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {

	//rootCmd.AddCommand(startCmd)
	StartCmd.Flags().StringVarP(&gaiaBuildPath, "gaiaBuildPath", "g", "go/src/github.com/cosmos/gaia/", "gaia repo location HOME +")
	StartCmd.Flags().StringVarP(&configPath, "configPath", "c", ".gaiad/config/", "gaia config location HOME +")
	StartCmd.Flags().StringVarP(&link, "link", "l", "https://raw.githubusercontent.com/cosmos/testnets/master/gaia-13k/genesis.json", "link to genesis")
	StartCmd.Flags().StringVarP(&version, "version", "v", "", "provide correct git tag e.x. v2.0.0")
	StartCmd.MarkFlagRequired("version")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GitFetchCommand(dir string) {
	cmd := exec.Command("git", "fetch", "--all")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q\n", out.String())
}
func GitCheckoutCommand(dir string) {
	cmd := exec.Command("git", "checkout", ".")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Checkout unnecessary files: %q\n", out.String())
}

func GitCheckoutVersionCommand(dir, version string) {
	cmd := exec.Command("git", "checkout", version)
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Checkout your prefered version: %q\n", out.String())
}

func GitCheckoutCleanFDCommand(dir string) {
	cmd := exec.Command("git", "clean", "-fd")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Git clean dir: %q\n", out.String())
}

func GitCheckoutCleanFXCommand(dir string) {
	cmd := exec.Command("git", "clean", "-fx")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Git clean files: %q\n", out.String())
}

func StopGaia(dir string) {
	cmd := exec.Command("sudo", "systemctl", "stop", "gaiad")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Stop gaia: %q\n", out.String())
}

func GoVersionCheck(dir string) {
	cmd := exec.Command("go", "version")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GOVERSION: %q\n", out.String())
}

func CheckGOPATH() {
	path, err := exec.LookPath("go")
	if err != nil {
		log.Fatal("installing go is in your future/please set correct environment")
	}
	log.Printf("go is available at %s\n", path)
}

func MakeGoModCache(dir string) {
	cmd := exec.Command("make", "go-mod-cache")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Log: %q\n", out.String())
}

func MakeInstall(dir string) {
	cmd := exec.Command("make", "install")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Log: %q\n", out.String())
}

func CheckVersion(dir string) {
	cmd := exec.Command("gaiad", "version")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("New binary version: %q\n", out.String())
}

func GaiaUnsafeResetAll(dir string) {
	cmd := exec.Command("gaiad", "unsafe-reset-all")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Unsafe Reset ALL: %q\n", out.String())
}

func RemoveGenesis(dir string) {
	cmd := exec.Command("rm", "genesis.json")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Remove Genesis: %q\n", out.String())
}

func GetGenesis(dir, link string) {
	cmd := exec.Command("wget", link)
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Get New Genesis: %q\n", out.String())
}

func ChecksumGenesis(dir string) {
	cmd := exec.Command("shasum", "-a", "256", "genesis.json")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Checksum Genesis: %q\n", out.String())
}

func StartGaia(dir string) {
	cmd := exec.Command("sudo", "systemctl", "start", "gaiad")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start gaia: %q\n", out.String())
}


