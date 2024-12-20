## 解き方
コンテナ実行しただけではflag.txtの内容は参照できないので、flag.txtの内容を出力するコマンドを実行するようにコンテナを起動する。

```bash
docker run --rm --entrypoint cat stajima/skill-test:v1_q1-2 flag.txt

# 出力
f1nat3xthd{b82df727-593a-46cc-9949-dabe8b8c10a7}
```

## 回答
f1nat3xthd{b82df727-593a-46cc-9949-dabe8b8c10a7}
