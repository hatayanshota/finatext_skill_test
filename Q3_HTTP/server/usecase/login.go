package usecase

import (
	"crypto/sha1"
	"encoding/hex"
)

type LoginParam struct {
	Username string
	Password string
}

type LoginResult struct {
	Token string
}

func Login(param LoginParam) LoginResult {
	// username と password を結合
	combined := param.Username + param.Password

	// SHA1のチェックサムを計算
	hash := sha1.New()
	hash.Write([]byte(combined))
	checksum := hash.Sum(nil)

	// チェックサムを16進数文字列に変換
	checksumHex := hex.EncodeToString(checksum)

	return LoginResult{
		Token: checksumHex,
	}
}
