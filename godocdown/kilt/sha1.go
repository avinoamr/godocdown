// This file was AUTOMATICALLY GENERATED by kilt-import (smuggol) from github.com/robertkrimen/kilt

package kilt

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

// Sha1Path
func (self Kilt) Sha1Path(path ...string) string {
	return Sha1Path(path...)
}

func Sha1Path(path ...string) string {
	file, err := os.Open(filepath.Join(path...))
	if err != nil {
		return ""
	}
	return Sha1Of(file)
}

// Sha1Of
func (self Kilt) Sha1Of(src io.Reader) string {
	return Sha1Of(src)
}

func Sha1Of(src io.Reader) string {
	hash := sha1.New()
	_, err := io.Copy(hash, src)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha1
func (self Kilt) Sha1(data []byte) string {
	return Sha1(data)
}

func Sha1(data []byte) string {
	hash := sha1.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
