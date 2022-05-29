package secret

import "testing"

func TestSecretManagerGetAndPut(t *testing.T) {
	sm := NewSecretManager()
	err := sm.Put("hello", "its me")
	if err != nil {
		t.Error(err)
	}
	value, err := sm.Get("hello")
	if err != nil {
		t.Error(err)
	}
	if value != "its me" {
		t.Error("fail")
	}
}
