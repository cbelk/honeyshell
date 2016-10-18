package honey

import (
    "fmt"
    "log"
    "net"

    "github.com/cbelk/honeyshell/config"
    "golang.org/x/crypto/ssh"
)

func HoneyLow(logger *log.Logger, hostPrivateKeySigner ssh.Signer) {
    conf := ssh.ServerConfig {
        PublicKeyCallback: KeyAuthLow,
        PasswordCallback: PassAuthLow,
        ServerVersion: config.SerVer,
    }
    conf.Config.SetDefaults()
    conf.AddHostKey(hostPrivateKeySigner)
    socket, err := net.Listen("tcp", ":"+config.Port)
    if err != nil {
        logger.Fatal(err)
    }
    fmt.Printf("Server Version -- %v\n", config.SerVer)
    for {
        conn, err := socket.Accept()
        if err != nil {
            logger.Fatal(err)
        }
        sshConn, _, _, err := ssh.NewServerConn(conn, &conf)
        if err != nil {
            logger.Printf("Connection to sshoney from %v .. error: %v --END\n", conn.RemoteAddr().String(), err)
        } else {
            logger.Printf("Connection on sshoney from %v --END\n", sshConn.RemoteAddr())
            sshConn.Close()
        }
    }
}
