package strugen

import (
	"os"
	"testing"
)

var (
	toml  = "../toml/test.toml"
	toml1 = "../toml/test1.toml"
)

func TestScanToml(t *testing.T) {
	bytes, err := os.ReadFile(toml)
	failP1OnErr("%v", err)
	lines := splitLn(string(bytes))
	attrs1, attrs2 := scanToml(lines)
	fPln(attrs1)
	fPln(attrs2)
}

func TestAttrsRange(t *testing.T) {
	bytes, err := os.ReadFile(toml)
	failP1OnErr("%v", err)
	lines := splitLn(string(bytes))
	for k, v := range attrsRange(lines) {
		fPln(k, v)
	}
}

func TestAttrTypes(t *testing.T) {
	bytes, err := os.ReadFile(toml)
	failP1OnErr("%v", err)
	lines := splitLn(string(bytes))
	for k, v := range attrTypes(lines, "") {
		fPln(k, v)
	}
}

func TestGenStruct(t *testing.T) {
	enableWarnDetail(false)
	cfgSrc := "../Config.go"
	os.Remove(cfgSrc)
	GenStruct(toml, "Config", "gonfig", cfgSrc)
	GenStruct(toml1, "Config1", "gonfig", cfgSrc)
	GenNewCfg(cfgSrc)
}
