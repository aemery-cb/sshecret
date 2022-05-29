package server

import (
	"io"

	"github.com/aemery-cb/sshecret/config"
	"github.com/aemery-cb/sshecret/secret"
	"github.com/gliderlabs/ssh"
)

func NewServer(conf *config.Config, sm *secret.SecretManager) *ssh.Server {

	server := ssh.Server{
		Addr: ":2222",
		Handler: ssh.Handler(func(s ssh.Session) {
			commands := s.Command()
			if len(commands) == 0 {
				return
			}
			switch s.Command()[0] {
			case "put":
				err := sm.Put(s.Command()[1], s.Command()[2])
				if err != nil {
					io.WriteString(s, "Error\n")
				} else {
					io.WriteString(s, "Done\n")
				}

			case "get":
				val, err := sm.Get(s.Command()[1])
				if err != nil {
					io.WriteString(s, "")
				} else {
					io.WriteString(s, val+"\n")
				}
			default:
				io.WriteString(s, "Help string\n")
			}
			val, _ := sm.Get(s.Command()[0])
			io.WriteString(s, val)
		}),
		PublicKeyHandler: ssh.PublicKeyHandler(func(ctx ssh.Context, key ssh.PublicKey) bool {
			for k, v := range conf.AuthorizedKeys {
				if k != ctx.User() {
					continue
				}
				stored, _, _, _, err := ssh.ParseAuthorizedKey([]byte(v))
				if err != nil {
					return false
				}
				if ssh.KeysEqual(key, stored) {
					return true
				}
			}
			return false
		}),
	}

	return &server
}
