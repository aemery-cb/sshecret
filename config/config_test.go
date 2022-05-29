package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gliderlabs/ssh"
)

func TestServerWriteAndRead(t *testing.T) {
	key := "asldkfjalskdjfla;skdfjl;askdjfal;ksdjfl;aksjdfl;kasjdf"
	configPath := "config.toml"
	stored, _, _, _, err := ssh.ParseAuthorizedKey([]byte(key))
	if err != nil {
		t.Error(err)
	}
	conf := Config{
		ConfigPath: configPath,
		AuthorizedKeys: map[string]string{
			"secureKey": string(stored.Marshal()),
		},
	}

	err = conf.Write()
	if err != nil {
		t.Error(err)
	}
	conf2 := Config{ConfigPath: configPath}
	conf2.Read()

	fmt.Print(conf)
	fmt.Print(conf2)
	if !reflect.DeepEqual(conf, conf2) {
		t.Error("SNOT EQUAL")
	}

}
