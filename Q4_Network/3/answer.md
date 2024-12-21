## 解き方
main.goを実行してパケットのアプリケーションレイヤーのペイロードを眺める。
プログラムの処理に関しては、main.goを参照
パケットを解析するにあたってgithub.com/google/gopacketライブラリを使用する。

```bash
go run main.go

# 出力
2024/12/21 18:02:12 payload: J
8.0.3U%'860v�����k!g,O
                      Cr^CsVmysql_native_password
2024/12/21 18:02:12 payload: ���finatext��dh��q���U��;�x�C�mysql_native_passwordp_pid94 _platformx86_64_osLinux
                                                                                                               _client_namlibmysqlos_userroot_client_version8.0.31
                                      program_namemysql
2024/12/21 18:02:12 payload: 
2024/12/21 18:02:12 payload: #select @@version_comment limit 1
2024/12/21 18:02:12 payload: 'def@@version_comment
                                                 UU�MySQL Community Server - GPL�
2024/12/21 18:02:12 payload: show databases
2024/12/21 18:02:12 payload: 6deSCHEMATschematDatabasDatabase
                                                            @��information_schemaperformance_schema

skill_test�"
2024/12/21 18:02:12 payload: SELECT DATABASE()
2024/12/21 18:02:12 payload:  def
DATABASE()
         "���
2024/12/21 18:02:12 payload: 
                             skill_test
2024/12/21 18:02:12 payload: @

skill_test
2024/12/21 18:02:12 payload: show tables
2024/12/21 18:02:12 payload: JdefTABLEStablesTables_in_skill_testTables_in_skill_test
                                                                                    @��flag�
2024/12/21 18:02:12 payload: select hex(flag) from flag
2024/12/21 18:02:12 payload: def        hex(flag)
                                                ��a`626F6775735F666C61677B61353335616335372D356639322D346539392D623636302D3736393239653934636563387Da`66316E617433787468647B66623439613661372D646132632D346139342D616561352D3936373032333061363063337Da`74686973697366616B657B36383566386561362D386532312D343133352D396164652D3039393338373039326263377Da`69676E6F7265746869737B36663165613432302D613961332D343666622D383366392D3238343935336563373465317D�"
2024/12/21 18:02:12 payload: 
```

上記の`select hex(flag) from flag`の結果にflagが入ってそうなことがわかる
したがって下記をそれぞれhexデコードしてみる
- 626F6775735F666C61677B61353335616335372D356639322D346539392D623636302D3736393239653934636563387D
- 66316E617433787468647B66623439613661372D646132632D346139342D616561352D3936373032333061363063337D
- 74686973697366616B657B36383566386561362D386532312D343133352D396164652D3039393338373039326263377D
- 69676E6F7265746869737B36663165613432302D613961332D343666622D383366392D3238343935336563373465317D

結果
- boggus_flag{a535ac57-5f92-4e99-b660-76929e94cec8}
- f1nat3xthd{fb49a6a7-da2c-4a94-aea5-9670230a60c3}
- thisisfake{685f8ea6-8e21-4135-9ade-099387092bc7}
- ignorethis{6f1ea420-a9a3-46fb-83f9-284953ec74e1}

## 回答
f1nat3xthd{fb49a6a7-da2c-4a94-aea5-9670230a60c3}
