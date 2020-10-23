package gonfig

// Config : AUTO Created From /home/qmiao/Desktop/gonfig/toml/test.toml
type Config struct {
	LogFile string
	Path interface{}
	Port int
	Service interface{}
	IP interface{}
	Storage struct {
		BadgerDBPath string
		DataBase string
		MetaPath string
	}
	WebService struct {
		Port interface{}
		Service string
		Version interface{}
	}
	Route struct {
		Delete string
		GetHash string
		LsID string
		LsObject string
		Update string
		Get string
		GetID string
		LsContext string
		LsUser string
		ROOT string
	}
	File struct {
		ClientMac string
		ClientWin64 string
		MaskConfig string
		MaskLinux64 string
		MaskMac string
		MaskWin64 string
		ClientConfig string
		ClientLinux64 string
	}
	Server struct {
		IP interface{}
		Port interface{}
		Protocol string
		Service string
	}
	Access struct {
		Timeout int
	}
	SchoolPrograms struct {
		BOOLEAN []string
	}
	CollectionStatus struct {
		NUMERIC []string
	}
	Journal struct {
		NUMERIC []string
	}
	SectionInfo struct {
		NUMERIC []string
	}
}

// Config1 : AUTO Created From /home/qmiao/Desktop/gonfig/toml/test1.toml
type Config1 struct {
	Version string
	Storage struct {
		BadgerDBPath string
		DataBase string
		MetaPath string
	}
	WebService struct {
		Port interface{}
		Service string
		Version interface{}
	}
	Route struct {
		ROOT string
		GetID string
		LsUser string
		GetHash string
		LsContext string
		LsID string
		LsObject string
		Update string
		Delete string
		Get string
	}
	File struct {
		MaskWin64 string
		ClientConfig string
		ClientLinux64 string
		ClientMac string
		ClientWin64 string
		MaskConfig string
		MaskLinux64 string
		MaskMac string
	}
	Server struct {
		Protocol string
		Service string
		IP interface{}
		Port interface{}
	}
	Access struct {
		Timeout int
	}
	SchoolPrograms struct {
		BOOLEAN []string
	}
	CollectionStatus struct {
		NUMERIC []string
	}
	Journal struct {
		NUMERIC []string
	}
	SectionInfo struct {
		NUMERIC []string
	}
}

// NewCfg :
func NewCfg(cfgStruName string, mReplExpr map[string]string, cfgPaths ...string) interface{} {
	var cfg interface{}
	switch cfgStruName {
	case "Config":
		cfg = &Config{}
	case "Config1":
		cfg = &Config1{}
	default:
		return nil
	}
	return InitEnvVar(cfg, mReplExpr, cfgStruName, cfgPaths...)
}
