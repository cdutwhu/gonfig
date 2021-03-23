package gonfig

// Config : AUTO Created From "gonfig/toml/test.toml"
type Config struct {
	Port int
	Service interface{}
	IP interface{}
	LogFile string
	Path interface{}
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
		Update string
		Delete string
		LsID string
		LsObject string
		LsContext string
		LsUser string
		ROOT string
		Get string
		GetHash string
		GetID string
	}
	File struct {
		ClientWin64 string
		MaskConfig string
		MaskLinux64 string
		MaskMac string
		MaskWin64 string
		ClientConfig string
		ClientLinux64 string
		ClientMac string
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

// Config1 : AUTO Created From "gonfig/toml/test1.toml"
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
		Get string
		GetID string
		LsContext string
		LsID string
		LsUser string
		ROOT string
		Update string
		Delete string
		GetHash string
		LsObject string
	}
	File struct {
		MaskConfig string
		MaskLinux64 string
		MaskMac string
		MaskWin64 string
		ClientConfig string
		ClientLinux64 string
		ClientMac string
		ClientWin64 string
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
