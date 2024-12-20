## 解き方
与えられた公開鍵をid_rsa.pubとして保存し、pkcs8フォーマットに変換する
(今回使用する検証ライブラリがpkcs1もしくはpkcs8フォーマットの公開鍵を要求しているため c.f.https://pkg.go.dev/github.com/dgrijalva/jwt-go#ParseRSAPublicKeyFromPEM)
```bash
ssh-keygen -f id_rsa -e -m pkcs8 > id_rsa.pkcs8
```
main.goを実行する
プログラムの処理に関しては、main.goを参照

```bash
go run main.go

# 出力
verified JWT: &map[flag:f1nat3xthd{f53c73d9-278e-4b2e-af2e-26b10eed5224}]
```

## 回答
f1nat3xthd{f53c73d9-278e-4b2e-af2e-26b10eed5224}
