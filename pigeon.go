package main

import (
	"bitbucket.org/chrj/smtpd"
//	"errors"
	"fmt"
	"strings"
)

func main() {
	var server *smtpd.Server

	server = &smtpd.Server{

		HeloChecker: func(peer smtpd.Peer, name string) error {
			if !strings.HasPrefix(peer.Addr.String(), "127.0.0.1:") {
				//return errors.New("Denied")
			}
			return nil
		},

		Handler: func(peer smtpd.Peer, env smtpd.Envelope) error {
			fmt.Printf("Sender: %s\nRecipients: %s\n%s\n\n",
				env.Sender,
				env.Recipients,
				env.Data)
			return nil
		},
	}

	server.ListenAndServe("172.16.1.164:10025")
}
