package gonfig

// Config : AUTO Created From "gonfig/toml/test.toml"
type Config struct {
	IP interface{}
	LogFile string
	Path interface{}
	Port int
	Service interface{}
	Storage struct {
		BadgerDBPath string
		DataBase string
		MetaPath string
	}
	WebService struct {
		Version interface{}
		Port interface{}
		Service string
	}
	Route struct {
		GetHash string
		LsContext string
		LsID string
		Update string
		Delete string
		Get string
		GetID string
		LsObject string
		LsUser string
		ROOT string
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
		Delete string
		Get string
		GetHash string
		GetID string
		LsObject string
		ROOT string
		LsContext string
		LsID string
		LsUser string
		Update string
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
