package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"log"
	"time"

	"github.com/aemery-cb/sshecret/config"

	"github.com/aemery-cb/sshecret/secret"
	"github.com/aemery-cb/sshecret/server"
	_ "modernc.org/sqlite"
)

func main() {

	conf, err := config.NewConfigFrom("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	sm := secret.NewSecretManager()
	srv := server.NewServer(conf, sm)
	log.Fatal(srv.ListenAndServe())

}

// Append extra 0s if the length of otp is less than 6
// If otp is "1234", it will return it as "001234"
func prefix0(otp string) string {
	if len(otp) == 6 {
		return otp
	}
	for i := (6 - len(otp)); i > 0; i-- {
		otp = "0" + otp
	}
	return otp
}

func hotp(key string, counter uint64, digits int) int {
	bytes := []byte(key)
	h := hmac.New(sha1.New, bytes)
	binary.Write(h, binary.BigEndian, counter)
	sum := h.Sum(nil)
	v := binary.BigEndian.Uint32(sum[sum[len(sum)-1]&0x0F:]) & 0x7FFFFFFF
	d := uint32(1)
	for i := 0; i < digits && i < 8; i++ {
		d *= 10
	}
	return int(v % d)
}

func totp(key string, t time.Time, digits int) int {
	return hotp(key, uint64(t.UnixNano())/30e9, digits)
}
