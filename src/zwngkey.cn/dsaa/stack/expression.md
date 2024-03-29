# 前缀表达式（波兰表达式）
- 前缀表达式是一种没有括号的算术表达式，与中缀表达式不同的是，其将运算符写在前面，操作数写在后面。
- 举例说明：
  - (3+4)*5-6 对应的前缀表达式就是 - * + 3 4 5 6
  - \- 1 + 2 3，它等价于1-(2+3)。




> 前缀表达式的计算机求值

从右至左扫描表达式，遇到数字时，将数字压入堆栈，遇到运算符时，弹出栈顶的两个数，用运算符对它们做相应的计算（栈顶元素 和 次顶元素），并将结果入栈；
重复上述过程直到表达式最左端，最后运算得出的值即为表达式的结果

例如 (3+4)×5-6 对应的前缀表达式就是 - × + 3 4 5 6 , 针对前缀表达式求值步骤如下:

- 从右至左扫描，将6、5、4、3压入堆栈
- 遇到+运算符，因此弹出3和4（3为栈顶元素，4为次顶元素），计算出3+4的值，得7，再将7入栈
- 接下来是×运算符，因此弹出7和5，计算出7×5=35，将35入栈
- 最后是-运算符，计算出35-6的值，即29，由此得出最终结果




> 中缀转前缀算法

1. 首先构造一个运算符栈(也可放置括号)，运算符(以括号为分界点)在栈内遵循越往栈顶优先级不降低的原则进行排列。
2. 从右至左扫描中缀表达式，从右边第一个字符开始判断：
   1. 如果当前字符是数字，则分析到数字串的结尾并将数字串直接输出。
   2. 如果是运算符，则比较优先级。如果当前运算符的优先级大于等于栈顶运算符的优先级(当栈顶是括号时，直接入栈)，则将运算符直接入栈；否则将栈顶运算符出栈并输    出，  直到当前运算符的优先级大于等于栈顶运算符的优先级(当栈顶是括号时，直接入栈)，再将当前运算符入栈。
   3. 如果是括号，则根据括号的方向进行处理。如果是右的括号，则直接入栈；否则，遇左括号前将所有的运算符全部出栈并输出，遇右括号后将左、向右的两括号一起出栈(并不 输出)。
3. 重复上述操作(2)直至扫描结束，将栈内剩余运算符全部出栈并输出，再逆缀输出字符串。中缀表达式也就转换为前缀表达式了。




# 中缀表达式

- 中缀表达式就是常见的运算表达式，如(3+4)*5-6
- 中缀表达式的求值是平常最为熟悉的，但是对计算机说却不好操作,因此，在计算结束时，往往会将中缀表达式转成其它表达式来操作（一般是转成后缀表达式）




# 后缀表达式

- 后缀表达式又称逆波兰表达式,与前缀表达式相似，只是运算符位于操作数之后
- 举例说明： (3+4)×5-6 对应的后缀表达式就是 3 4 + 5 × 6 –




> 后缀表达式的计算机求值

从左至右扫描表达式，遇到数字时，将数字压入堆栈，遇到运算符时，弹出栈顶的两个数，用运算符对它们做相应的计算（次顶元素 和 栈顶元素），并将结果入栈；重复上述过程直到表达式最右端，最后运算得出的值即为表达式的结果

例如: (3+4)×5-6 对应的后缀表达式就是 3 4 + 5 × 6 - , 针对后缀表达式求值步骤如下:

1. 从左至右扫描，将3和4压入堆栈.
2. 遇到+运算符，因此弹出4和3（4为栈顶元素，3为次顶元素），计算出3+4的值，得7，再将7入栈
3. 将5入栈
4. 接下来是×运算符，因此弹出5和7，计算出7×5=35，将35入栈
5. 将6入栈
6. 最后是-运算符，计算出35-6的值，即29，由此得出最终结果






> 中转后缀算法

1. 初始化两个栈：运算符栈s1和储存中间结果的栈s2
2. 从左至右扫描中缀表达式
3. 遇到操作数时，将其压s2
4. 遇到运算符时，比较其与s1栈顶运算符的优先级：
   1. 如果s1为空，或栈顶运算符为左括号“(”，则直接将此运算符入栈
   2. 否则，若优先级比栈顶运算符的高，也将运算符压入s1
   3. 否则，将s1栈顶的运算符弹出并压入到s2中，再次与s1中新的栈顶运算符相比
5. 遇到括号时：
   1. 如果是左括号“(”，则直接压入s1
   2. 如果是右括号“)”，则依次弹出s1栈顶的运算符，并压入s2，直到遇到左括号为止，此时将这一对括号丢弃
6. ........
7. 重复步骤2至5，直到表达式的最右边
8. 将s1中剩余的运算符依次弹出并压入s2
9. 依次弹出s2中的元素并输出，结果的逆序即为中缀表达式对应的后缀表达式


```
(a+b)*c-(a+b)/e的后缀表达式为：
(a+b)*c-(a+b)/e 
→((a+b)*c)((a+b)/e)-
→((a+b)c*)((a+b)e/)-
→(ab+c*)(ab+e/)-
→ab+c*ab+e/-
```