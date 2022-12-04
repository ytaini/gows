# 双向链表

　　1.双向链表：在单链表的每个结点中，再设置一个指向其前驱结点的指针域。所以在双向链表中的结点都有两个指针域，一个指向直接后继，另一个指向直接前驱。

　　2.双向链表也可以有循环链表，即对于某一结点p，它的后继的前驱是它自己，它的前驱的后继还是它自己：p->next->prior = p = p->prior->next。

　　3.双向链表在插入和删除时，需要改变两个指针变量：

　　（1）插入时：假设要将存储元素e的结点s插入到结点p和p->next之间
```c
s->prior = p;        // 把p赋值给s的前驱
s->next  = p->next;  // 把p->next赋值给s的后继
p->next->prior = s;  // 把s赋值给p->next的前驱
p->next  = s;        // 把s赋值给p的后继
```
　　（2）删除时：假设要将结点p删除
```c
p->prior->next = p->next;     // 把p->next 赋值给p->prior的后继
p->next->prior = p->prior;    // 把p->prior赋值给p->next 的前驱
free(p);
```