# No.5: インターフェース汚染

インターフェース設計時に、インターフェースがいくつメソッドを含むかを念頭に入れるべき。そして「インターフェースが大きければ大きいほど、抽象化は弱くなる。」という格言を忘れない

インターフェースを作成する３つの基本ユースケース
* 共通のふるまい
* 具体的な実装との分離
* ふるまいの制限

抽象化は作成するものではなく、発見するもの。つまり、必要な時に作成するのであって、必要だと予想されるときに作成するものではない。

インターフェースを過剰に使うと、コードの流れが複雑になりすぎてしまう。


