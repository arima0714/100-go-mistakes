# No.35: ループ内で defer を使う

**取り囲んでいる関数が戻ったときに関数の呼び出しが実行される**

解決方法：
* 反復処理の間に呼び出される別の関数の導入
* defer を使わず、手動で defer でしたかったことを実行する