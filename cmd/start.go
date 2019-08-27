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
package cmd

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var gaiaRepoPath string
var configPath string
var link string
var home string
var version string
var shasum string

// startCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start update",
	Long:  `start update and specify version`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("YOUR GAIA-SRC-HOME address: " + gaiaRepoPath)
		log.Println("YOUR GAIA-CONFIG-HOME address: " + configPath)
		log.Println("YOUR GAIA-GENESIS-LINK address: " + link)
		log.Println("YOUR GAIA-VERSION-TO-INSTALL: " + version)
		GitFetchCommand(gaiaRepoPath)
		GitCheckoutCleanFDCommand(gaiaRepoPath)
		GitCheckoutCleanFXCommand(gaiaRepoPath)
		GitCheckoutCommand(gaiaRepoPath)
		GitCheckoutVersionCommand(gaiaRepoPath, version)
		GoVersionPrint(gaiaRepoPath)
		CheckGOPATH()
		MakeGoModCache(gaiaRepoPath)
		MakeInstall(gaiaRepoPath)
		PrintGaiadVersion(gaiaRepoPath)
		GaiaUnsafeResetAll(home)
		RemoveGenesis(configPath)
		GetGenesis(configPath, link)
		if shasum != "" {
			ChecksumGenesis(configPath)
		}
	},
}

func init() {
	home, err := os.UserHomeDir()
	StartCmd.Flags().StringVarP(&gaiaRepoPath, "gaiaRepoPath", "g", filepath.Join(home, "go/src/github.com/cosmos/gaia/"), "gaia repo location")
	StartCmd.Flags().StringVarP(&configPath, "configPath", "c", filepath.Join(home, ".gaiad/config/"), "gaia config location")
	StartCmd.Flags().StringVarP(&link, "link", "l", "https://raw.githubusercontent.com/cosmos/testnets/master/gaia-13k/genesis.json", "link to genesis")
	StartCmd.Flags().StringVarP(&shasum, "shasum", "s", "", "provide sha256sum of genesis.json file")
	StartCmd.Flags().StringVarP(&version, "version", "v", "", "provide correct git tag e.x. v2.0.0")
	StartCmd.MarkFlagRequired("version")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
	log.Printf("Reset current branch: %q\n", out.String())
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
	log.Printf("Clean dir: %q\n", out.String())
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
	log.Printf("Clean files: %q\n", out.String())
}

func GoVersionPrint(dir string) {
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
		log.Fatal("Updater couldn't find an installation of go. Please install go and retry")
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

func PrintGaiadVersion(dir string) {
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
	err := os.Remove(filepath.Join(dir, "genesis.json"))
	if err != nil {
		if os.IsNotExist(err) {
			println("Couldn't delete genesis.json, file does not exist")
			return
		} else {
			log.Fatal(err)
		}
	}
	log.Printf("Remove genesis.json")
}

func GetGenesis(dir, link string) {
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath.Join(dir, "genesis.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded new genesisfile")
}

func ChecksumGenesis(dir string) {
	file, err := os.Open(filepath.Join(dir, "genesis.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}
	n, err := hex.DecodeString(shasum)
	if err != nil {
		log.Fatal(err)
	}
	if bytes.Equal(h.Sum(nil), n) {
		log.Printf("Correct checksum genesis.json: %x", h.Sum(nil))
	} else {
		log.Fatalln("False checksum genesis.json: ", hex.EncodeToString(h.Sum(nil)))
	}
}
