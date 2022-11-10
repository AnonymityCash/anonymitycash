Anonymitycash classic
======

[![Build Status](https://travis-ci.org/Anonymitycash/anonymitycash.svg)](https://travis-ci.org/Anonymitycash/anonymitycash) [![AGPL v3](https://img.shields.io/badge/license-AGPL%20v3-brightgreen.svg)](./LICENSE)

**Official golang implementation of the Anonymitycash classic protocol.**

Automated builds are available for stable releases and the unstable master branch. Binary archives are published at https://github.com/Anonymitycash/anonymitycash/releases.

## What is Anonymitycash classic?

Anonymitycash classic is software designed to operate and connect to highly scalable blockchain networks confirming to the Anonymitycash classic Blockchain Protocol, which allows partipicants to define, issue and transfer digitial assets on a multi-asset shared ledger. Please refer to the [White Paper](https://github.com/Anonymitycash/wiki/blob/master/en-US/Anonymitycash-Technical-White-Paper-EN.pdf) for more details.

In the current state `anonymitycash classic` is able to:

- Manage key, account as well as asset
- Send transactions, i.e., issue, spend and retire asset

## Installing with Homebrew

```
brew tap anonymitycash/anonymitycash && brew install anonymitycash
```

## Building from source

### Requirements

- [Go](https://golang.org/doc/install) version 1.8 or higher, with `$GOPATH` set to your preferred directory

### Installation

Ensure Go with the supported version is installed properly:

```bash
$ go version
$ go env GOROOT GOPATH
```

- Get the source code

``` bash
$ git clone https://github.com/Anonymitycash/anonymitycash.git $GOPATH/src/github.com/anonymitycash/anonymitycash
```

- Build source code

``` bash
$ cd $GOPATH/src/github.com/anonymitycash/anonymitycash
$ GO111MODULE=off make anonymitycashcd    # build anonymitycash daemon
$ GO111MODULE=off make anonymitycashccli  # build anonymitycash command line
```

When successfully building the project, the `anonymitycashd` and `anonymitycashcli` binary should be present in `cmd/anonymitycashcd` and `cmd/anonymitycashccli` directory, respectively.

### Executables

The Anonymitycash project comes with several executables found in the `cmd` directory.

| Command      | Description                                                  |
| ------------ | ------------------------------------------------------------ |
| **anonymitycashcd**   | anonymitycashd command can help to initialize and launch anonymitycash domain by custom parameters. `anonymitycashcd --help` for command line options. |
| **anonymitycashccli** | Our main Anonymitycash CLI client. It is the entry point into the Anonymitycash network (main-, test- or private net), capable of running as a full node archive node (retaining all historical state). It can be used by other processes as a gateway into the Anonymitycash network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `anonymitycashccli --help` and the [anonymitycashccli Wiki page](https://github.com/Anonymitycash/anonymitycash/wiki/Command-Line-Options) for command line options. |

## Running anonymitycash classic

Currently, anonymitycash classic is still in active development and a ton of work needs to be done, but we also provide the following content for these eager to do something with `anonymitycash classic`. This section won't cover all the commands of `anonymitycashcd` and `anonymitycashccli` at length, for more information, please the help of every command, e.g., `anonymitycashccli help`.

### Initialize

First of all, initialize the node:

```bash
$ cd ./cmd/anonymitycashcd
$ ./anonymitycashcd init --chain_id mainnet
```

There are three options for the flag `--chain_id`:

- `mainnet`: connect to the mainnet.
- `testnet`: connect to the testnet wisdom.
- `solonet`: standalone mode.

After that, you'll see `config.toml` generated, then launch the node.

### launch

``` bash
$ ./anonymitycashcd node
```

available flags for `anonymitycashcd node`:

```
Flags:
      --auth.disable                     Disable rpc access authenticate
      --chain_id string                  Select network type
  -h, --help                             help for node
      --log_file string                  Log output file (default "log")
      --log_level string                 Select log level(debug, info, warn, error or fatal)
      --mining                           Enable mining
      --p2p.dial_timeout int             Set dial timeout (default 3)
      --p2p.handshake_timeout int        Set handshake timeout (default 30)
      --p2p.keep_dial string             Peers addresses try keeping connecting to, separated by ',' (for example "1.1.1.1:46654;2.2.2.2:46653")
      --p2p.laddr string                 Node listen address. (0.0.0.0:0 means any interface, any port) (default "tcp://0.0.0.0:46655")
      --p2p.lan_discoverable             Whether the node can be discovered by nodes in the LAN (default true)
      --p2p.max_num_peers int            Set max num peers (default 50)
      --p2p.node_key string              Node key for p2p communication
      --p2p.proxy_address string         Connect via SOCKS5 proxy (eg. 127.0.0.1:1086)
      --p2p.proxy_password string        Password for proxy server
      --p2p.proxy_username string        Username for proxy server
      --p2p.seeds string                 Comma delimited host:port seed nodes
      --p2p.skip_upnp                    Skip UPNP configuration
      --prof_laddr string                Use http to profile anonymitycashd programs
      --simd.enable                      Enable SIMD mechan for tensority
      --vault_mode                       Run in the offline enviroment
      --wallet.disable                   Disable wallet
      --wallet.rescan                    Rescan wallet
      --wallet.txindex                   Save global tx index
      --web.closed                       Lanch web browser or not
      --ws.max_num_concurrent_reqs int   Max number of concurrent websocket requests that may be processed concurrently (default 20)
      --ws.max_num_websockets int        Max number of websocket connections (default 25)

Global Flags:
      --home string   root directory for config and data
  -r, --root string   DEPRECATED. Use --home (default "/Users/zcc/Library/Application Support/Anonymitycashclassic")
      --trace         print out full stack trace on errors
```

Given the `anonymitycashcd` node is running, the general workflow is as follows:

- create key, then you can create account and asset.
- send transaction, i.e., build, sign and submit transaction.
- query all kinds of information, let's say, avaliable key, account, key, balances, transactions, etc.

__simd feature:__

You could enable the _simd_ feature to speed up the _PoW_ verification (e.g., during mining and block verification) by simply:
```
anonymitycashcd node --simd.enable
```

To enable this feature you will need to compile from the source code by yourself, and `make anonymitycashcd-simd`. 

What is more,

+ if you are using _Mac_, please make sure _llvm_ is installed by `brew install llvm`.
+ if you are using _Windows_, please make sure _mingw-w64_ is installed and set up the _PATH_ environment variable accordingly.

For more details about using `anonymitycashcli` command please refer to [API Reference](https://github.com/Anonymitycash/anonymitycash/wiki/API-Reference)

### Dashboard

Access the dashboard:

```
$ open http://localhost:1313/
```

### In Docker

Ensure your [Docker](https://www.docker.com/) version is 17.05 or higher.

```bash
$ docker build -t anonymitycash .
```

For the usage please refer to [running-in-docker-wiki](https://github.com/Anonymitycash/anonymitycash/wiki/Running-in-Docker).

## Contributing

Thank you for considering helping out with the source code! Any contributions are highly appreciated, and we are grateful for even the smallest of fixes!

If you run into an issue, feel free to [anonymitycash classic issues](https://github.com/Anonymitycash/anonymitycash/issues/) in this repository. We are glad to help!

## License

[AGPL v3](./LICENSE)
