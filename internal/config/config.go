package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Read() (Config, error) {
	var cfg Config
	cfgPath, err := getConfigFilePath()
	if err != nil {
	    return Config{}, err
	} 

	cfgJson, err := os.ReadFile(cfgPath)
	if err != nil {
	    return Config{}, err
	} 

	if err := json.Unmarshal(cfgJson, &cfg); err != nil {
	    return Config{}, err
	} 
	return cfg, nil
} 

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
} 

func write(cfg Config) error {
	cfgJson, err := json.Marshal(cfg)
	if err != nil {
	    return err
	} 

	configPath, err := getConfigFilePath()
	if err != nil {
	    return err
	} 

	err = os.WriteFile(configPath, cfgJson, os.ModeAppend)
	if err != nil {
		return err
	} 
	return nil
} 

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
	    return "", nil
	} 
	return filepath.Join(home, configFileName), err
} 
