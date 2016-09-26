package main

import (
    "bufio"
    "encoding/base64"
    "fmt"
    "io/ioutil"
    "log"
    "log/syslog"
//    "net"
    "os"
    "strings"

    "golang.org/x/crypto/ssh"
)

var (
    hostPrivateKeySigner    ssh.Signer
    keyFile                 string
    port                    string
    ilevel                  string
)

func readConfig(config string) {
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
                port = str[1]
            } else if str[0] == "HONEY_KEY" {
                keyFile = str[1]
            } else if str[0] == "ILEVEL" {
                ilevel = str[1]
            }
        }
    }
}

func init() {
    logwriter, err := syslog.New(syslog.LOG_NOTICE, "honeyshell")
    if err == nil {
        log.SetOutput(logwriter)
    }
    config := ""
    if os.Getenv("HONEY_CONFIG") != "" {
        config = os.Getenv("HONEY_CONFIG")
    } else {
        log.Fatal("HONEY_CONFIG must be set")
    }
    readConfig(config)
    fmt.Printf("port = %s\n", port)
    fmt.Printf("key = %s\n", keyFile)
    fmt.Printf("ilevel = %s\n", ilevel)
    if keyFile == "" {
        //log.Fatal("HONEY_KEY must be set")
        fmt.Println("HONEY_KEY must be set")
    }
    hostPrivateKey, err := ioutil.ReadFile(keyFile)
    if err != nil {
        log.Fatal(err)
    }
    hostPrivateKeySigner, err = ssh.ParsePrivateKey(hostPrivateKey)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    /*
    config := ssh.ServerConfig {
        PublicKeyCallback: keyAuth,
        PasswordCallback: passAuth,
        ServerVersion: "SSH-2.0-OpenSSH_7.2p2 Ubuntu 4ubuntu2.1",
    }
    config.Config.SetDefaults()
    config.AddHostKey(hostPrivateKeySigner)
    socket, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatal(err)
    }
    for {
        conn, err := socket.Accept()
        if err != nil {
            log.Fatal(err)
        }
        sshConn, _, _, err := ssh.NewServerConn(conn, &config)
        if err != nil {
            log.Printf("Connection to sshoney from %v .. error: %v --END\n", conn.RemoteAddr().String(), err)
        } else {
            log.Printf("Connection on sshoney from %v --END\n", sshConn.RemoteAddr())
            sshConn.Close()
        }
    }
    */
}

func keyAuth(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
    log.Printf("%v attempted to connect as user %v with the key %v --END\n", conn.RemoteAddr(), conn.User(), base64.StdEncoding.EncodeToString(key.Marshal()))
    return nil, nil
}

func passAuth(conn ssh.ConnMetadata, psswd []byte) (*ssh.Permissions, error) {
    log.Printf("%v attempted to connect as user %v with the password %v --END\n", conn.RemoteAddr(), conn.User(), psswd)
    return nil, nil
}
