一、静态链表

　　1.静态链表：用数组描述的链表叫做静态链表。C语言中，让数组的元素都是由两个数据域组成，data和cur。数组的每个下标都对应着一个data和一个cur。数据域data，用来存放数据元素，也就是要处理的数据；而cur相当于单链表中的next指针，存放该元素的后继在数据中的下标，把cur叫游标。另外，数组的第一个和最后一个元素作为特殊元素处理，不存数据。数组的第一个元素，即下标为0的元素的cur存放备用链表（未被使用的数组元素）第一个结点的下标，而数组的最后一个元素的cur则存放第一个有数值的元素的下标，相当于单链表中的头结点的作用。

　　2.静态链表的优缺点：

在插入和删除操作时，只需要修改游标，不需要移动元素，从而改进了在顺序存储结构中的插入和删除操作需要移动大量元素的缺点。
没有解决连续存储分配带来的表长难以确定的问题；失去了顺序存储结构随机存取的特性。


静态链表的C语言代码实现：

```c
#include "string.h"
#include "ctype.h"      

#include "stdio.h"    
#include "stdlib.h"   
#include "io.h"  
#include "math.h"  
#include "time.h"

#define OK 1
#define ERROR 0
#define TRUE 1
#define FALSE 0

#define MAXSIZE 1000 /* 存储空间初始分配量 */

typedef int Status;           /* Status是函数的类型,其值是函数结果状态代码，如OK等 */
typedef char ElemType;        /* ElemType类型根据实际情况而定，这里假设为char */


Status visit(ElemType c)
{
    printf("%c ",c);
    return OK;
}

/* 线性表的静态链表存储结构 */
typedef struct 
{
    ElemType data;
    int cur;  /* 游标(Cursor) ，为0时表示无指向 */
} Component,StaticLinkList[MAXSIZE];


/* 将一维数组space中各分量链成一个备用链表，space[0].cur为头指针，"0"表示空指针 */
Status InitList(StaticLinkList space) 
{
    int i;
    for (i=0; i<MAXSIZE-1; i++)  
        space[i].cur = i+1;
    space[MAXSIZE-1].cur = 0; /* 目前静态链表为空，最后一个元素的cur为0 */
    return OK;
}


/* 若备用空间链表非空，则返回分配的结点下标，否则返回0 */
int Malloc_SSL(StaticLinkList space) 
{ 
    int i = space[0].cur;                   /* 当前数组第一个元素的cur存的值 */
                                            /* 就是要返回的第一个备用空闲的下标 */
    if (space[0]. cur)         
        space[0]. cur = space[i].cur;       /* 由于要拿出一个分量来使用了， */
                                            /* 所以我们就得把它的下一个 */
                                            /* 分量用来做备用 */
    return i;
}


/*  将下标为k的空闲结点回收到备用链表 */
void Free_SSL(StaticLinkList space, int k) 
{  
    space[k].cur = space[0].cur;    /* 把第一个元素的cur值赋给要删除的分量cur */
    space[0].cur = k;               /* 把要删除的分量下标赋值给第一个元素的cur */
}

/* 初始条件：静态链表L已存在。操作结果：返回L中数据元素个数 */
int ListLength(StaticLinkList L)
{
    int j=0;
    int i=L[MAXSIZE-1].cur;
    while(i)
    {
        i=L[i].cur;
        j++;
    }
    return j;
}

/*  在L中第i个元素之前插入新的数据元素e   */
Status ListInsert(StaticLinkList L, int i, ElemType e)   
{  
    int j, k, l;   
    k = MAXSIZE - 1;   /* 注意k首先是最后一个元素的下标 */
    if (i < 1 || i > ListLength(L) + 1)   
        return ERROR;   
    j = Malloc_SSL(L);   /* 获得空闲分量的下标 */
    if (j)   
    {   
        L[j].data = e;   /* 将数据赋值给此分量的data */
        for(l = 1; l <= i - 1; l++)   /* 找到第i个元素之前的位置 */
           k = L[k].cur;           
        L[j].cur = L[k].cur;    /* 把第i个元素之前的cur赋值给新元素的cur */
        L[k].cur = j;           /* 把新元素的下标赋值给第i个元素之前元素的ur */
        return OK;   
    }   
    return ERROR;   
}

/*  删除在L中第i个数据元素   */
Status ListDelete(StaticLinkList L, int i)   
{ 
    int j, k;   
    if (i < 1 || i > ListLength(L))   
        return ERROR;   
    k = MAXSIZE - 1;   
    for (j = 1; j <= i - 1; j++)   
        k = L[k].cur;   
    j = L[k].cur;   
    L[k].cur = L[j].cur;   
    Free_SSL(L, j);   
    return OK;   
} 

Status ListTraverse(StaticLinkList L)
{
    int j=0;
    int i=L[MAXSIZE-1].cur;
    while(i)
    {
            visit(L[i].data);
            i=L[i].cur;
            j++;
    }
    return j;
    printf("\n");
    return OK;
}


int main()
{
    StaticLinkList L;
    Status i;
    i=InitList(L);
    printf("初始化L后：L.length=%d\n",ListLength(L));

    i=ListInsert(L,1,'F');
    i=ListInsert(L,1,'E');
    i=ListInsert(L,1,'D');
    i=ListInsert(L,1,'B');
    i=ListInsert(L,1,'A');

    printf("\n在L的表头依次插入FEDBA后：\nL.data=");
    ListTraverse(L); 

    i=ListInsert(L,3,'C');
    printf("\n在L的“B”与“D”之间插入“C”后：\nL.data=");
    ListTraverse(L); 

    i=ListDelete(L,1);
    printf("\n在L的删除“A”后：\nL.data=");
    ListTraverse(L); 

    printf("\n");

    return 0;
}


输出为：
初始化L后：L.length=0
在L的表头依次插入FEDBA后：
L.data=A B D E F 
在L的“B”与“D”之间插入“C”后：
L.data=A B C D E F 
在L的删除“A”后：
L.data=B C D E F
```

