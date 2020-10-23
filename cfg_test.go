package gonfig

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
)

func TestPrjName(t *testing.T) {
	SetDftCfgVal("dftPrjName", "0.0.0")
	fPln(PrjName())
}

func TestGitVer(t *testing.T) {
	SetDftCfgVal("dftPrjName", "0.0.0")
	fPln(GitVer())
}

func TestGitTag(t *testing.T) {
	fPln(GitTag())
}

func TestModify(t *testing.T) {
	cfg := &Config{}
	_, err := toml.DecodeFile("./toml/test.toml", cfg)
	failOnErr("%v", err)
	Icfg := Modify(cfg, map[string]interface{}{
		"[PORT]": 1234,
		"[s]":    "gonfig",
		"[v]":    "v1.2.3",
	})
	cfg = Icfg.(*Config)
	spew.Dump(cfg)
}

func TestNewCfg(t *testing.T) {
	SetDftCfgVal("dftPrjName", "0.0.0")
	prj, _ := PrjName()
	cfg := &Config{}
	ok := New(
		cfg,
		map[string]string{
			"[s]":   "WebService.Service",
			"[v]":   "WebService.Version",
			"[p]":   "Port",
			"[prj]": prj,
		},
		"./toml/test.toml",
	)
	fPln(ok)
	Save("./saved.toml", cfg)
}

func TestEvalCfgValue(t *testing.T) {
	cfg := &Config{}
	cfg = New(
		cfg,
		map[string]string{
			"[p]": "Port",
			"[s]": "WebService.Service",
			"[v]": "WebService.Version",
		},
		"./toml/test.toml",
	).(*Config)

	spew.Dump(cfg)

	fPln(EvalCfgValue(cfg, "Port"))
	fPln(EvalCfgValue(cfg, "Storage.MetaPath"))
	fPln(EvalCfgValue(cfg, "WebService.Port"))
	fPln(EvalCfgValue(cfg, "Server.Service"))
	fPln(EvalCfgValue(cfg, "WebService.Service"))

	Save("./saved.toml", cfg)
}

func TestToEnvVar(t *testing.T) {
	InitEnvVar(
		&Config{},
		map[string]string{
			"[p]": "Port",
			"[s]": "WebService.Service",
			"[v]": "WebService.Version",
		},
		"KEY",
		"./toml/test.toml",
	)
	cfg := env2Struct("KEY", &Config{})
	spew.Dump(cfg)
}

// echo 'password' | sudo -S env "PATH=$PATH" go test -v -count=1 ./ -run TestRegister
func TestRegister(t *testing.T) {
	prj, _ := PrjName()
	pkg := "Server"
	ok, file := Register("qmiao", "./toml/test.toml", prj, pkg)
	fPln(ok, file)
}
