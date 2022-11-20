[TOC]

# 分治算法

字面上的解释是“分而治之”，就是把一个复杂的问题分成两个或更多的相同或相似的子问题，再把子问题分成更小的子问题……直到最后子问题可以简单的直接求解，原问题的解即子问题的解的合并。这个技巧是很多高效算法的基础，如排序算法(快速排序，归并排序)，傅立叶变换(快速傅立叶变换)……

<br>

分治策略是：对于一个规模为n的问题，若该问题可以容易地解决（比如说规模n较小）则直接解决，否则将其分解为k个规模较小的子问题，这些子问题互相独立且与原问题形式相同，递归地解这些子问题，然后将各子问题的解合并得到原问题的解。这种算法设计策略叫做分治法。

<br>

如果原问题可分割成k个子问题，1<k≤n ，且这些子问题都可解并可利用这些子问题的解求出原问题的解，那么这种分治法就是可行的。由分治法产生的子问题往往是原问题的较小模式，这就为使用递归技术提供了方便。在这种情况下，反复应用分治手段，可以使子问题与原问题类型一致而其规模却不断缩小，最终使子问题缩小到很容易直接求出其解。这自然导致递归过程的产生。分治与递归像一对孪生兄弟，经常同时应用在算法设计之中，并由此产生许多高效算法。

分治法所能解决的问题一般具有以下几个特征：
1) **该问题的规模缩小到一定的程度就可以容易地解决**
2) **该问题可以分解为若干个规模较小的相同问题，即该问题具有最优子结构性质。**
3) **利用该问题分解出的子问题的解可以合并为该问题的解；**
4) **该问题所分解出的各个子问题是相互独立的，即子问题之间不包含公共的子问题。**

上述的第一条特征是绝大多数问题都可以满足的，因为问题的计算复杂性一般是随着问题规模的增加而增加；
**第二条特征是应用分治法的前提它也是大多数问题可以满足的，此特征反映了递归思想的应用；**
**第三条特征是关键，能否利用分治法完全取决于问题是否具有第三条特征，如果具备了第一条和第二条特征，而不具备第三条特征，则可以考虑用贪心法或动态规划法。**
**第四条特征涉及到分治法的效率，如果各子问题是不独立的则分治法要做许多不必要的工作，重复地解公共的子问题，此时虽然可用分治法，但一般用动态规划法较好。**


<br>

- 分治算法可以求解的一些经典问题
    - 二分搜索
    - 大整数乘法
    - 棋盘覆盖
    - 合并排序
    - 快速排序
    - 线性时间选择
    - 最接近点对问题
    - 循环赛日程表
    - 汉诺塔
    - 分治算法的基本步骤

分治法在每一层递归上都有三个步骤：
- 分解：将原问题分解为若干个规模较小，相互独立，与原问题形式相同的子问题
- 解决：若子问题规模较小而容易被解决则直接解，否则递归地解各个子问题
- 合并：将各个子问题的解合并为原问题的解

<br>

> 根据分治法的分割原则，原问题应该分为多少个子问题才较适宜？

> 各个子问题的规模应该怎样才为适当？

答: 但人们从大量实践中发现，在用分治法设计算法时，最好使子问题的规模大致相同。换句话说，将一个问题分成大小相等的k个子问题的处理方法是行之有效的。许多问题可以取 k = 2。这种使子问题规模大致相等的做法是出自一种平衡(balancing)子问题的思想，它几乎总是比子问题规模不等的做法要好

<br>
### 递归与分治的区别

递归是一种编程技巧，一种解决问题的思维方式；分治算法很大程度上是基于递归的，解决更具体问题的算法思想。

