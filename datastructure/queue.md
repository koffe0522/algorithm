# キュー/Queue

・キューはリストの一方の端で挿入が行われ反対の端で削除が行われる 「FIFO」 の構造

> FIFO (First In, First Out) とは「最初に入ったものが最初にでてゆく」という意味

- 要素を挿入する操作は Enqueue
- 要素を削除する(取り出す)操作は Dequeue
- 先頭は Front
- 最後尾は Rear

```py
import Queue

q = Queue.Queue()

for i in range(5):
    q.put(i) # Enqueue

while not q.empty():
    print q.get() # Dequeue

'''
$ python Queue_fifo.py
0
1
2
3
4
'''
```
