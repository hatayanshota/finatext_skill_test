## 解き方
与えられた秘密鍵をid_rsaとして保存し、pkcs8フォーマットに変換する
(今回使用する検証ライブラリがpkcs1もしくはpkcs8フォーマットの公開鍵を要求しているため c.f.https://pkg.go.dev/github.com/dgrijalva/jwt-go#ParseRSAPrivateKeyFromPEM)
```bash
openssl pkcs8 -topk8 -inform PEM -in id_rsa -outform PEM -out id_rsa.pkcs8 -nocrypt
```
main.goを実行する
プログラムの処理に関しては、main.goを参照

```bash
go run main.go

# 出力
verified JWT: &map[flag:f1nat3xthd{0ad9dc6b-9e07-4e0f-b142-ff170e6d2ac1}]
```

## 回答
f1nat3xthd{0ad9dc6b-9e07-4e0f-b142-ff170e6d2ac1}
