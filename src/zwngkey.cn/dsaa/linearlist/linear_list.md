# 线性表的顺序存储结构
一、线性表的定义：线性表就是零个或多个数据元素的有效序列。

二、线性表的顺序存储结构

1. 定义：指的是用一段地址连续地存储单元依次存储线性表的数据元素。

2. 顺序存储方式：线性表的每个数据元素的类型都相同，所以可以用一维数组来实现顺序存储结构，即把第一个数据元素存到数组下标为0的位置中，接着把线性表相邻的元素存储在数组中相邻的位置。线性表中，估算最大存储容量就是数组的长度。

3. 描述顺序存储结构需要三个属性：
    - 存储空间的起始位置：数组data，它的存储位置就是存储空间的存储位置。
    - 线性表的最大存储容量：数组长度MaxSize。
    - 线性表的当前长度：length。
4. 数组的长度是存放线性表的存储空间的长度，存储分配后这个量一般是不变的。而线性表的长度是线性表中数据元素的个数，随着线性表插入和删除操作的进行，这个量是变化的。

5. 地址计算方法：由于数据是从0开始第一个下标的，因此线性表的第i个元素是要存储在数据下标为i - 1 的位置。存储器中每个存储单元都有自己的编号，这个编号称为地址。

6. 线性表顺序存储结构的优缺点：
    - 优点：无须为表示表中元素之间的逻辑关系而增加额外的存储空间；可以快速地存取表中任意位置的元素。
    - 缺点：插入和删除操作需要移动大量元素；当线性表长度变化较大时，难以确定存储空间的容量；造成存储空间的“碎片”。（实际存储空间小于分配的空间）

