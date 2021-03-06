package gonfig

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
)

// SetDftCfgVal :
func SetDftCfgVal(prjName, prjVer string) {
	dftPrjName = prjName
	dftPrjVer = prjVer
}

// PrjName :
func PrjName() (string, bool) {
	const check = "/.git"
NEXT:
	for i := 1; i < 64; i++ {
		for _, ln := range splitLn(trackCaller(i)) {
			if sHasPrefix(ln, "/") {
				ln = rmTailFromLast(ln, ":")
			AGAIN:
				dir := filepath.Dir(ln)
				if dir == "/" {
					continue NEXT
				}
				_, err := os.Stat(dir + check)
				if os.IsNotExist(err) {
					ln = dir
					goto AGAIN
				} else {
					return filepath.Base(dir), true
				}
			}
		}
	}
	return dftPrjName, false
}

// GitVer :
func GitVer() (string, bool) {
	tag, err := GitTag()
	if err != nil {
		return dftPrjVer, false
	}
	if r := rxMustCompile(`^v\d+\.\d+\.\d+$`); r.MatchString(tag) {
		return tag, true
	}
	return dftPrjVer, false
}

// GitTag :
func GitTag() (tag string, err error) {
	defer func() {
		if r := recover(); r != nil {
			tag, err = "", fEf("%v", r)
		}
	}()

	// check git existing
	_, oriWD := prepare("git") // maybe invoke panic
	os.Chdir(oriWD)            // under .git project dir to get `git tag`

	// run git
	cmd := exec.Command("bash", "-c", "git describe --tags")
	output, err := cmd.Output()
	// failOnErr("cmd.Output() error @ %v", err) // DO NOT PANIC
	if outstr := sTrim(string(output), " \n\t"); outstr != "" {
		return sSplit(outstr, "-")[0], nil
	}
	return "", nil
}

// Modify : only 2 levels struct variable could be modified. that is enough for config
func Modify(cfg interface{}, mRepl map[string]interface{}) interface{} {
	if len(mRepl) == 0 {
		return cfg
	}
	if vof(cfg).Kind() == typPTR {
		if cfgElem := vof(cfg).Elem(); cfgElem.Kind() == typSTRUCT {
			for i, nField := 0, cfgElem.NumField(); i < nField; i++ {
				for key, value := range mRepl {
					var ivalue interface{} = value
					repVal, isstrValue := value.(string)
					field := cfgElem.Field(i)
					if oriVal, ok := field.Interface().(string); ok && sContains(oriVal, key) {
						if isstrValue {
							ivalue = sReplaceAll(oriVal, key, repVal)
						}
						field.Set(vof(ivalue))
					}
					// go into struct element
					if field.Kind() == typSTRUCT {
						for j, nFieldSub := 0, field.NumField(); j < nFieldSub; j++ {
							fieldSub := field.Field(j)
							if oriVal, ok := fieldSub.Interface().(string); ok && sContains(oriVal, key) {
								if isstrValue {
									ivalue = sReplaceAll(oriVal, key, repVal)
								}
								fieldSub.Set(vof(ivalue))
							}
						}
					}
				}
			}
			return cfg
		}
	}
	failP1OnErr("%v", fEf("input cfg MUST be struct pointer"))
	return nil
}

