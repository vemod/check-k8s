package environment

import (
	"os"

	"github.com/spf13/pflag"
)

// EnvSettings describes all of the environment settings.
type EnvSettings struct {
	// Verbose indicates the level of additional debug output
	Verbose int
	// KubeContext is the name of the kubeconfig context.
	KubeContext string
}

// envMap maps flag names to envvars
var envMap = map[string]string{
	"debug": "CHECK_K8S_DEBUG",
}

// AddFlags binds flags to the given flagset.
func (s *EnvSettings) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.KubeContext, "kube-context", "", "name of the kubeconfig context to use")
	fs.CountVarP(&s.Verbose, "verbose", "v", "enable verbose output")
}

// Init sets values from the environment.
func (s *EnvSettings) Init(fs *pflag.FlagSet) {
	for name, envar := range envMap {
		setFlagFromEnv(name, envar, fs)
	}
}

func setFlagFromEnv(name, envar string, fs *pflag.FlagSet) {
	if fs.Changed(name) {
		return
	}
	if v, ok := os.LookupEnv(envar); ok {
		fs.Set(name, v)
	}
}
