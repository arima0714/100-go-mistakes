# No.4: ゲッターとセッターの多用

ゲッターとセッターの自動補完的なものはgolangには存在しない。標準ライブラリにはフィールドに直接アクセスできる構造体が実装されてる場合があり、直接修正できることもある。これらはゲッター・セッターをgolangは強制していないことを表している。

ゲッターメソッドの名前は `GetXXXXX` ではなく `XXXXX` とすべきな一方で、セッターメソッドの名前は `SetXXXXX` とすべき。

基本的にゲッター・セッターは使用すべきではない。しかし、場合によってはその必要があれば使ってよい。
