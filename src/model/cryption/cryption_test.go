package cryption

import(
	"testing"
)

func TestCryp(t *testing.T){
	username := "Jack88552"
	//password := "TestBBBBBB"
	password := "yo8@%gur$T"
	crypt := EncryptionByString(password, username)
	decrypt := Decryption(crypt, username)
	if password != decrypt {
		t.Errorf("Bad Crypt: %s should be %s. Crypt: %s", decrypt, password, crypt)
	}
}

func TestRandBytes(t *testing.T){
	
}
