package go_jwt

import "testing"

func init() {
	Secret = "123456"
}

func TestSign(t *testing.T) {
	token, _ := Sign(struct {
		Name string `json:"name"`
	}{Name: "yunze"})
	t.Log(token)

	err := Verify(token)
	if err != nil {
		t.Error("failed")
	} else {
		t.Log("ok")
	}
}
