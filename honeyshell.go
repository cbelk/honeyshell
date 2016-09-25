package main

import (
    "encoding/base64"
    "io/ioutil"
    "log"
    "log/syslog"
    "net"
    "os"

    "golang.org/x/crypto/ssh"
)

const (
    port string = "22"
)

var (
    hostPrivateKeySigner    ssh.Signer
    keyFile                 string
    port                    string
    keyPath                 string
)

func init() {
    logwriter, err := syslog.New(syslog.LOG_NOTICE, "sshoney")
    if err == nil {
        log.SetOutput(logwriter)
    }
    config := ""
    if os.Getenv("HONEY_CONFIG") != "" {
        config = os.Getenv("HONEY_CONFIG")
    } else {
        log.Fatal("HONEY_CONFIG must be set")
    }
    keyPath := ""
    if os.Getenv("HONEY_KEY") != "" {
        keyPath = os.Getenv("HONEY_KEY")
    }
    if keyPath == "" {
        log.Fatal("HONEY_KEY must be set")
    }
    hostPrivateKey, err := ioutil.ReadFile(keyPath)
    if err != nil {
        log.Fatal(err)
    }
    hostPrivateKeySigner, err = ssh.ParsePrivateKey(hostPrivateKey)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
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
}

func keyAuth(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
    log.Printf("%v attempted to connect as user %v with the key %v --END\n", conn.RemoteAddr(), conn.User(), base64.StdEncoding.EncodeToString(key.Marshal()))
    return nil, nil
}

func passAuth(conn ssh.ConnMetadata, psswd []byte) (*ssh.Permissions, error) {
    log.Printf("%v attempted to connect as user %v with the password %v --END\n", conn.RemoteAddr(), conn.User(), psswd)
    return nil, nil
}
