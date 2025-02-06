## 使用技術
- Gazelle:
<br>Bazel を利用する際に、ソースコードのディレクトリ構造や依存関係を解析して、適切な BUILD ファイルを生成・更新するツール

- Bzlmod
<br>Bazel における新しい外部依存関係管理システム。従来の WORKSPACE ファイルを使った依存管理に代わる仕組みとして導入。MODULE.bazel というファイルに外部モジュールの依存関係やバージョン情報などを記述。


## 利用コマンド
### BUILD ファイルの自動生成・更新
-  Go のパッケージ構造やソースコードに変更があったとき
- 手動での BUILD ファイルのメンテナンス負担を軽減したい場合
- 新たにファイルやパッケージを追加した後、ビルドルールを最新状態に反映させたいとき
```
bazel run //:gazelle
```
<br>

### Bzlmod（Bazelモジュールシステム）の依存関係整理
- Bazel の外部依存関係（Bzlmod を使っている場合）の設定を最新の状態に保ちたいとき
- Bazel モジュールの依存関係が複雑になり、不要な定義やバージョンの調整が必要な場合
- Bazel のビルド環境全体の再現性を確保するために、依存関係を整えたいとき

```
bazel mod tidy
```
<br>

### Go Modules の依存関係整理
- Go プロジェクト内で、Go Modules の依存関係情報を最新状態に保ちたいとき
- 例えば、コードを変更して利用するライブラリが変わった場合など、go.mod を自動整備するために実行する
- Bazel ビルド前に Go の依存関係が正しく管理されていることを確認したい場合
<br>** go mod tidyとほぼ同じことをしている<br>

```
bazel run @rules_go//go -- mod tidy
```
<br>

### Go Modules の依存情報を Bazel のリポジトリ定義へ反映
- Go の依存関係（go.mod）に変更があった場合に、Bazel 側の外部リポジトリ定義を更新したいとき
- Go Modules と Bazel のリポジトリ定義との整合性を保ちたいとき
- 不要な外部リポジトリ定義を自動で削除し、最新の依存状態を反映させたい場合
```
bazel run //:update-go-repos
```
<br>


## Build, Dockerについて

### Imageを作成する
```
 bazel run //cmd:image_amd64_load
```

### Imageを動かす
```
docker run --rm cmd:latest
```

**macで動かすため下記設定をcmd/BUILD.bazel二追加
```
go_cross_binary(
    name = "cmd_linux_amd64",
    platform = "@rules_go//go/toolchain:linux_amd64",
    target = ":cmd",
)
```
