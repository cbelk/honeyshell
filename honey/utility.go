package honey

import (
    "encoding/base64"
    "log"
    "log/syslog"

    "github.com/cbelk/honeyshell/config"
    "golang.org/x/crypto/ssh"
)

func init() {
    logwriter, err := syslog.New(syslog.LOG_NOTICE, config.LoggerName)
    if err == nil {
        log.SetOutput(logwriter)
    }
}

func KeyAuthLow(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
    log.Printf("%v attempted to connect as user %v with the key %v --END\n", conn.RemoteAddr(), conn.User(), base64.StdEncoding.EncodeToString(key.Marshal()))
    return nil, nil
}

func PassAuthLow(conn ssh.ConnMetadata, psswd []byte) (*ssh.Permissions, error) {
    log.Printf("%v attempted to connect as user %v with the password %v --END\n", conn.RemoteAddr(), conn.User(), psswd)
    return nil, nil
}
