# No.13: ユーティリティパッケージの作成

ユーティリティパッケージがあると `util.XXXXX()` のように無駄に `util` を書く必要が出てくるので、パッケージ名はそもそも `util` じゃない方が良い。

ユーティリティパッケージの構成を変えて、メソッド化しておくと、`<パッケージ名>.YYYYY()` のパッケージ名を省けてより分かりやすいコードに仕立てられる
