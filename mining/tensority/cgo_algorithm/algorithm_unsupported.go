//go:build !simd
// +build !simd

package cgo_algorithm

import (
	log "github.com/sirupsen/logrus"

	"github.com/anonymitycash/anonymitycash/mining/tensority/go_algorithm"
	"github.com/anonymitycash/anonymitycash/protocol/bc"
)

func SimdAlgorithm(bh, seed *bc.Hash) *bc.Hash {
	log.Warn("SIMD feature is not supported on release version, please compile the lib according to README to enable this feature.")
	return go_algorithm.LegacyAlgorithm(bh, seed)
}
