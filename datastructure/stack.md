# スタック/Stack

・スタックは要素の挿入と削除がリストの先頭だけで行われる 「LIFO」 のデータ構造

> LIFO (Last In, First Out) とは「最後に入ったものが最初に出てゆく」という意味

・挿入する(積む)操作は push
・要素を削除する(取り出す)操作は pop
・リストの先頭(頂上)は top
・終端(底)は bottom と呼びます。

```py
stack = ['A', 'B', 'C']
stack.push('D') # ['A', 'B', 'C', 'D']
stack.pop() # ['A', 'B', 'C']
#  top -> ['A', 'B', 'C'] <- bottom
```
