## 解き方
(今回のサーバーの実装は問題2のものとまとめて実装しています)
server配下でmain.goを実行してサーバーを立ち上げる。
指定のAPIエンドポイントにリクエストを送信する。
今回も、ローカルで立ち上げたサーバーが指定のサーバーからリクエストを受け付けられるようにするため、ngrokを使用した。
本サーバーのURLは以下とする。
https://example.com/

```bash
cd ../server
go run main.go

# 別シェルで実行
curl -X POST https://skill-test.st8.workers.dev/v1/q3-3/agent -d '{"target":"https://example.com/api/"}'
# 出力
flag: f1nat3xthd{02f401fd-140e-4fe5-a7f5-29ff996139fc}
```

## 回答
f1nat3xthd{02f401fd-140e-4fe5-a7f5-29ff996139fc}
