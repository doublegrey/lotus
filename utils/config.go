package utils

import "github.com/BurntSushi/toml"

// Config TomlConfig
var Config TomlConfig

// TomlConfig struct
type TomlConfig struct {
	Server struct {
		Version     string `toml:"version"`
		Addr        string `toml:"addr"`
		Development bool   `toml:"development"`
	} `toml:"server"`
	Database struct {
		URI     string `toml:"uri"`
		DB      string `toml:"db"`
		Timeout int    `toml:"timeout"`
	} `toml:"database"`
}

// ParseConfig decodes toml config to Config
func ParseConfig(p ...string) {
	path := "./config.toml"
	if len(p) > 0 {
		path = p[0]
	}
	if _, err := toml.DecodeFile(path, &Config); err != nil {
		panic(err)
	}
}
