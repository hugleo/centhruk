package main

import (
    "os"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Server struct {
        Host string `yaml:"host"`
        Port string `yaml:"port"`
    } `yaml:"server"`
    
    Database struct {
        Name string `yaml:"name"`
        Host string `yaml:"host"`
        Port string `yaml:"port"`
        Username string `yaml:"username"`
        Password string `yaml:"password"`
    } `yaml:"database"`
}

func get_config() (*Config) {
    config := &Config{}

    file, err := os.Open("/etc/centhruk.conf")
    if err != nil {
        return nil
    }
    defer file.Close()

    d := yaml.NewDecoder(file)

    if err := d.Decode(&config); err != nil {
        return nil
    }

    return config
}
