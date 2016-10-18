package config

import (
    "bufio"
    "log"
    "os"
    "strings"
)

var (
    KeyFile                 string
    Ilevel                  string
    LoggerName              string
    Port                    string
    SerVer                  string
)

func ReadConfig(config string) {
    conf, err := os.Open(config)
    if err != nil {
        log.Fatal(err)
    }
    defer conf.Close()
    reader := bufio.NewReader(conf)
    scanner := bufio.NewScanner(reader)
    for scanner.Scan() {
        line := scanner.Text()
        if line != "" && string(line[0]) != "#" {
            str := strings.Split(line, "=")
            if str[0] == "PORT" {
                Port = str[1]
            } else if str[0] == "HONEY_KEY" {
                KeyFile = str[1]
            } else if str[0] == "ILEVEL" {
                Ilevel = str[1]
            } else if str[0] == "VERSION" {
                SerVer = str[1]
            } else if str[0] == "LOGGER_NAME" {
                LoggerName = str[1]
            }
        }
    }
}

