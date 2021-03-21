package config

import (
	"basic/memutils/redis"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-sql-driver/mysql"
)

type site struct {
	SiteName        string `json:"sitename"`
	Description     string `json:"description"`
	URL             string `json:"url"`
	Keywords        string `json:"keywords"`
	DefaultTimezone string `json:"defaultTimezone"`
}

type configuration struct {
	RedisConfig redis.Config `json:"redis"`
	MySqlConfig mysql.Config `json:"mysql"`
	Site        site         `json"site"`
}

//Current the current configuration
var Current configuration

//Init init the configuration
func init() {
	Current = configuration{}

	basePath, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(basePath)
	configFile := filepath.Join(filepath.Dir(path), "config.json")
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	file, err := os.Open(configFile)

	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Current)
	if err != nil {
		fmt.Print(err)
		return
	}
}

//LoadFromFile load config from file
func LoadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&Current)
}
