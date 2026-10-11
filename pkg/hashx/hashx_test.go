package hashx

import (
	"encoding/hex"
	"testing"
)

func TestHashSHA256(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "my-sha-256-testing-string",
			want: "7a014a26d6672054025c348e78eee8ec8d381dd249fa52fbf122bd3ddfb4f1c5",
		},
		{
			in:   "/file/spajz/random-path/this-is-certainly-image.jpg",
			want: "2b504c3192031359de369f056c5564a0845df1d04bdc4a05ecf0b8acb8118357",
		},
	}

	for _, test := range tests {
		hash := hex.EncodeToString(HashSHA256([]byte(test.in)))
		if hash != test.want {
			t.Errorf("SHA256 hash for '%s' is incorrect '%s', want: '%s'", test.in, hash, test.want)
		}
	}
}

func TestHashMD5(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "my-md5-testing-string",
			want: "ac3feb8f6f4c9a4ee44694bef061485c",
		},
		{
			in:   "/file/spajz/random-path/this-is-certainly-image.jpg",
			want: "16e3cf586b19043c4db9ec1992bad156",
		},
	}

	for _, test := range tests {
		hash := hex.EncodeToString(HashMD5([]byte(test.in)))
		if hash != test.want {
			t.Errorf("MD5 hash for '%s' is incorrect '%s', want: '%s'", test.in, hash, test.want)
		}
	}
}