C代码实现:
```c
#include "stdio.h"    
#include "stdlib.h"   
#include "io.h"  
#include "math.h"  
#include "time.h"

#define OK 1
#define ERROR 0
#define TRUE 1
#define FALSE 0

#define MAXSIZE 20 /* 存储空间初始分配量 */

typedef int Status;          /* Status是函数的类型,其值是函数结果状态代码，如OK等 */
typedef int ElemType;        /* ElemType类型根据实际情况而定，这里假设为int */

/* 将数据元素c打印出来 */
Status visit(ElemType c)
{
    printf("%d ",c);
    return OK;
}

/* 线性表的顺序存储的结构 */
typedef struct
{
    ElemType data[MAXSIZE];        /* 数组，存储数据元素 */
    int length;                    /* 线性表当前长度 */
}SqList;                           /* 给结构体命名 */20

/* 初始化顺序线性表，建立一个空的线性表L */
Status InitList(SqList *L) 
{ 
    L->length=0;
    return OK;
}

/* 初始条件：顺序线性表L已存在。操作结果：若L为空表，则返回TRUE，否则返回FALSE */
Status ListEmpty(SqList L)
{ 
    if(L.length==0)
        return TRUE;
    else
        return FALSE;
}

/* 初始条件：顺序线性表L已存在。操作结果：将L重置为空表 */
Status ClearList(SqList *L)
{ 
    L->length=0;
    return OK;
}

/* 初始条件：顺序线性表L已存在。操作结果：返回L中数据元素个数 */
int ListLength(SqList L)
{
    return L.length;
}

/* 获得元素操作：将线性表L中的第i个位置元素值返回，*/
/* 初始条件：顺序线性表L已存在，1≤i≤ListLength(L) */
/* 操作结果：用e返回L中第i个数据元素的值,注意i是指位置，第1个位置的数组是从0开始 */
Status GetElem(SqList L,int i,ElemType *e)
{
    if(L.length==0 || i<1 || i>L.length)
            return ERROR;
    *e=L.data[i-1];  /* 将线性表的第i个位置，也就是数组的第i-1个位置的元素返回 */

    return OK;
}

/* 初始条件：顺序线性表L已存在 */
/* 操作结果：返回L中第1个与e满足关系的数据元素的位序。 */
/* 若这样的数据元素不存在，则返回值为0 */
int LocateElem(SqList L,ElemType e)
{
    int i;
    if (L.length==0)
            return 0;
    for(i=0;i<L.length;i++)
    {
        if (L.data[i]==e)
            break;
    }
    if(i>=L.length)
            return 0;

    return i+1;
}


/* 初始条件：顺序线性表L已存在,1≤i≤ListLength(L)， */
/* 操作结果：在L中第i个位置之前插入新的数据元素e，L的长度加1 */
Status ListInsert(SqList *L,int i,ElemType e)
{ 
    int k;
    if (L->length==MAXSIZE)  /* 顺序线性表已经满 */
        return ERROR;
    if (i<1 || i>L->length+1)/* 当i比第一位置小或者比最后一位置后一位置还要大时 */
        return ERROR;

    if (i<=L->length)        /* 若插入数据位置不在表尾 */
    {
        for(k=L->length-1;k>=i-1;k--)  /* 将要插入位置之后的数据元素向后移动一位 */
            L->data[k+1]=L->data[k];
    }
    L->data[i-1]=e;          /* 将新元素插入 */
    L->length++;             /* 线性表L的长度增加1 */

    return OK;
}

/* 初始条件：顺序线性表L已存在，1≤i≤ListLength(L) */
/* 操作结果：删除L的第i个数据元素，并用e返回其值，L的长度减1 */
Status ListDelete(SqList *L,int i,ElemType *e) 
{ 
    int k;
    if (L->length==0)               /* 线性表为空 */
        return ERROR;
    if (i<1 || i>L->length)         /* 删除位置不正确 */
        return ERROR;
    *e=L->data[i-1];                /* 将被删除的数据元素返回给e */
    if (i<L->length)                /* 如果删除不是最后位置 */
    {
        for(k=i;k<L->length;k++)    /* 将删除位置后继元素前移1个位置 */
            L->data[k-1]=L->data[k];
    }
    L->length--;                    /* 线性表L的长度减1 */
    return OK;
}

/* 初始条件：顺序线性表L已存在 */
/* 操作结果：依次对L的每个数据元素输出 */
Status ListTraverse(SqList L)
{
    int i;
    for(i=0;i<L.length;i++)
        visit(L.data[i]);
    printf("\n");
    return OK;
}

/* 将所有的在线性表Lb中但不在La中的数据元素插入到La中 */
void unionL(SqList *La,SqList Lb)
{
    int La_len,Lb_len,i;    /* 声明线性表La和Lb的长度 */
    ElemType e;             /* 声明线性表La和Lb相同的数据元素e */
    La_len=ListLength(*La); /* 求线性表La的长度 */
    Lb_len=ListLength(Lb);  /* 求线性表Lb的长度 */
    for (i=1;i<=Lb_len;i++)
    {
        GetElem(Lb,i,&e);   /* 去Lb中第i个数据元素赋给e */
        if (!LocateElem(*La,e))  /* 如果线性表La中不存在和e相同的数据元素 */
            ListInsert(La,++La_len,e); /* 就把线性表中La中不存在的这个元素插入线性表La中 */
    }
}



int main()
{
        
    SqList L;
    SqList Lb;
    
    ElemType e;
    Status i;
    int j,k;

    i=InitList(&L);
    printf("初始化L后：L.length=%d\n",L.length);
    printf("\n");

    for(j=1;j<=5;j++)
            i=ListInsert(&L,1,j);
    printf("在L的表头依次插入1～5后：L.data=");
    ListTraverse(L); 
    printf("L.length=%d \n",L.length);
    printf("\n");

    i=ListEmpty(L);
    printf("L是否空：i=%d(1:是 0:否)\n",i);
    i=ClearList(&L);
    printf("清空L后：L.length=%d\n",L.length);
    i=ListEmpty(L);
    printf("L是否空：i=%d(1:是 0:否)\n",i);
    printf("\n");

    for(j=1;j<=10;j++)
            ListInsert(&L,j,j);
    printf("在L的表尾依次插入1～10后：L.data=");
    ListTraverse(L); 
    printf("L.length=%d \n",L.length);
    printf("\n");


    ListInsert(&L,1,0);
    printf("在L的表头插入0后：L.data=");
    ListTraverse(L); 
    printf("L.length=%d \n",L.length);
    printf("\n");

    GetElem(L,5,&e);
    printf("第5个元素的值为：%d\n",e);
    for(j=3;j<=4;j++)
    {
            k=LocateElem(L,j);
            if(k)
                    printf("第%d个元素的值为：%d\n",k,j);
            else
                    printf("没有值为%d的元素：\n",j);
    }
    printf("\n");
    
    printf("依次输出L的元素：");
    ListTraverse(L);
    k=ListLength(L); /* k为表长 */
    for(j=k+1;j>=k;j--)
    {
            i=ListDelete(&L,j,&e); /* 删除第j个数据 */
            if(i==ERROR)
                    printf("删除第%d个数据失败\n",j);
            else
                    printf("删除第%d个的元素值为：%d\n",j,e);
    }
    printf("依次输出L的元素：");
    ListTraverse(L); 
    printf("\n");

    j=5;
    ListDelete(&L,j,&e); /* 删除第5个数据 */
    printf("删除第%d个的元素值为：%d\n",j,e);

    printf("依次输出L的元素：");
    ListTraverse(L); 
    printf("\n");

    //构造一个有10个数的Lb
    i=InitList(&Lb);
    for(j=6;j<=15;j++)
        i=ListInsert(&Lb,1,j);

    printf("依次输出Lb的元素：");
    ListTraverse(Lb);

    unionL(&L,Lb);

    printf("依次输出合并了Lb的L的元素：");
    ListTraverse(L); 

    return 0;
}

输出：
初始化L后：L.length=0

在L的表头依次插入1～5后：L.data=5 4 3 2 1 
L.length=5 

L是否空：i=0(1:是 0:否)
清空L后：L.length=0
L是否空：i=1(1:是 0:否)

在L的表尾依次插入1～10后：L.data=1 2 3 4 5 6 7 8 9 10 
L.length=10 

在L的表头插入0后：L.data=0 1 2 3 4 5 6 7 8 9 10 
L.length=11 

第5个元素的值为：4
第4个元素的值为：3
第5个元素的值为：4

依次输出L的元素：0 1 2 3 4 5 6 7 8 9 10 
删除第12个数据失败
删除第11个的元素值为：10
依次输出L的元素：0 1 2 3 4 5 6 7 8 9 

删除第5个的元素值为：4
依次输出L的元素：0 1 2 3 5 6 7 8 9 

依次输出Lb的元素：15 14 13 12 11 10 9 8 7 6 
依次输出合并了Lb的L的元素：0 1 2 3 5 6 7 8 9 15 14 13 12 11 10 
```


