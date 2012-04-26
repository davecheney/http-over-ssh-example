// HTTP serving using your ssh server as a 
// frontend proxy -- it's 100% wacky
package main

import (
	"code.google.com/p/go.crypto/ssh"
	"fmt"
	"flag"
	"log"
	"net/http"
)

var (
        sshuser    = flag.String("ssh.user", "", "ssh username")
        sshpass    = flag.String("ssh.pass", "", "ssh password")
)

// password implements the ClientPassword interface
type password string

func (p password) Password(user string) (string, error) {
        return string(p), nil
}

func init() {
	flag.Parse()
}

func main() {
	config := &ssh.ClientConfig{
                User: *sshuser,
                Auth: []ssh.ClientAuth{
                        ssh.ClientAuthPassword(password(*sshpass)),
                },
        }
        conn, err := ssh.Dial("tcp", "localhost:22", config)
        if err != nil {
                log.Fatalf("unable to connect: %s", err)
        }
        defer conn.Close()

	l, err := conn.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("unable to register tcp forward: %v", err)
	}
	defer l.Close()
	http.Serve(l, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Hello world!\n")	
	}))
}
