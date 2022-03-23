package xxtea

import (
	"bytes"
	"testing"
	"time"
)

func zeroBytes(n int) []byte {
	return make([]byte, n)
}

var data = []byte("xxtea-go test case")
var key = []byte{82, 253, 252, 7, 33, 130, 101, 79, 22, 63, 95, 15, 154, 98, 29, 114}

var enc = []byte{165, 49, 49, 102, 222, 29, 124, 20, 219, 59, 80, 14, 80, 113, 186, 239, 7, 66, 98, 216}
var hexEnc = "a5313166de1d7c14db3b500e5071baef074262d8"
var b64Enc = "pTExZt4dfBTbO1AOUHG67wdCYtg="

func TestEncrypt(t *testing.T) {
	v, err := Encrypt(data, key, true, 0)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(enc, v) {
		t.Errorf("%+v != %+v", enc, v)
	}
}

func TestDecrypt(t *testing.T) {
	v, err := Decrypt(enc, key, true, 0)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, v) {
		t.Errorf("%+v != %+v", data, v)
	}
}

func TestEncryptBase64(t *testing.T) {
	v, err := EncryptBase64(data, key, true, 0)
	if err != nil {
		t.Error(err)
	}

	if b64Enc != v {
		t.Errorf("%s != %s", hexEnc, v)
	}
}

func TestDecryptBase64(t *testing.T) {
	v, err := DecryptBase64(b64Enc, key, true, 0)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, v) {
		t.Errorf("%+v != %+v", data, v)
	}
}

func TestEncryptHex(t *testing.T) {
	v, err := EncryptHex(data, key, true, 0)
	if err != nil {
		t.Error(err)
	}

	if hexEnc != v {
		t.Errorf("%s != %s", hexEnc, v)
	}
}

func TestDecryptHex(t *testing.T) {
	v, err := DecryptHex(hexEnc, key, true, 0)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, v) {
		t.Errorf("%+v != %+v", data, v)
	}
}

func TestRandom(t *testing.T) {
	for i := 0; i < 2048; i++ {
		seedV := time.Now().UnixNano()
		data, _ := URandom(i, seedV)
		key, _ := URandom(16, seedV)

		enc, err := Encrypt(data, key, true, 0)
		if err != nil {
			t.Error(err)
		}

		dec, err := Decrypt(enc, key, true, 0)
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(data, dec) {
			t.Errorf("data=%+v not equal dec=%+v, key=%+v\n", data, dec, key)
		}

		key = zeroBytes(16)
		enc, err = Encrypt(data, key, true, 0)
		if err != nil {
			t.Error(err)
		}
		dec, err = Decrypt(enc, key, true, 0)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(data, dec) {
			t.Errorf("data=%+v not equal dec=%+v, key=%+v\n", data, dec, key)
		}
	}
}

func TestZeroBytes(t *testing.T) {
	for i := 0; i < 2048; i++ {
		data := zeroBytes(i)
		key, _ := URandom(16, time.Now().UnixNano())

		enc, err := Encrypt(data, key, true, 0)
		if err != nil {
			t.Error(err)
		}
		dec, err := Decrypt(enc, key, true, 0)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(data, dec) {
			t.Errorf("data=%+v not equal dec=%+v, key=%+v\n", data, dec, key)
		}
	}
}

func TestEncryptNoPadding(t *testing.T) {
	seedV := time.Now().UnixNano()
	key, _ := URandom(16, seedV)
	for _, v := range []int{8, 12, 16, 20} {
		data, _ := URandom(v, seedV)
		enc, err := Encrypt(data, key, false, 0)
		if err != nil {
			t.Error(err)
		}

		dec, err := Decrypt(enc, key, false, 0)
		if err != nil {
			t.Error(err)
		}

		//t.Logf("log: data=%+v, enc=%+v, dec=%+v", data, enc, dec)
		if !bytes.Equal(data, dec) {
			t.Errorf("data=%+v not equal dec=%+v, key=%+v\n", data, dec, key)
		}
	}
}

func TestEncryptRandomRounds(t *testing.T) {
	seedV := time.Now().UnixNano()
	key, _ := URandom(16, seedV)
	data, _ := URandom(64, seedV)
	for i := 1; i < 2048; i++ {
		enc, err := Encrypt(data, key, true, uint32(i))
		if err != nil {
			t.Error(err)
		}

		dec, err := Decrypt(enc, key, true, uint32(i))
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(data, dec) {
			t.Errorf("data=%+v not equal dec=%+v, key=%+v\n", data, dec, key)
		}
	}
}

func TestEncryptNoPaddingZero(t *testing.T) {
	seedV := time.Now().UnixNano()
	key, _ := URandom(16, seedV)
	for _, v := range []int{8, 12, 16, 20} {
		data := zeroBytes(v)

		enc, err := Encrypt(data, key, false, 0)
		if err != nil {
			t.Error(err)
		}

		dec, err := Decrypt(enc, key, false, 0)
		if err != nil {
			t.Error(err)
		}

		//t.Logf("log: data=%+v, enc=%+v, dec=%+v", data, enc, dec)
		if !bytes.Equal(data, dec) {
			t.Errorf("data=%+v not equal dec=%+v, key=%+v\n", data, dec, key)
		}
	}
}
