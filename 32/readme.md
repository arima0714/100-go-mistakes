# No.32: range ループでポインタ要素を使う影響を無視する

range ループを使ってデータ構造を反復処理する場合、すべての値は一意のアドレスを持つ一意の変数に代入される。

反復処理ごとにこの変数を参照するポインタを保存すると、最新の要素のコピーを保持する一意な変数のポインタを保存することになる。

解決策

* ループのスコープに強制的にローカル変数を作成する
* インデックスを介してスライス要素を参照するポインタを作成する
