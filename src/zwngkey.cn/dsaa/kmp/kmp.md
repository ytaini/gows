# kmp算法

https://www.cnblogs.com/zzuuoo666/p/9028287.html
https://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html


- KMP算法是一种改进的字符串匹配算法
- KMP算法的时间复杂度O(m+n) 
- KMP方法算法就利用之前判断过信息，通过一个next数组(或部分匹配表)，保存模式串中最长公共前后缀的长度. 每次回溯时，通过next数组找到，前面匹配过的位置，省去了大量的计算时间

