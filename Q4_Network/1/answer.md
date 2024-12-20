## 解き方
d=st.fntxt.co;
これが送信元のドメイン。

s=dk4419;
これがセレクタ。

```bash
nslookup -type=TXT dk4419._domainkey.st.fntxt.co

# 出力
Server:         192.168.0.1
Address:        192.168.0.1#53

Non-authoritative answer:
dk4419._domainkey.st.fntxt.co   text = "v=DKIM1;k=rsa;p=f1nat3xthd{eb690513-0b2b-45a4-b6a3-591c56ac1a91}"

Authoritative answers can be found from:
```

参考
- https://www.proofpoint.com/jp/threat-reference/dkim
- https://baremail.jp/blog/2021/03/16/1124/

## 回答
f1nat3xthd{eb690513-0b2b-45a4-b6a3-591c56ac1a91}