// EvalCfgValue :
func EvalCfgValue(cfg interface{}, key string) interface{} {
	bytes, err := json.MarshalIndent(cfg, "", "\t")
	failP1OnErr("%v", err)
	lines := splitLn(string(bytes))
	if sCount(key, ".") == 0 {
		for _, ln := range lines {
			if sHasPrefix(ln, fSf("\t\"%s\":", key)) {
				sval := sTrim(sSplit(ln, ": ")[1], ",\"")
				switch {
				case isNumeric(sval) && !sContains(sval, "."):
					ret, _ := strconv.ParseInt(sval, 10, 64)
					return int(ret)
				case isNumeric(sval) && sContains(sval, "."):
					ret, _ := strconv.ParseFloat(sval, 64)
					return ret
				case sval == "true" || sval == "false":
					ret, _ := strconv.ParseBool(sval)
					return ret
				default:
					return sval
				}
			}
		}
	} else if sCount(key, ".") == 1 {
		ss := sSplit(key, ".")
		part1, part2 := ss[0], ss[1]
	NEXT:
		for i, ln1 := range lines {
			if sHasPrefix(ln1, fSf("\t\t\"%s\":", part2)) {
				for j := i - 1; j >= 0; j-- {
					ln2 := lines[j]
					if sHasPrefix(ln2, "\t\"") {
						if sHasPrefix(ln2, fSf("\t\"%s\":", part1)) {
							sval := sTrim(sSplit(ln1, ": ")[1], ",\"")
							switch {
							case isNumeric(sval) && !sContains(sval, "."):
								ret, _ := strconv.ParseInt(sval, 10, 64)
								return int(ret)
							case isNumeric(sval) && sContains(sval, "."):
								ret, _ := strconv.ParseFloat(sval, 64)
								return ret
							case sval == "true" || sval == "false":
								ret, _ := strconv.ParseBool(sval)
								return ret
							default:
								return sval
							}
						}
						continue NEXT
					}
				}
			}
		}
	}
	return nil
}

// ------------------------------------------------------------------------------- //

// InitEnvVar : initialize the global variables
func InitEnvVar(cfg interface{}, mReplExpr map[string]string, key string, cfgPaths ...string) interface{} {
	cfg = New(cfg, mReplExpr, append(cfgPaths, "./config.toml")...)
	if cfg == nil {
		return nil
	}
	struct2Env(key, cfg)
	return cfg
}

// New :
func New(cfg interface{}, mReplExpr map[string]string, cfgPaths ...string) interface{} {
	defer func() { mux.Unlock() }()
	mux.Lock()
	for _, f := range cfgPaths {
		if _, e := os.Stat(f); e == nil {
			return initCfg(f, cfg, mReplExpr)
		}
	}
	return nil
}

func initCfg(fpath string, cfg interface{}, mReplExpr map[string]string) interface{} {
	_, e := toml.DecodeFile(fpath, cfg)
	failPnOnErr(2, "%v", e)
	abs, e := filepath.Abs(fpath)
	failOnErr("%v", e)
	home, e := os.UserHomeDir()
	failOnErr("%v", e)
	prj, _ := PrjName()
	ver, _ := GitVer()

	cfg = Modify(cfg, map[string]interface{}{
		"~":      home,
		"[DATE]": time.Now().Format("2006-01-02"),
		"[PATH]": abs,
		"[IP]":   localIP(),
		"[PRJ]":  prj,
		"[VER]":  ver,
	})

	mRepl := make(map[string]interface{})
	for k, v := range mReplExpr {
		value := EvalCfgValue(cfg, v)
		if value != nil {
			mRepl[k] = value
		} else {
			mRepl[k] = v
		}
	}
	return Modify(cfg, mRepl)
}

// Save :
func Save(fpath string, cfg interface{}) {
	if !sHasSuffix(fpath, ".toml") {
		fpath += ".toml"
	}
	f, e := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	failP1OnErr("%v", e)
	defer f.Close()
	failP1OnErr("%v", toml.NewEncoder(f).Encode(cfg))
}

// --------------------------------------------------------------------------------------------------- //

// Register : echo 'password' | sudo -S env "PATH=$PATH" go test -v -count=1 ./ -run TestRegister
// func Register(funcOSUser, tomlFile, prjName, pkgName string) (bool, string) {
// 	enableLog2F(true, logfile)
// 	if funcOSUser == "" {
// 		user, err := user.Current()
// 		failOnErr("%v", err)
// 		funcOSUser = user.Name
// 	}

