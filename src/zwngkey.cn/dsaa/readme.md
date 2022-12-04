# 学习基于go 的数据结构与算法.

一、数据结构基础

1. 数据结构是一门研究非数值计算的程序设计问题中的操作对象，以及它们之间的关系和操作等相关问题的学科。

2. 程序设计 = 数据结构 + 算法

3. 数据：是描述客观事物的符号，是计算机中可以操作的对象，是能被计算机识别，并输入给计算机处理的符号集合。

4. 数据元素：是组成数据的、有一定意义的基本单位，在计算机中通常作为整体处理。也被称为记录。

5. 数据项：一个数据元素可以由若干个数据项组成。数据项是数据不可分割的最小单位。

6. 数据对象：是性质相同的数据元素的集合，是数据的自己。

7. 数据结构：是相互之间存在一种或多种特定关系的数据元素的集合。

二、结构

1. 逻辑结构：是指数据对象中数据元素之间的相互关系。

（1）集合结构：集合结构中的数据元素除了同属于一个集合外，它们之间没有其他关系。

（2）线性结构：线性结构中的数据元素之间是一对一的关系。

（3）树形结构：树形结构中的数据元素之间存在一种一对多的层次关系。

（4）图形结构：图形结构的数据元素是多对多的关系。

2. 物理结构：是指数据的逻辑结构在计算机中的存储形式。

（1）顺序存储结构：是把数据元素存放在地址连续地存储单元里，其数据间的逻辑关系和物理关系是一致的。

（2）链式存储结构：是把数据元素存放在任意的存储单元里，这组存储单元可以是连续的，也可以是不连续的。

三、抽象数据类型

1. 数据类型：是指一组性质相同的值的集合及定义在此集合上的一些操作的总称。

2. 抽象：是指抽取出事务具有的普遍性的本质。

3. 抽象数据类型（Abstract Data Type， ADT）：是指一个数据模型及定义在该模型的一组操作。“抽象”的意义在于数据类型的数学抽象特性。事实上，抽象数据类型体现了程序设计中问题分解、抽象和信息隐藏的特性。



## 算法基础与复杂度

一、算法基础

　　1.算法是解决特定问题求解步骤的描述，在计算机中表现为指令的有限序列，并且每条指令表示一个或多个操作。

　　2.算法具有五个基本特性：输入、输出、有穷性、确定性和可行性。

　　（1）输入输出：算法具有零个或多个输入，但是至少有一个或多个输出。　

　　（2）有穷性：指算法在执行有限的步骤之后，自动结束而不会出现无限循环，并且每一个步骤在可接受的时间内完成。

　　（3）确定性：算法的每一步骤都具有确定的含义，不会出现二义性。

　　（4）可行性：算法的每一步都必须是可行的，也就是说，每一步都能够通过执行有限次数完成。

　　3.算法设计的要求：正确性、可读性、健壮性、时间效率高和存储量低

　　（1）正确性：是指算法至少应该具有输入、输出和加工处理无歧义性、能正确反映问题的需求、能够得到问题的正确答案。

算法程序没有语法错误。
算法程序对于合法的输入数据能够产生满足要求的输出结果。
算法程序对于非法的输入数据能够得出满足规格说明的结果。
算法程序对于精心选择的，甚至刁难的测试数据都有满足要求的输出结果。　　
　　（2）可读性：算法设计的另一目的是为了便于阅读、理解和交流。

　　（3）健壮性：当输入数据不合法时，算法也能做出相关处理，而不是产生异常或莫名其妙的结果。

　　（4）时间效率高和存储量低：对于同一个问题，如果有多个算法能够解决，执行时间短的算法效率高，程序运行时所占用的内存或外部硬盘存储空间少的存储量低。

　　4.算法效率的度量方法

　　（1）事后统计方法：通过设计好的测试程序和数据，利用计算机计时器对不同算法编制的程序的运行时间进行比较，从而确定算法效率的高低。

　　（2）事前分析估算方法：在计算机程序编制前，依据统计方法对算法进行估算。一个用高级程序语言编写的程序在计算机上运行时所消耗的时间取决于下列因素：①算法采用的策略、方法。②编译产生的代码质量。③问题的输入规模。④机器执行指令的速度。

　　

二、算法时间复杂度

　　1.定义：在进行算法分析时，语句总的执行次数T(n)是关于问题规模n的函数，进而分析T(n)随n的变化情况并确定T(n)的数量级。算法的时间复杂度，也就是算法的时间度量度，记作：T(n) = O(f(n))。它表示随问题规模n的增大，算法执行时间的增长率和f(n)的增长率相同，称作算法的渐进时间复杂度，简称为时间复杂度。其中f(n)是问题规模n的某个函数。

　　2.推到大O阶的方法：①用常数1取代运行时间中的所有加法常数。②在修改后的运行次数函数中，只保留最高阶项。③如果最高阶项存在且不是1，则去除于这个项相乘的常数。

　　3.算法时间复杂度可以分为：O(1)常数阶、O(n)线性阶、O(n^2)平方阶、O(log2n)对数阶、O(mxn)平方阶变体。 

 

三、算法空间复杂度

　　1.算法的空间复杂度通过计算算法所需的存储空间实现，算法空间复杂度的计算公式记作：S(n) = O(f(n))，其中n为问题的规模，f(n)为语句关于n所占存储空间的函数。

　　2.当不限定词地使用“复杂度”时，通常都是指时间复杂度。