静态链表的Java语言代码实现：

```java
public class StaticLinkList {

    private Element[] L = null; 
    private int MAXSIZE = 1000;//默认存储大小
    
    class Element{
        int data;
        int cur;
    }

    // 静态链表的初始化
    public StaticLinkList(){
        L = new Element[MAXSIZE];
        for (int i = 0; i < MAXSIZE-1; i++) {
            L[i] = new Element();
            L[i].cur = i+1;
        }
        L[MAXSIZE-1] = new Element();
        L[MAXSIZE-1].cur = 0;
    }
    
    public int listLength() {
        int j = 0;
        int i = L[MAXSIZE-1].cur;
        while (i!=0) {
            i = L[i].cur;
            j++;
        }
        return j;
    }
    
    // 获得静态链表中存放备用链表的第一个结点的下标，即第一个空闲空间的数组下标
    public int mallocSLL() {
        int i = L[0].cur;
        if (L[0].cur!=0) {  // 链表为空时，空闲元素的下标即为1
            L[0].cur = L[i].cur;
        }
        return i;
    }
    
    public void listInsert(int i, int e) throws Exception{
        int k = MAXSIZE-1;
        if (i < 1 || i > listLength() + 1) 
            throw new Exception("超出范围，无法插入");
        int j = mallocSLL();
        if (j!=0) { 
            L[j].data = e;
            for (int l = 1; l <= i -1 ; l++) 
                k = L[k].cur;
            L[j].cur = L[k].cur;
            L[k].cur = j;
        }
    }
    
    public void freeSSL(int k) {
        L[k].cur = L[0].cur;
        L[0].cur = k;
    }
    
    public void listDelete(int i) throws Exception {
        if (i < 1 || i > listLength()) 
            throw new Exception("超出范围，无法删除");
        int k = MAXSIZE - 1 ;
        for (int l = 1; l <= i - 1; l++) 
            k = L[k].cur;
        int j = L[k].cur;
        L[k].cur = L[j].cur;
        freeSSL(j);
    }
    
    public void ListTraverse() {
        int i = L[MAXSIZE-1].cur;
        while (i!=0) {
            System.out.print(L[i].data + " ");
            i = L[i].cur;
        }
    }
    

    
    public static void main(String[] args) throws Exception {
        StaticLinkList sList = new StaticLinkList();
        sList.listInsert(1, 1);
        System.out.print("此时链表的输出为:");
        sList.ListTraverse();
        System.out.println("此时链表的长度为" + sList.listLength());
        
        sList.listInsert(1, 2);
        System.out.print("此时链表的输出为:");
        sList.ListTraverse();
        System.out.println("此时链表的长度为" + sList.listLength());
        
        sList.listInsert(1, 3);
        System.out.print("此时链表的输出为:");
        sList.ListTraverse();
        System.out.println("此时链表的长度为" + sList.listLength());
        
        sList.listInsert(1, 4);
        System.out.print("此时链表的输出为:");
        sList.ListTraverse();
        System.out.println("此时链表的长度为" + sList.listLength());
        
        sList.listInsert(1, 5);
        System.out.print("此时链表的输出为:");
        sList.ListTraverse();
        System.out.println("此时链表的长度为" + sList.listLength());
        
        sList.listInsert(3, 6);
        System.out.print("在第三个元素之前插入一个值6，此时链表的输出为:");
        sList.ListTraverse();
        System.out.println("此时链表的长度为" + sList.listLength());
        
        sList.listDelete(1);
        System.out.print("删除第一个元素后，此时链表的输出为:");
        sList.ListTraverse();
        System.out.println("此时链表的长度为" + sList.listLength());
        
        sList.listDelete(3);
        System.out.print("删除第三个元素后，此时链表的输出为:");
        sList.ListTraverse();
        System.out.println("此时链表的长度为" + sList.listLength());
    }
    
    
}

输出为：
此时链表的输出为:1 此时链表的长度为1
此时链表的输出为:2 1 此时链表的长度为2
此时链表的输出为:3 2 1 此时链表的长度为3
此时链表的输出为:4 3 2 1 此时链表的长度为4
此时链表的输出为:5 4 3 2 1 此时链表的长度为5
在第三个元素之前插入一个值6，此时链表的输出为:5 4 6 3 2 1 此时链表的长度为6
删除第一个元素后，此时链表的输出为:4 6 3 2 1 此时链表的长度为5
删除第三个元素后，此时链表的输出为:4 6 2 1 此时链表的长度为4
```