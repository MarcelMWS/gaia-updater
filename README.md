# gaia-upgrade

### Install [go](https://golang.org/doc/install)

```
go version go1.12.7 darwin/amd64
``` 
---
(first time)

create go path:

```
mkdir -p $HOME/go/bin
```
clone cosmos-sdk gaia repository (if you haven't got it yet):
```
mkdir -p $GOPATH/src/github.com/cosmos
cd $GOPATH/src/github.com/cosmos
git clone https://github.com/cosmos/gaia
cd cosmos-sdk
```

set environment variables (one time):

```
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

set environment variables (permanent):

```
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
source ~/.bash_profile
echo "export GOBIN=$GOPATH/bin" >> ~/.bash_profile
source ~/.bash_profile
echo "export PATH=$PATH:$GOBIN" >> ~/.bash_profile
source ~/.bash_profile
```

---

check:
```
echo $GOPATH
echo $GOBIN
echo $PATH
```

in gaia-updater dir run:

```
go install
```

start:

```
gaia-updater
```


man pages:

```

                   _         _
                  | |       | |
 _   _  _ __    __| |  __ _ | |_   ___  _ __
| | | || '_ \  / _| | / _| |__|_| / _ \| '__|
| |_| || |_) || (_| || (_| || |_ |  __/| |
 \__,_|| .__/  \__,_| \__,_| \__| \___||_|
       | |
       |_|
Update and compile cosmos-sdk-gaia repository:

You have to specify the git version/tag to checkout and compile the right version

Usage:
  go-gaiad-updater [command]

Available Commands:
  help        Help about any command
  start       start update
  version     print version

Flags:
  -h, --help   help for go-gaiad-updater

Use "go-gaiad-updater [command] --help" for more information about a command.
```

start command man pages:

```
                   _         _
                  | |       | |
 _   _  _ __    __| |  __ _ | |_   ___  _ __
| | | || '_ \  / _| | / _| |__|_| / _ \| '__|
| |_| || |_) || (_| || (_| || |_ |  __/| |
 \__,_|| .__/  \__,_| \__,_| \__| \___||_|
       | |
       |_|
Error: required flag(s) "version" not set
Usage:
  go-gaiad-updater start [flags]

Flags:
  -c, --configPath string     gaia config location (default "/Users/***/.gaiad/config")
  -g, --gaiaRepoPath string   gaia repo location (default "/Users/***/go/src/github.com/cosmos/gaia")
  -h, --help                  help for start
  -l, --link string           link to genesis (default "https://raw.githubusercontent.com/cosmos/testnets/master/gaia-13k/genesis.json")
  -v, --version string        provide correct git tag e.x. v2.0.0

required flag(s) "version" not set
```

### upgrade with new genesisfile, steps executed automatically 
```
cd
cd go/src/github.com/cosmos/gaia/
git fetch --tags
git checkout .
git checkout v1.0.0-rc3                       #variable?
git clean -fd
git clean -fx
go version                                    #check minimal version
make go-mod-cache                             #or set env to?: GO111MODULE=on
make install
gaiad version                                 #variable to check version?
cd
gaiad unsafe-reset-all
cd .gaiad/config/
rm genesis.json
wget https://raw.githubusercontent.com/cosmos/testnets/master/gaia-13k/genesis.json #variable?
sha256sum genesis.json                        #variable?
```
