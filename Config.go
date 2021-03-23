package gonfig

// Config : AUTO Created From "gonfig/toml/test.toml"
type Config struct {
	Service interface{}
	IP interface{}
	LogFile string
	Path interface{}
	Port int
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
		GetHash string
		LsContext string
		ROOT string
		Delete string
		GetID string
		LsID string
		LsObject string
		LsUser string
		Update string
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
		Port interface{}
		Protocol string
		Service string
		IP interface{}
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
		Service string
		Version interface{}
		Port interface{}
	}
	Route struct {
		LsObject string
		ROOT string
		Get string
		GetHash string
		LsID string
		LsUser string
		Update string
		Delete string
		GetID string
		LsContext string
	}
	File struct {
		ClientLinux64 string
		ClientMac string
		ClientWin64 string
		MaskConfig string
		MaskLinux64 string
		MaskMac string
		MaskWin64 string
		ClientConfig string
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
