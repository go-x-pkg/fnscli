package fnscli

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

func IsPFlagSet(flags *pflag.FlagSet, name string) bool {
	found := false

	flags.Visit(func(f *pflag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}

func DecodeYAMLFromPath(into interface{}, path string, isPathFromArguments bool) error {
	_, err := os.Stat(path)

	// raise error if path to file was directly set via
	// command line arguments
	if isPathFromArguments && err != nil {
		return fmt.Errorf("read config file passed from arguments failed (:path %s): %w",
			path, err)
	}

	// try to read config file only if it is exists
	// if file is not exists - do nothing
	if os.IsNotExist(err) {
		return nil
	}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config file failed (:path %s): %w", path, err)
	}

	if err := yaml.Unmarshal(raw, into); err != nil {
		return fmt.Errorf("decode config file from yaml failed (:path %s): %w", path, err)
	}

	return nil
}
