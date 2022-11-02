# kmp算法

https://www.cnblogs.com/zzuuoo666/p/9028287.html
https://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html


- KMP算法是一种改进的字符串匹配算法
- KMP算法的时间复杂度O(m+n) 
- KMP方法算法就利用之前判断过信息，通过一个next数组(部分匹配表)，保存模式串中最长公共前后缀的长度. 每次回溯时，通过next数组找到，前面匹配过的位置，省去了大量的计算时间

next数组:
- next数组是对于模式串而言的。
- 模式串P 的 next 数组定义为：next[i] 表示 P[0:i+1] 这一个子串，next 数组各值的含义：代表当前字符之前的字符串中，有多大长度的相同前缀后缀。
  - 例如如果next [i] = k，代表j 之前的字符串中有最大长度为k 的相同前缀后缀。 特别地，k不能取i+1（因为这个子串一共才 i+1 个字符，自己肯定与自己相等，就没有意义了）。

- 快速构建next数组，是KMP算法的精髓所在，核心思想是“P自己与自己做匹配”。
 
