# リスト(配列)

・リストは要素が順番に並んだデータ構造
・要素は順番に並べられ、順序付けには連続した範囲の整数が使用されます。個々の要素へアクセスするにはインデックス(順序番号)を指定します。

```py
list = [4, 2, 0, 7, 5, 9, 1, 8]
list[0] # 4
list[1] # 2
list[5] # 9
```

## リストと配列の違い

|                    | リスト    | 配列         |
| ------------------ | --------- | ------------ |
| 要素の数           | 可変      | 不変         |
| 格納できるデータ型 | 何でも OK | 決められた型 |

・配列のデータに配列を用いれば「多次元配列」
・リストにリストを格納すれば「多次元リスト」
