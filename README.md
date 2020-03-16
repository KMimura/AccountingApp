# 概要
以前、個人のAWSアカウントを用いて、[個人用の家計簿webアプリ](https://qiita.com/KMim/items/f3975e308d07df4b359f)を、サーバーレスで作りました。
それを今回、EC2上のDocker上に移植します。
# 使用技術
### 静的ファイル
Vue.jsを用いて作成しました。これは、以前サーバーレスで作った際のモノをそのまま流用します。
### API
以前はPythonのプログラムがLambda上で動いていたのですが、今回はそれをGo言語のフレームワークであるGinを用いたものに書き換えました。
### DB
MySQLを用います。データはホストのEC2サーバーにマウンティングさせておきます。
## 自分用メモ
- [GinでCRUD](https://qiita.com/hyo_07/items/59c093dda143325b1859)
- [docker-composeでアプリからDBにアクセスする](https://qiita.com/M_Nagata/items/120831bb4e4a3deace13)