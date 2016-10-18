package main

import (
//    "bufio"
    "bytes"
    "fmt"
    "io/ioutil"
    "log"
    "log/syslog"
//    "net"
    "os"
//    "strings"

    "github.com/cbelk/honeyshell/honey"
    "github.com/cbelk/honeyshell/config"
    "golang.org/x/crypto/ssh"
)

var (
    hostPrivateKeySigner    ssh.Signer
    buf                     bytes.Buffer
    logger                  *log.Logger
)

func init() {
    conf := ""
    if os.Getenv("HONEY_CONFIG") != "" {
        conf = os.Getenv("HONEY_CONFIG")
    } else {
        log.Fatal("HONEY_CONFIG must be set")
    }
    config.ReadConfig(conf)
    logger = log.New(&buf, config.LoggerName, log.Lshortfile)
    logwriter, err := syslog.New(syslog.LOG_NOTICE, config.LoggerName)
    if err == nil {
        logger.SetOutput(logwriter)
    }
    if config.KeyFile == "" {
        //logger.Fatal("HONEY_KEY must be set")
        fmt.Println("HONEY_KEY must be set")
    }
    hostPrivateKey, err := ioutil.ReadFile(config.KeyFile)
    if err != nil {
        logger.Fatal(err)
    }
    hostPrivateKeySigner, err = ssh.ParsePrivateKey(hostPrivateKey)
    if err != nil {
        logger.Fatal(err)
    }
}

func main() {
    if config.Ilevel == "LOW" {
        honey.HoneyLow(logger, hostPrivateKeySigner)
    }
}