// 	pkgName = sToLower(pkgName)
// 	dir, _ := callerSrc()
// 	gonfigDir := dir                                                      // filepath.Dir(dir)
// 	gonfigDir = sReplace(gonfigDir, "/root/", "/home/"+funcOSUser+"/", 1) // sudo root go pkg --> input OS-user go pkg
// 	file := gonfigDir + fSf("/.cache/%s/%s/Config.go", prjName, pkgName)  // cfg struct Name as to be go fileName

// 	logger("ready to generate: %v", file)
// 	if !strugen.GenStruct(tomlFile, "Config", pkgName, file) {
// 		return false, ""
// 	}
// 	logger("finish generating: %v", file)

// 	// file LIKE `/home/qmiao/go/pkg/mod/github.com/cdutwhu/gonfig@v0.1.2/***/***/Config.go`
// 	pkgmark := "/go/pkg/mod/"
// 	if sContains(file, pkgmark) {
// 		fullpkg := filepath.Dir(sSplit(file, pkgmark)[1])
// 		logger("generated package path: %v", fullpkg)
// 		pos := rxMustCompile(`@[^/]+/`).FindAllStringIndex(fullpkg, -1)
// 		pkg := replByPosGrp(fullpkg, pos, []string{""}, 0, 1)
// 		logger("generated package: %v", pkg)
// 		// make necessary functions for using
// 		mkFuncs(pkg, prjName, pkgName, gonfigDir)
// 		return true, pkg
// 	}
// 	return false, file
// }

// func mkFuncs(impt, prj, pkg, fnDir string) {
// 	pkg = sToLower(pkg)
// 	fnFile := fnDir + "/auto_" + prj + "_" + pkg + ".go"

// 	prj = replAllOnAny(prj, []string{"-", " "}, "")
// 	pkg = replAllOnAny(pkg, []string{"-", " "}, "")

// 	prj, pkg = sTitle(prj), sTitle(pkg)
// 	fnNew := `New` + prj + pkg
// 	fnToEnv := `ToEnv` + prj + pkg
// 	fnFromEnv := `FromEnv` + prj + pkg

// 	src := `package gonfig` + "\n\n"
// 	src += `import auto "` + impt + `"` + "\n"
// 	src += `import "os"` + "\n\n"
// 	src += `func ` + fnNew + `(mReplExpr map[string]string, cfgPaths ...string) *auto.Config {` + "\n"
// 	src += `    defer func() { mux.Unlock() }()` + "\n"
// 	src += `    mux.Lock()` + "\n"
// 	src += `    cfg := &auto.Config{}` + "\n"
// 	src += `    for _, f := range cfgPaths {` + "\n"
// 	src += `        if _, e := os.Stat(f); e == nil {` + "\n"
// 	src += `            return initCfg(f, cfg, mReplExpr).(*auto.Config)` + "\n"
// 	src += `        }` + "\n"
// 	src += `    }` + "\n"
// 	src += `    return nil` + "\n"
// 	src += `}` + "\n\n"
// 	src += `// -------------------------------- //` + "\n\n"
// 	src += `func ` + fnToEnv + `(mReplExpr map[string]string, key string, cfgPaths ...string) *auto.Config {` + "\n"
// 	src += `    cfg := ` + fnNew + `(mReplExpr, append(cfgPaths, "./config.toml")...)` + "\n"
// 	src += `    if cfg == nil {` + "\n"
// 	src += `        return nil` + "\n"
// 	src += `    }` + "\n"
// 	src += `    struct2Env(key, cfg)` + "\n"
// 	src += `    return cfg` + "\n"
// 	src += `}` + "\n\n"
// 	src += `// -------------------------------- //` + "\n\n"
// 	src += `func ` + fnFromEnv + `(key string) *auto.Config {` + "\n"
// 	src += `    return env2Struct(key, &auto.Config{}).(*auto.Config)` + "\n"
// 	src += `}` + "\n\n"

// 	mustWriteFile(fnFile, []byte(src))
// }
