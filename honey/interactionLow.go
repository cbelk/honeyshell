package honey

import (
    "fmt"

    "github.com/cbelk/honeyshell/config"
)

func HoneyLow() {
/*
    config := ssh.ServerConfig {
        PublicKeyCallback: KeyAuthLow,
        PasswordCallback: PassAuthLow,
        ServerVersion: SerVer,
    }
    config.Config.SetDefaults()
    config.AddHostKey(hostPrivateKeySigner)
    socket, err := net.Listen("tcp", ":"+honeyshell.Port)
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
    fmt.Printf("Server Version -- %v\n", config.SerVer)
}