# 线性表的链式存储结构（单链表）

一、线性表的顺序存储结构的不足：线性表的顺序结构最大的缺点就是插入和删除时需要移动大量元素，这显然就需要耗费时间。原因就在于相邻两元素的存储位置也具有邻居关系。它们编号是1,2,3...n，它们在内存中的位置也是挨着的吗，中间没有空隙，当然就无法快速插入，而删除后，当中就会留出空隙，自然需要弥补。

二、线性表的链式存储结构的特点是用一组任意的存储单元存储线性表的数据元素，这组存储单元可以是连续的，也可以是不连续的。这就意味着，这些数据元素可以存在内存未被占用的任意位置。在链式结构中，除了要存数据元素以外，还要存储后继元素的存储地址。

三、线性表的链式存储结构的组成：把存储数据元素信息的域称为数据域，把存储直接后继位置的域称为指针域。指针域中存储的信息称作指针或链。这两部分信息组成数据元素的存储映像，称为结点node。n个结点链结成一个链表，即为线性表的链式存储结构，因为此链表的每个结点中只包含一个指针域，所以叫做单链表。

四、链表中第一个结点的存储位置叫做头指针，之后每一个结点，其实就是上一个的后继指针指向的位置，最后一个结点为“空”。为了更加方便地对链表进行操作，会在单链表的第一个结点前附设一个结点，称为头结点。头结点的数据域可以不存储任何信息，也可以存储如线性表的长度等附加信息，头结点的指针域存储指向第一个结点的指针。

五、单链表结构与顺序存储结构优缺点：
- 存储分配方式：顺序存储结构用一段连续的存储单元依次存储线性表的数据元素；单链表采用链式存储结构，用一组任意的存储单元存放线性表的元素。
- 时间性能：查找时，顺序存储结构O(1)和单链表O(n)；插入和删除时，顺序存储结构需要平均移动表长一半的元素O(n)而单链表直接找出某位置的指针后，插入和删除的时间仅为O()

