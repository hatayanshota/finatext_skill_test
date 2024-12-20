package main

import (
	"crypto/rsa"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

const (
	privateKeyFile = "id_rsa.pkcs8"
	jwtFile        = "jwts.rand.txt"
)

func main() {
	// 秘密鍵ファイルを読み込む
	privateKey, err := os.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatalf("failed to read jwt file: %v", err)
	}

	// 秘密鍵をパース
	parsedPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		log.Fatalf("failed to parse public key: %v", err)
	}

	publicKey := parsedPrivateKey.PublicKey

	// jwts.rand.txtを読み込む
	jwts, err := os.ReadFile(jwtFile)
	if err != nil {
		log.Fatalf("failed to read jwt file: %v", err)
	}

	// 各JWTを検証
	tokens := strings.Split(string(jwts), "\n")
	for _, token := range tokens {
		if token == "" {
			continue
		}

		// 検証
		claims, err := verifyJWT(token, &publicKey)
		if err != nil {
			log.Fatalf("failed to verify JWT: %v", err)
		}

		if claims != nil {
			// 検証結果が入っていればログに出力して終了
			log.Printf("verified JWT: %v", claims)
			break
		}
	}
}

func verifyJWT(tokenString string, pubKey *rsa.PublicKey) (*jwt.MapClaims, error) {
	// JWTを解析
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})
	if err != nil {
		return nil, nil
	}

	// 検証結果を取得
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil
	}

	return &claims, nil
}
