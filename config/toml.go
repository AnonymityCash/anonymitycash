package config

import (
	"path"

	cmn "github.com/tendermint/tmlibs/common"
)

/****** these are for production settings ***********/
func EnsureRoot(rootDir string, network string) {
	cmn.EnsureDir(rootDir, 0700)
	cmn.EnsureDir(rootDir+"/data", 0700)

	configFilePath := path.Join(rootDir, "config.toml")

	// Write default config file if missing.
	if !cmn.FileExists(configFilePath) {
		cmn.MustWriteFile(configFilePath, []byte(selectNetwork(network)), 0644)
	}
}

var defaultConfigTmpl = `# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml
fast_sync = true
db_backend = "leveldb"
api_addr = "0.0.0.0:1313"
node_alias = ""
`

var mainNetConfigTmpl = `chain_id = "mainnet"
[p2p]
laddr = "tcp://0.0.0.0:46654"
seeds = "158.101.199.138:46654"
`

var testNetConfigTmpl = `chain_id = "wisdom"
[p2p]
laddr = "tcp://0.0.0.0:46655"
seeds = "158.101.199.138:46655"
`

var soloNetConfigTmpl = `chain_id = "solonet"
[p2p]
laddr = "tcp://0.0.0.0:46653"
seeds = ""
`

// Select network seeds to merge a new string.
func selectNetwork(network string) string {
	switch network {
	case "mainnet":
		return defaultConfigTmpl + mainNetConfigTmpl
	case "testnet":
		return defaultConfigTmpl + testNetConfigTmpl
	default:
		return defaultConfigTmpl + soloNetConfigTmpl
	}
}
