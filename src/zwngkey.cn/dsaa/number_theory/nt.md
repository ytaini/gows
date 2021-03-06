# 复数

- 我们把形如 z=a+bi（a、b均为实数）的数称为复数。其中，a 称为[实部](https://baike.baidu.com/item/实部/53626919)，b 称为虚部，i 称为[虚数](https://baike.baidu.com/item/虚数)单位,且![img](https://bkimg.cdn.bcebos.com/formula/ea3de96c5bc7e213f5cd4e32a0d752ac.svg)
- 当 z 的虚部 b＝0 时，则 z 为实数；当 z 的[虚部](https://baike.baidu.com/item/虚部/5231815) b≠0 时，实部 a＝0 时，常称 z 为[纯虚数](https://baike.baidu.com/item/纯虚数/3386848)。

- 复数域是实数域的代数闭包，即任何复系数多项式在复数域中总有根。

- 复数的集合用C表示，实数的集合用R表示，显然，R是C的真子集。 

- 复数集是无序集，不能建立大小顺序。

- 实际意义:

  - 我们可以在平面直角坐标系中画出虚数系统。如果利用横轴表示全体实数，那么纵轴即可表示虚数。整个平面上每一点对应着一个复数，称为复平面。横轴和纵轴也改称为实轴和虚轴。在此时，一点P坐标为P （a，bi），将坐标乘上i即点绕圆心逆时针旋转90度。

- i的性质

  i 的高次方会不断作以下的循环：

  i<sup>1</sup>= i，

  i<sup>2</sup>= - 1，

  i<sup>3</sup> = - i，

  i<sup>4</sup> = 1，

  i<sup>5</sup> = i，

  i<sup>6</sup> = - 1.



# 实数

- 实数的集合用符号**R**表示

# 虚数

- 在数学中，虚数就是形如a+b*i的数，其中a,b是实数，且b≠0,i² = - 1
- 后来发现虚数a+b\*i的实部a可对应平面上的横轴，虚部b与对应平面上的纵轴，这样虚数a+b*i可与平面内的点(a,b)对应。
- 可以将虚数bi添加到实数a以形成形式a + bi的复数，其中实数a和b分别被称为复数的实部和虚部。
- 一些作者使用术语纯虚数来表示所谓的虚数，**虚数表示具有非零虚部的任何复数**

# 有理数

- 有理数是[整数](https://baike.baidu.com/item/整数/1293937)（正整数、0、负整数）和分数的统称，是整数和分数的集合。
- 有理数集可以用大写黑正体符号**Q**代表。但**Q**并不表示有理数，有理数集与有理数是两个不同的概念。有理数集是元素为全体有理数的**集合**，而有理数则为有理数集中的所有**元素**。



# 无理数

无理数，也称为无限不循环小数，不能写作两整数之比。若将它写成小数形式，小数点之后的数字有无限多个，并且不会循环。 常见的无理数有非完全平方数的平方根、π和e（其中后两者均为超越数）等。无理数的另一特征是无限的连分数表达式。

# 整数

- 整数（integer）是正整数、零、负整数的集合。

- 在整数系中，零和正整数统称为自然数

- 整数不包括小数、分数。

- 整数也可分为[奇数](https://baike.baidu.com/item/奇数)和偶数两类。

- 整数中，能够被2整除的数，叫做偶数。不能被2整除的数则叫做奇数

- 整数的集合用 **Z**来表示

  # 负整数

  负整数是在自然数前面加上负号(-)所得的数。例如，-1、-2、-3、-38……都是负整数，负整数是小于0的整数，用Z<sup>-</sup>表示。 

  # 正整数

  - 除了0以外的自然数就是正整数。正整数又可分为质数，1和合数
  - N*/ N<sup>+</sup>：正整数集合

  

# 自然数

- 自然数是指用以计量事物的件数或表示事物次序的数。

- 自然数由0开始，一个接一个，组成一个无穷的集体。

- 自然数有有序性，无限性。

- 分为偶数和奇数，合数和质数等。

- 常用 **N** 来表示

  

# 最大公约数

- 如果数a能被数b整除，a就叫做b的**倍数**，b就叫做a的**约数**

- 对于任意整数i,j,k,它们的最大公约数记为**gcd(i,j,k)**
- 两个或多个整数共有约数中最大的一个

```
gcd(a,b)=gcd(b,a)
gcd(ma,mb)=m * gcd(a,b)
gcd(a,b)=gcd(b,a%b)
gcd(a,0)=a
gcd(0,b)=b
```





# 互质数

- 如果两个自然数是互质数，那么它们的最大公约数是1，最小公倍数是这两个数的乘积

- 两个整数分别除以它们的最大公约数，所得的商是互质数。







# 数根

- 在数学中，数根(又称位数根或数字根Digital root)是自然数的一种性质，换句话说，每个自然数都有一个数根。
  - 将一正整数的各个位数相加(即横向相加)后，若加完后的值大于等于10的话，则继续将各位数进行横向相加直到其值小于十为止所得到的数，即为数根。换句话说，数根是将一数字重复做其数字之和，直到其值小于十为止，则所得的值为该数的**数根**。
    - 例如54817的数根为7，因为5+4+8+1+7=25，25大于10则再加一次，2+5=7，7小于十，则7为54817的数根。



- *树根用途:*

  - *数根可以计算模运算的同余，对于非常大的数字的情况下可以节省很多时间。*

  - *数字根可作为一种检验计算正确性的方法。例如，两数字的和的数根等于两数字分别的数根的和。*

  -  *另外，数根也可以用来判断数字的整除性，如果数根能被3或9整除，则原来的数也能被3或9整除。*
    

- **公式法求数根：**

  **a的数根: b = (a-1) % 9+1, 即 mod(a-1, 9)+1，且a ∈ N***

  - 树根的规律:

    把 1 到 30 的树根列出来。

    原数:0 1 2 3 4 5 6 7 8 **9** 10 11 12 13 14 15 16 17 **18** 19 20 21 22 23 24 25 26 **27** 28 29 30
    数根:0 1 2 3 4 5 6 7 8 **9**  1  2   3   4   5  6   7  8  **9**   1   2  3   4  5   6   7   8   **9**   1   2   3 

    **可以发现除0以外,数根 9 个为一组， 1 - 9 循环出现。**

    结合上边的规律，对于给定的 n 有三种情况。

    - n 是 0 ，数根就是 0。

    - n 不是 9 的倍数，数根就是 n 对 9 取余，即 n mod 9。

    - n 是 9 的倍数，数根就是 9。

    **将所有情况统一起来**,我们将给定的数字减 `1`，相当于原数整体向左偏移了 `1`，然后再将得到的数字对 `9` 取余，最后将得到的结果加 `1` 即可。

    原数是 n，树根就可以表示成 **(n-1) mod 9 + 1**，可以结合下边的过程理解。

    ```
    原数: 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30
    偏移: 0 1 2 3 4 5 6 7 8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 
    取余: 0 1 2 3 4 5 6 7 8  0  1  2  3  4  5  6  7  8  0  1  2  3  4  5  6  7  8  0  1  2  
    数根: 1 2 3 4 5 6 7 8 9  1  2  3  4  5  6  7  8  9  1  2  3  4  5  6  7  8  9  1  2  3 
    ```

    

- 数学证明:

  对于一个n位的10进制数x, x可以表示成x = a<sub>0</sub>•10<sup>0</sup>+a<sub>1</sub>•10<sup>1</sup>+...+a<sub>n-1</sub>•10<sup>n-1</sup>，其中a<sub>i</sub>表示从低到高的每一位.

  即![[公式]](https://www.zhihu.com/equation?tex=x+%3D+%5Csum_%7Bi%3D0%7D%5E%7Bn-1%7D%7Ba_i%7D%7B10%5Ei%7D)

  因为:  ![[公式]](https://www.zhihu.com/equation?tex=10%5En+%5Cequiv+1%5En+%5Cequiv+1+%5Cmod+9)

  那么 ![[公式]](https://www.zhihu.com/equation?tex=x+%5Cequiv+%5Csum_%7Bi%3D0%7D%5E%7Bn-1%7Da_i+%5Cmod+9), 也就是一个数和它的各数位之和的模9相同。

  不如我们把这个操作记为f即![[公式]](https://www.zhihu.com/equation?tex=f%28x%29+%3D++%5Csum_%7Bi%3D0%7D%5E%7Bn-1%7Da_i+)

  也就是![[公式]](https://www.zhihu.com/equation?tex=f%28x%29+%5Cequiv+x+%5Cmod+9)

  所以![[公式]](https://www.zhihu.com/equation?tex=f%28f%28x%29%29+%5Cequiv+f%28x%29+%5Cequiv+x+%5Cmod+9)

  也就是说每做一次这样的操作，它对于9的模始终是不变的

  所以最终求出的数根和原数对9的模相同。







# 同余定理

- 给定一个正整数m，如果两个整数a和b满足a-b能够被m整除，即(a-b)/m得到一个整数，那么就称整数a与b对模m同余，记作a≡b(mod m)。对模m同余是整数的一个等价关系

  

- 两个整数a、b，若它们除以整数m所得的余数相等，则称a与b对于模m同余或a同余于b模m。

  记作：a≡b (mod m)，

  读作：a同余于b模m，或读作a与b对模m同余，例如26≡2(mod 12)。

