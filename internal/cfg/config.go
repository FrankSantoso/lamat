package cfg

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

// Config is the main config file which you'll populate
type Config struct {
	APIKeys apiKeys
}

type apiKeys struct {
	Nominatim     string
	GoogleGeocode string
}

var (
	defaultConfig = map[string]interface{}{
		"apiKeys.Nominatim":     "PLEASE_INSERT_ANYTHING",
		"apiKeys.GoogleGeocode": "PLEASE_INSERT_ANYTHING",
	}
)

// ReadConfig will read input config path
// returning Config and error if any.
func ReadConfig(cfgpath string) (*Config, error) {
	var newConf Config
	vp, err := readConfig(cfgpath, defaultConfig)
	if err != nil {
		return nil, err
	}
	err = vp.Unmarshal(&newConf)
	return &newConf, err
}

// getConfigFile returns config-file-path and the name of config file.
func getConfigFile(filepath string) (string, string) {
	cfgSlice := strings.Split(filepath, "/")
	log.Printf("Path: %s, filename: %s",
		strings.Join(cfgSlice[:len(cfgSlice)-1], "/"),
		cfgSlice[len(cfgSlice)-1])
	return strings.Join(cfgSlice[:len(cfgSlice)-1], "/"),
		cfgSlice[len(cfgSlice)-1]
}

func readConfig(filepath string, defaults map[string]interface{}) (*viper.Viper, error) {
	vp := viper.New()
	cfgPath, filename := getConfigFile(filepath)
	for k, v := range defaults {
		vp.SetDefault(k, v)
	}
	vp.SetConfigName(filename)
	vp.AddConfigPath(cfgPath)
	vp.AutomaticEnv()
	err := vp.ReadInConfig()
	return vp, err
}
