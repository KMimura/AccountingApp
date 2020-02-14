# 概要
以前、個人のAWSアカウントを用いて、[個人用のサーバーレスの家計簿webアプリ](https://qiita.com/KMim/items/f3975e308d07df4b359f)を作りました。
それを、EC2上のDocker上に移植します。
# 使用技術
### 静的ファイル
Vue.jsを用いて作成しました。これは、以前サーバーレスで作った際のモノをそのまま流用します。
### API
以前はLambda上で動いていたのですが、PythonのフレームワークのFlaskを用いるように書き換えました。
### DB
MySQLを用います。データはホストのEC2サーバーにマウンティングさせておきます。
## 自分用メモ
- [ターミナル放置で固まる件](https://rcmdnk.com/blog/2014/08/23/computer-linux-putty/)
- [GinでCRUD](https://qiita.com/hyo_07/items/59c093dda143325b1859)
