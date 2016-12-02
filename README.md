# qiist

- 指定した Qiita User がストックしている投稿で、最新の 5 件を表示するコマンドです
- `config.json.sample` を基に `config.json` を作成してください
- `$ sh build.sh` でビルド
- `$ ./qiish` で実行できます



# 補足 - Qiita API の注意点
## 利用制限
- [Qiita API Doc | 利用制限](http://qiita.com/api/v2/docs#認証中のユーザ)

> 利用制限
> 認証している状態ではユーザごとに1時間に1000回まで、認証していない状態ではIPアドレスごとに1時間に60回までリクエストを受け付けます。

## Qiita:Team のデータを利用する場合
- [Qiita API Doc | 概要](http://qiita.com/api/v2/docs#概要)
> リクエスト
> APIとの全ての通信にはHTTPSプロトコルを利用します。アクセス先のホストには、Qiitaのデータを利用する場合には qiita.com を利用し、Qiita:Teamのデータを利用する場合には *.qiita.com を利用します (*には所属しているTeamのIDが入ります)。
