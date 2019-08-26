# go-gaia-upgrade

### Install go

```
go version go1.12.7 darwin/amd64
``` 

in dir run:

```
go install
```

man pages:

```
                   _         _
                  | |       | |
 _   _  _ __    __| |  __ _ | |_   ___  _ __
| | | || '_ \  / _` | / _` || __| / _ \| '__|
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

```

start command man pages:

```
                   _         _
                  | |       | |
 _   _  _ __    __| |  __ _ | |_   ___  _ __
| | | || '_ \  / _` | / _` || __| / _ \| '__|
| |_| || |_) || (_| || (_| || |_ |  __/| |
 \__,_|| .__/  \__,_| \__,_| \__| \___||_|
       | |
       |_|
Error: required flag(s) "version" not set
Usage:
  go-gaiad-updater start [flags]

Flags:
  -c, --configPath string      gaia config location HOME + (default ".gaiad/config/")
  -g, --gaiaBuildPath string   gaia repo location HOME + (default "go/src/github.com/cosmos/gaia/")
  -h, --help                   help for start
  -l, --link string            link to genesis (default "https://raw.githubusercontent.com/cosmos/testnets/master/gaia-13k/genesis.json")
  -v, --version string         provide correct git tag e.x. v2.0.0

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
sudo systemctl stop gaiad
ps -ax | grep gaia                            #check gaiad killed?
go version                                    #check minimal version
echo $GOPATH                                  #check envs
echo $GOBIN
make go-mod-cache                             #or set env to?: GO111MODULE=on
make install
gaiad version                                 #variable to check version?
cd
df -h                                         #check disk space
cp -r .gaiad/data/ .gaiad/data_backup         #backup blockchain data
gaiad unsafe-reset-all
df -h                                       #check disk space if old database is clean/removed
cd .gaiad/config/
vim config.toml                               #check configs peers, pri_ids
ls                                            #check dir
rm genesis.json
wget https://raw.githubusercontent.com/cosmos/testnets/master/gaia-13k/genesis.json #variable?
sha256sum genesis.json                        #variable?
sudo systemctl start gaiad                    #ensure set ulimit in systemd service file
```
