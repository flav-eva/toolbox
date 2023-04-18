package crypto

import "testing"

func TestHashPassword(t *testing.T) {
	pwd := "123456"
	hashPassword, err := HashPassword(pwd)
	if err != nil {
		t.Logf("err: %v", err)
		panic(err)
	}
	t.Logf("hashPassword: %v", hashPassword)

	verify, err := VerifyPassword(pwd, hashPassword)
	if err != nil {
		panic(err)
	}

	t.Logf("verify: %v", verify)
}

func TestMD5(t *testing.T) {
	pwd := "123456"
	md5Pwd := MD5(pwd)
	t.Logf("ret: %v", MD5Verify(md5Pwd, pwd))
}
