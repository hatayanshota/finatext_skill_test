package main

import (
	"crypto/rsa"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

const (
	pubKeyFile = "id_rsa.pkcs8"
	jwtFile    = "jwts.rand.txt"
)

func main() {
	// 公開鍵ファイルを読み込む
	pubKey, err := os.ReadFile(pubKeyFile)
	if err != nil {
		log.Fatalf("failed to read jwt file: %v", err)
	}

	// 公開鍵をパース
	parsedPubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		log.Fatalf("failed to parse public key: %v", err)
	}

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
		claims, err := verifyJWT(token, parsedPubKey)
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
