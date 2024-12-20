## 解き方
指定のエンドポイントに対してcurlで`POST`リクエストを送信し、レスポンスを受け取る。
```bash
curl -X POST https://skill-test.st8.workers.dev/v1/q3-1/login -d '{"username":"admin","password":"supersecretpassword"}'

# 出力
{"token": "d97329d6c03c90a0c82bfec27e9427e796324861", "flag": "f1nat3xthd{8727d491-8190-4a3a-a937-d94a71d89c6c}"}
```

## 回答
f1nat3xthd{8727d491-8190-4a3a-a937-d94a71d89c6c}
