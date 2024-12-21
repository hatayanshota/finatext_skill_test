## 解き方
server配下でmain.goを実行してサーバーを立ち上げる。
指定のAPIエンドポイントにリクエストを送信する。
今回は、ローカルで立ち上げたサーバーが指定のサーバーからリクエストを受け付けられるようにするため、ngrokを使用した。
本サーバーのURLは以下とする。
https://example.com/

```bash
cd ../server
go run main.go

# 別シェルで実行
curl -X POST https://skill-test.st8.workers.dev/v1/q3-2/agent -d '{"target":"https://example.com/api/"}'
# 出力
flag: f1nat3xthd{f1096ca4-f298-464f-97e0-feedbe237f78}
```

## 回答
f1nat3xthd{f1096ca4-f298-464f-97e0-feedbe237f78}
