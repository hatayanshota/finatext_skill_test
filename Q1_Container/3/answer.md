## 解き方
### 失敗
2と同じやり方でやってみる

```bash
docker run --rm --entrypoint cat stajima/skill-test:v1_q1-3 flag.txt

# 出力
docker: Error response from daemon: failed to create task for container: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: exec: "cat": executable file not found in $PATH: unknown.
```

どうやらこのイメージのベースイメージにはcatコマンドが入っていない。

### 成功
一度イメージからコンテナを作成し、`docker cp`コマンドを使用してコンテナ内のファイルをホストマシンにコピーするやり方を試してみる。

```bash
docker create --name skill-test stajima/skill-test:v1_q1-3
docker cp skill-test:/flag.txt ./flag.txt
cat flag.txt

# 出力
f1nat3xthd{7d374833-4dcb-4a66-9aa7-8c3a7ff737de}
```

## 回答
f1nat3xthd{7d374833-4dcb-4a66-9aa7-8c3a7ff737de}
