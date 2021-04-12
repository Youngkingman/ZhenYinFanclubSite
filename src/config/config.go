package config

import (
	"basic/memutils/mysql"
	"basic/memutils/redis"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	Site        site         `json:"site"`
	Debug       bool         `json:"debug"`
}

//Current the current configuration
var Current configuration

//Init init the configuration
func init() {
	Current = configuration{
		RedisConfig: redis.Config{
			Host:     "",
			Port:     0,
			Auth:     "",
			Db:       0,
			PoolSize: 0,
		},
		MySqlConfig: mysql.Config{
			Host:         "127.0.0.1",
			Port:         3306,
			User:         "root",
			Password:     "root",
			Db:           "zhenyinfanclub",
			Dbprefix:     "czy_",
			ConnLifeTime: 40,
			MaxIdleConn:  2,
			MaxOpenConn:  30,
		},
		Site: site{
			SiteName:        "",
			Description:     "",
			URL:             "",
			Keywords:        "",
			DefaultTimezone: "",
		},
		Debug: false,
	}

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
