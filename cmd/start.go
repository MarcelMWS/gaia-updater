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

// startCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		gaiaPath := "/go/src/github.com/cosmos/gaia/"
		configPath := home + "/.gaiad/config/"
		fmt.Println("YOUR GAIA-SRC-HOME address: " + home + gaiaPath)
		GitFetchCommand(home + gaiaPath)
		GitCheckoutCommand(home + gaiaPath)
		GitCheckoutVersionCommand(home + gaiaPath)
		GitCheckoutCleanFDCommand(home + gaiaPath)
		GitCheckoutCleanFXCommand(home + gaiaPath)
		// StopGaia(home)
		GoVersionCheck(home + gaiaPath)
		CheckGOPATH()
		MakeGoModCache(home + gaiaPath)
		MakeInstall(home + gaiaPath)
		CheckVersion(home + gaiaPath)
		GaiaUnsafeResetAll(home)
		RemoveGenesis(configPath)
		GetGenesis(configPath)
		ChecksumGenesis(configPath)
		// StartGaia(home)
	},
}

func init() {

	//rootCmd.AddCommand(startCmd)

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

func GitCheckoutVersionCommand(dir string) {
	cmd := exec.Command("git", "checkout", "v2.0.0")
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

func GetGenesis(dir string) {
	cmd := exec.Command("wget", "https://raw.githubusercontent.com/cosmos/testnets/master/gaia-13k/genesis.json")
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