c语言代码实现
```c
#include "stdio.h"    
#include "string.h"
#include "ctype.h"      
#include "stdlib.h"   
#include "io.h"  
#include "math.h"  
#include "time.h"

#define OK 1
#define ERROR 0
#define TRUE 1
#define FALSE 0


typedef int Status;  /* Status是函数的类型,其值是函数结果状态代码，如OK等 */
typedef int ElemType;/* ElemType类型根据实际情况而定，这里假设为int */


Status visit(ElemType c)
{
    printf("%d ",c);
    return OK;
}

/* 线性表的单链表存储结构*/
typedef struct Node
{
    ElemType data;
    struct Node *next;
}Node;
typedef struct Node *LinkList; /* 定义LinkList */

/* 初始化顺序线性表 */
Status InitList(LinkList *L) 
{ 
    *L=(LinkList)malloc(sizeof(Node)); /* 产生头结点,并使L指向此头结点 */
    if(!(*L)) /* 存储分配失败 */
            return ERROR;
    (*L)->next=NULL; /* 指针域为空 */

    return OK;
}

/* 初始条件：顺序线性表L已存在。操作结果：若L为空表，则返回TRUE，否则返回FALSE */
Status ListEmpty(LinkList L)
{ 
    if(L->next)
            return FALSE;
    else
            return TRUE;
}

/* 初始条件：顺序线性表L已存在。操作结果：将L重置为空表 */
/*单链表整表删除的算法思路：
 * 1.声明一结点p和q
 * 2.将第一个结点赋值给p
 * 3.循环：
 *         将下一结点赋值给q
 *         释放p
 *         将q赋值给p*/
Status ClearList(LinkList *L)
{ 
    LinkList p,q;
    p=(*L)->next;           /*  p指向第一个结点 */
    while(p)                /*  没到表尾 */
    {
        q=p->next;
        free(p);
        p=q;
    }
    (*L)->next=NULL;        /* 头结点指针域为空 */
    return OK;
}

/* 初始条件：顺序线性表L已存在。操作结果：返回L中数据元素个数 */
int ListLength(LinkList L)
{
    int i=0;
    LinkList p=L->next; /* p指向第一个结点 */
    while(p)                        
    {
        i++;
        p=p->next;
    }
    return i;
}

/* 初始条件：顺序线性表L已存在，1≤i≤ListLength(L) */
/* 操作结果：用e返回L中第i个数据元素的值 */
/* 1.声明一个指针p指向链表第一个结点，初始化j从1开始
 * 2.当j<i时，就遍历链表，让p的指针向后移动，不断指向下一结点，j累加1
 * 3.若到链表末尾p为空，则说明第i个结点不存在
 * 4.否则查找成功，返回结点p的数据。*/
Status GetElem(LinkList L,int i,ElemType *e)
{
    int j;
    LinkList p;             /* 声明一结点p */
    p = L->next;         /* 让p指向链表L的第一个结点 */
    j = 1;                 /* j为计数器 */
    while (p && j<i)     /* p不为空或者计数器j还没有等于i时，循环继续 */
    {   
        p = p->next;     /* 让p指向下一个结点 */
        ++j;
    }
    if ( !p || j>i ) 
        return ERROR;    /*  第i个元素不存在 */
    *e = p->data;        /*  取第i个元素的数据 */
    return OK;
}

/* 初始条件：顺序线性表L已存在 */
/* 操作结果：返回L中第1个与e满足关系的数据元素的位序。 */
/* 若这样的数据元素不存在，则返回值为0 */
int LocateElem(LinkList L,ElemType e)
{
    int i=0;
    LinkList p=L->next;
    while(p)
    {
        i++;
        if(p->data==e) /* 找到这样的数据元素 */
                return i;
        p=p->next;
    }

    return 0;
}


/* 初始条件：顺序线性表L已存在,1≤i≤ListLength(L)， */
/* 操作结果：在L中第i个位置之前插入新的数据元素e，L的长度加1 */
/* 1.声明一指针p指向链表头结点，初始化j从1开始
 * 2.当j<i时，就遍历链表，让p的指针向后移动， 不断指向下一结点，j累加1
 * 3.若到链表末尾p为空，则说明第i个结点不存在
 * 4.否则查找成功，在系统中生成一个空节点s
 * 5.将数据元素e赋值给s->data
 * 6.单链表的插入标准语句s->next=p->next; p->next=s
 * 7.返回成功*/
Status ListInsert(LinkList *L,int i,ElemType e)
{ 
    int j;
    LinkList p,s;
    p = *L;   
    j = 1;
    while (p && j < i)     /* 寻找第i-1个结点 */
    {
        p = p->next;
        ++j;
    } 
    if (!p || j > i) 
        return ERROR;     /* 第i个元素不存在 */
    s = (LinkList)malloc(sizeof(Node));  /*  生成新结点(C语言标准函数) */
    s->data = e;  
    s->next = p->next;    /* 将p的后继结点赋值给s的后继  */
    p->next = s;          /* 将s赋值给p的后继 */
    return OK;
}

/* 初始条件：顺序线性表L已存在，1≤i≤ListLength(L) */
/* 操作结果：删除L的第i个数据元素，并用e返回其值，L的长度减1 */
/* 1.声明一指针p指向链表头结点，初始化j从1开始
 * 2.当j<i时，就遍历链表，让p的指针向后移动， 不断指向下一结点，j累加1
 * 3.若到链表末尾p为空，则说明第i个结点不存在
 * 4.否则查找成功，将想要删除的结点p->next赋值给q
 * 5.单链表标准删除语句：p->next=q->next
 * 6.将q结点中的数据赋值给e，作为返回
 * 7.释放q结点
 * 8.返回成功
 * */
Status ListDelete(LinkList *L,int i,ElemType *e) 
{ 
    int j;
    LinkList p,q;
    p = *L;
    j = 1;
    while (p->next && j < i)    /* 遍历寻找第i个元素 */
    {
        p = p->next;
        ++j;
    }
    if (!(p->next) || j > i) 
        return ERROR;           /* 第i个元素不存在 */
    q = p->next;
    p->next = q->next;            /* 将q的后继赋值给p的后继 */
    *e = q->data;               /* 将q结点中的数据给e */
    free(q);                    /* 让系统回收此结点，释放内存 */
    return OK;
}

/* 初始条件：顺序线性表L已存在 */
/* 操作结果：依次对L的每个数据元素输出 */
Status ListTraverse(LinkList L)
{
    LinkList p=L->next;
    while(p)
    {
        visit(p->data);
        p=p->next;
    }
    printf("\n");
    return OK;
}

/*单链表整表创建的算法思路(包括头插法和尾插法)：
 * 1.声明一指针p和计数器变量i
 * 2.初始化一空链表L
 * 3.让L的头结点的指针指向NULL，即建立一个带头结点的单链表
 * 4.循环：
 *          生成一新节点赋值给p
 *          随机生成一数字赋值给p的数据域p->data
 *          将p插入到头结点与一新节点之间*/


/*  随机产生n个元素的值，建立带表头结点的单链线性表L（头插法） */
void CreateListHead(LinkList *L, int n) 
{
    LinkList p;
    int i;
    srand(time(0));                         /* 初始化随机数种子 */
    *L = (LinkList)malloc(sizeof(Node));
    (*L)->next = NULL;                      /*  先建立一个带头结点的单链表 */
    for (i=0; i<n; i++) 
    {
        p = (LinkList)malloc(sizeof(Node)); /*  生成新结点 */
        p->data = rand()%100+1;             /*  随机生成100以内的数字 */
        p->next = (*L)->next;    
        (*L)->next = p;                        /*  插入到表头 */
    }
}

/*  随机产生n个元素的值，建立带表头结点的单链线性表L（尾插法） */
void CreateListTail(LinkList *L, int n) 
{
    LinkList p,r;
    int i;
    srand(time(0));                      /* 初始化随机数种子 */
    *L = (LinkList)malloc(sizeof(Node)); /* L为整个线性表 */
    r=*L;                                /* r为指向尾部的结点 */
    for (i=0; i<n; i++) 
    {
        p = (Node *)malloc(sizeof(Node)); /*  生成新结点 */
        p->data = rand()%100+1;           /*  随机生成100以内的数字 */
        r->next=p;                        /* 将表尾终端结点的指针指向新结点 */
        r = p;                            /* 将当前的新结点定义为表尾终端结点 */
    }
    r->next = NULL;                       /* 表示当前链表结束 */
}

int main()
{        
    LinkList L;
    ElemType e;
    Status i;
    int j,k;
    i=InitList(&L);
    printf("初始化L后：ListLength(L)=%d\n",ListLength(L));
    for(j=1;j<=5;j++)
            i=ListInsert(&L,1,j);
    printf("在L的表头依次插入1～5后：L.data=");
    ListTraverse(L); 

    printf("ListLength(L)=%d \n",ListLength(L));
    i=ListEmpty(L);
    printf("L是否空：i=%d(1:是 0:否)\n",i);

    i=ClearList(&L);
    printf("清空L后：ListLength(L)=%d\n",ListLength(L));
    i=ListEmpty(L);
    printf("L是否空：i=%d(1:是 0:否)\n",i);

    for(j=1;j<=10;j++)
            ListInsert(&L,j,j);
    printf("在L的表尾依次插入1～10后：L.data=");
    ListTraverse(L); 

    printf("ListLength(L)=%d \n",ListLength(L));

    ListInsert(&L,1,0);
    printf("在L的表头插入0后：L.data=");
    ListTraverse(L); 
    printf("ListLength(L)=%d \n",ListLength(L));

    GetElem(L,5,&e);
    printf("第5个元素的值为：%d\n",e);
    for(j=3;j<=4;j++)
    {
            k=LocateElem(L,j);
            if(k)
                    printf("第%d个元素的值为%d\n",k,j);
            else
                    printf("没有值为%d的元素\n",j);
    }
    

    k=ListLength(L); /* k为表长 */
    for(j=k+1;j>=k;j--)
    {
            i=ListDelete(&L,j,&e); /* 删除第j个数据 */
            if(i==ERROR)
                    printf("删除第%d个数据失败\n",j);
            else
                    printf("删除第%d个的元素值为：%d\n",j,e);
    }
    printf("依次输出L的元素：");
    ListTraverse(L); 

    j=5;
    ListDelete(&L,j,&e); /* 删除第5个数据 */
    printf("删除第%d个的元素值为：%d\n",j,e);

    printf("依次输出L的元素：");
    ListTraverse(L); 

    i=ClearList(&L);
    printf("\n清空L后：ListLength(L)=%d\n",ListLength(L));
    CreateListHead(&L,20);
    printf("整体创建L的元素(头插法)：");
    ListTraverse(L); 
    
    i=ClearList(&L);
    printf("\n删除L后：ListLength(L)=%d\n",ListLength(L));
    CreateListTail(&L,20);
    printf("整体创建L的元素(尾插法)：");
    ListTraverse(L); 


    return 0;
}



输出为：
初始化L后：ListLength(L)=0
在L的表头依次插入1～5后：L.data=5 4 3 2 1 
ListLength(L)=5 
L是否空：i=0(1:是 0:否)
清空L后：ListLength(L)=0
L是否空：i=1(1:是 0:否)
在L的表尾依次插入1～10后：L.data=1 2 3 4 5 6 7 8 9 10 
ListLength(L)=10 
在L的表头插入0后：L.data=0 1 2 3 4 5 6 7 8 9 10 
ListLength(L)=11 
第5个元素的值为：4
第4个元素的值为3
第5个元素的值为4
删除第12个数据失败
删除第11个的元素值为：10
依次输出L的元素：0 1 2 3 4 5 6 7 8 9 
删除第5个的元素值为：4
依次输出L的元素：0 1 2 3 5 6 7 8 9 

清空L后：ListLength(L)=0
整体创建L的元素(头插法)：78 68 67 81 41 35 54 35 96 91 45 33 20 66 6 55 27 1 29 46 

删除L后：ListLength(L)=0
整体创建L的元素(尾插法)：46 29 1 27 55 6 66 20 33 45 91 96 35 54 35 41 81 67 68 78 
```