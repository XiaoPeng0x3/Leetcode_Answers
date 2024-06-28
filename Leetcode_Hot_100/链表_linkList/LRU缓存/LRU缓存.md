# LRU缓存

#  0 前言

刷力扣的时候突然发现竟然有`LRU`缓存这道题目，恰好之前做过`CSAPP`的`cache lab`,现在来梳理一下`LRU`的步骤。

# 1 缓存Cache

在学习一项新的概念的时候，要多提几个“为什么”，这样就可以帮助我们加深印象。因此，为什么要有缓存这个结构呢？

## 1.1 存储结构

在计算机早期设计的时候，其实并没有像现在这样复杂的结构(缓存、页表机制)，然而随着CPU的频率和速度大大加快，高效和性能日益成为一个必不可少的要求。因此，对于CPU来说，如果需要访问的数据就在寄存器里面，那么几乎不消耗时间就可以得到数据。反之，如果是在内存上存储的数据，则需要去访问内存去获得这个数据，会消耗几百个时钟周期。因此，为了消除CPU与内存之间的高速差距，于是设计者就不得不在内存和CPU之间另外搭建了一座桥梁，也就是最先的`L1缓存`。

`L1缓存`的速度小于CPU但是大于直接去内存里面访问的速度，因此就可以把一些数据放到`L1缓存`里面，从而减少获取需要的数据所需要的时间。但是这里有两个问题，`L1缓存`的速度很快但是大小也是很小，即便是在现代的处理器上，一般的大小也就是MB级别。而有的时候我们的数据却远远大于这个级别，还有一点更重要的就是`Cache miss `与`Cache hit`。假设这里只存在`L1`级别的缓存，如果CPU需要的数据正好在`L1`缓存里面，那么就是`Cache hit`,相反，如果不在`L1`缓存里面的话，那么就得驱逐`L1`缓存里面一些数据(这个时候就要用到几种不同的算法，例如`LRU`和`LFU`),然后再去内存里面寻找数据。

到这里，应该不难发现，即便是使用了`L1`缓存，也还是有很大的可能去内存里面访问数据。因此对应的，各位计算机前辈们就对应的发明了`L2`、`L3`缓存来解决这个问题。整个的结构就是在速度上：`L1 > L2 > L3`,在大小上：`L1 < L2 < L3`,不难想到从造价上：`L1 > L2 > L3`，因此整个存储系统就像一座金字塔一样，越靠近塔尖的存储器造价越贵、容量越小、速度越快；越向下的存储结构造价越便宜、容量越大、速度越慢。

![image-20240628101716380](C:\Users\13481\Desktop\刷题\Leetcode_Hot_100\链表_linkList\LRU缓存\LRU缓存.png)

一句话来说：**L1缓存是CPU的cache,L2缓存是L1的cache,L3缓存是L2的cache**,如此下去，形成了一层存储结构。

## 1.2 缓存的性质

缓存比较有趣的一个性质就是局部性，一般来说，对于已经访问的数据来说很有可能会再一次访问。比如像下面这行代码

```c
for (int i = 0 ; i < length; i++) {
    sum += arr[i]
}
```

这是一个简单计算数组元素和的代码，在这个循环中，每次都重复使用的数据就是`i,length,sum,arr`,其中像`i,length,sum`等这些元素都是存储在寄存器里面的，而对数组来说，正因为局部性这个小性质使得我们不用每次都去内存里面加载数据。

一个局部性很好的代码就要利用到缓存的一些性质，这里使用一个`NxN`的矩阵来进行演示，对于这个矩阵来说，一个求和程序可能是

```c
for (int i = 0; i < N; i++) {
    for (int j = 0; j < N; j++) {
        sum += arr[i][j]
    }
}
```

这个求和程序就是一行一行的去求矩阵大小，最后加到一起，很直观了。来分析一下为什么这个程序的局部性比较好。PS:数组是一个线性模型

- 首先去访问`arr[0][0]`,发现这个元素不在缓存中，然后发生一次`cache miss`,那么就需要去内存中访问数据(主要开销)，然后把`arr[0]`这一行数据载入缓存(举个例子)，那么在访问`arr[1] ~ arr[n - 1]`的时候就不会发生`cache miss`
- 所以，总的来说，每次只需要发生1次`miss`,总共发生`n`次即可

这是一个局部性不好的求法

```c
for (int i = 0; i < N; i++) {
    for (int j = 0; j < N; j++) {
        sum += arr[j][i]
    }
}
```

与上面那种求法相比，这种求法是按照一列一列来求的。原因就是没有去利用到缓存的局部性，整个过程如下，

- 首先去访问`arr[0][0]`,发现这个元素不在缓存中，对应的，把`arr[0]`缓存起来，但是，因为我们是按列求和，那么访问第一个元素时，会发生`miss`,之后的每一次访问元素都会发生`miss`,都需要去内存里面重新寻找，因此这两段代码的目的相同，但是效率却大大不同！
- 对此，我们也应当尽力写出局部性好的代码，进而提高程序的效率！

## 1.3 地址的寻找

刚才我们说到了缓存的一些性质，但还没有说在存储器上是怎么组织的。对于存储器结构来说，把它抽象为一个字节数组可以说是简单直观了。

实际上计算机很多的结构都是简单的线性结构，考虑这样一个机器，它的最长寻址地址是$2^m$个地址，那么我们一般就把这个机器叫做`m`位机器，对这`m`位`bit`来说，我们就可以把它的每一部分地址划分为不同的traget来使用。我们该怎么使用这部分地址呢？先来看一看你该怎么划分一个cache吧！

上文提到，把cache当作数组来看待是比较合适的，因为机器上的地址是`m`位的，所以一个cache也应该是`m`位的，这样方便`CPU`直接取址。因此，就可以把cache看作**一间一间的空房间**，等到CPU需要的时候直接破门而入即可(不要敲`nullptr`的门哟)。

那么，对CPU来说，这`m`位地址已经它也不知道在哪啊？这个时候就要考虑编号了。因此，这里就引出了三种不同的`cache`的组织方式：直接映射高速缓存、组相连高速缓存、全相连高速缓存。先从组相连高速缓存开始！

- 组相连高速缓存

  - 上面说到我们可以把`m bits`长的地址划分为类似于一个数组，那么对于这些数组来说，可以单独为一组，也可以每`E`个数组为一组(可以住单间、也可以几个人住一起)

  - 我们假设整个缓存的大小是`C`，假设每个数组的大小是`B`,那么，所划分的行数就是`C/B`这么多行。对于这些行来说，在组相连的方法里，我们可以把这些行每`E`个划分为一组(一个房间里有E个人)。

  - 因此，这样就建立好了索引，要找到某一块地址在那个房间里，首先要找到它在哪一组里面，我们假设有`S`个组,找到了在某一组内后，还需要找到它的组内的行号，找到行号后，再去取出它存储的一些信息

  - **总的来说，就是先找到房间号(有S个房间)，再找到房间里的那个人(人数为E)，再向那个人问话。**也就是**组选择、行匹配、字抽取**。

  - 最后，我们就可以把地址分块了，首先是要匹配组号，假设有`S`组，那么只需要找到比`S`大的那个$2^n$的数即可。比如说，我们现在有16组，那么组索引只需要4位二进制就可以完全表示出来，假设有8组，那么只需要3位二进制就可以把组索引表示出来。我们把占用组索引位的位数叫做`s`(小写s)，显然，$S = 2 ^ s$,同理，找到房间里的那个人我们也可以用二进制来表示出来，假设需要`t`位，那么显然$E = 2 ^ t$,这样，剩下来的就可以表示信息。一般的，是先确定信息所占的位数(假设是b位，那么块大小$B = 2 ^b$)，然后根据组号确定`s`的位数，剩下来的标记位就可以使用减法`m - b - s`得到。

  - 也就是说，知二就可以求三！这样就把`m`个字节长的地址划分好了。

    ![image-20240628143827271](C:\Users\13481\Desktop\刷题\Leetcode_Hot_100\链表_linkList\LRU缓存\通用结构.png)

- 直接映射高速缓存

  - 这其实是组相联高速缓存的一种特殊情况，特殊之处就在每个房间只有一个人，这样其实我们就不用寻找是哪个人了，那个房间即那人！

  - 寻找过程与其相似不再叙述。

    ![image-20240628144020159](C:\Users\13481\Desktop\刷题\Leetcode_Hot_100\链表_linkList\LRU缓存\直相连.png)

    ![image-20240628144150319](C:\Users\13481\Desktop\刷题\Leetcode_Hot_100\链表_linkList\LRU缓存\直接映射寻址.png)

- 全相联映射

  - 这个家伙可能有不尊重人权的嫌疑，这种方法说，为什么还要一个一个去寻找房间再去找人呢？直接集中力量办大事，所有人全部都去一个屋子里面吧。因此，所有人没有组之分，全部都在一个房间内。

  - 所以这个时候实际上就是只有一组，所以我们就可以把其它给多的位用来表示。

    ![image-20240628144327738](C:\Users\13481\Desktop\刷题\Leetcode_Hot_100\链表_linkList\LRU缓存\全相联组成.png)

    ![image-20240628144406250](C:\Users\13481\Desktop\刷题\Leetcode_Hot_100\链表_linkList\LRU缓存\全相联寻址.png)

## 1.4 缓存的hit与miss

cache的miss一般是对应的组号内的那个标记不存在或者有效位为0而发生的(注意，如果有效位为0，经过组号、行号都符合也算miss)。对应的,hit就是成功找到那个组号、行号的数据。

## 1.5 矩阵优化

这一部分其实算是一种额外的优化，矩阵间的运算几乎是机器学习和统计的基石，所以我们**需要根据实际条件来找到高效的运算方法**。一般的矩阵的运算可以通过行与列的乘积得到，可是我们在前文中已经探讨过了，当按列访问元素时，由于发生`cache miss`而会导致重复的去内存中寻址，从而引起计算效率的下降。对此，我觉得硬件设计完全可以背这个锅(反正不能怪我想不出来),因此就有了一种常见的运算方法，就是矩阵的分块(block)

对于两个`MxM`的矩阵来说，一个常见的乘法运算是

```c
int ans[M][M];
for (int i = 0; i < M; i++) {
    for (int j = 0; j < M; j++) {
        ans[i][j] = 0
        for (int k = 0; k < M; k++) {
            ans[i][k] += arr1[i][k] * arr2[k][j]
        }
    }
}
```

但是，这个计算效率并不高，主要是因为cache 的 miss而导致的，对此，我们可以使用分块矩阵乘法来解决，把整个矩阵划分为更小的子矩阵，同时还可以使用并行计算的方式来，就像numpy库里面的那样，这里是一个可行的分块矩阵乘法

```c++
#include <iostream>
#define M 64
#define BLOCK_SIZE 16

void multiplyBlock(int A[][M], int B[][M], int C[][M], int block_size) {
    for (int i = 0; i < block_size; i++) {
        for (int j = 0; j < block_size; j++) {
            int sum = 0;
            for (int k = 0; k < M; k++) {
                sum += A[i][k] * B[k][j];
            }
            C[i][j] += sum;
        }
    }
}

int main() {
    int A[M][M] = { /* 初始化矩阵A */ };
    int B[M][M] = { /* 初始化矩阵B */ };
    int C[M][M] = {0};

    for (int i = 0; i < M; i += BLOCK_SIZE) {
        for (int j = 0; j < M; j += BLOCK_SIZE) {
            for (int k = 0; k < M; k += BLOCK_SIZE) {
                // 有可能不是方阵
                multiplyBlock(&A[i][k], &B[k][j], &C[i][j], std::min(BLOCK_SIZE, M - i), std::min(BLOCK_SIZE, M - j));
            }
        }
    }

    // 打印结果矩阵C
    for (int i = 0; i < M; i++) {
        for (int j = 0; j < M; j++) {
            std::cout << C[i][j] << " ";
        }
        std::cout << std::endl;
    }

    return 0;
}
```

# 2 LRU算法

说了那么多铺垫知识，现在来看一看这个题目怎么写

题目描述：

> 请你设计并实现一个满足 [LRU (最近最少使用) 缓存](https://baike.baidu.com/item/LRU) 约束的数据结构。
>
> 实现 `LRUCache` 类：
>
> - `LRUCache(int capacity)` 以 **正整数** 作为容量 `capacity` 初始化 LRU 缓存
> - `int get(int key)` 如果关键字 `key` 存在于缓存中，则返回关键字的值，否则返回 `-1` 。
> - `void put(int key, int value)` 如果关键字 `key` 已经存在，则变更其数据值 `value` ；如果不存在，则向缓存中插入该组 `key-value` 。如果插入操作导致关键字数量超过 `capacity` ，则应该 **逐出** 最久未使用的关键字。
>
> 函数 `get` 和 `put` 必须以 `O(1)` 的平均时间复杂度运行。
>
>  
>
> **示例：**
>
> ```
> 输入
> ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
> [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
> 输出
> [null, null, null, 1, null, -1, null, -1, 3, 4]
> 
> 解释
> LRUCache lRUCache = new LRUCache(2);
> lRUCache.put(1, 1); // 缓存是 {1=1}
> lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
> lRUCache.get(1);    // 返回 1
> lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
> lRUCache.get(2);    // 返回 -1 (未找到)
> lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
> lRUCache.get(1);    // 返回 -1 (未找到)
> lRUCache.get(3);    // 返回 3
> lRUCache.get(4);    // 返回 4
> ```

可以看到，查询是根据`key`值去查询的，而且要求时间是`O(1)`的，所以第一时间想到的就是哈希表的结构也就是`HashMap`结构。此外还要求可以动态的调整位置，所以我们这里用到了双向链表。

思路：`Hashmap + DoubleLinkList`即可。

## 2.1 LRU算法的实现

LRU算法，全称为`Least Recently Used（最近最少使用）`，是一种常用的缓存淘汰策略和页面置换算法。这种算法主要用于决定在有限的内存空间中，当没有足够的资源保留所有数据时，应该优先释放或淘汰哪些数据。LRU算法的核心思想是，假设某个数据项在最近一段时间内没有被访问，那么在未来一段时间内，它很可能也不会被访问。基于这种观察，当缓存满载需要腾出空间时，LRU算法会选择最近最少使用的数据项进行淘汰。

下面我们来看一看实现思路：

- 首先是`get`方法，在原始题目中的描述是这样的
  ```java
  public int get(int key) {}
  ```

  目的就是去查询一个值，如果这个值存在，那么就返回对应的`value`,如果不存在的话就返回`-1`。需要注意的一点是，如果要查询的值是存在的，那么相当于这个值是被我们访问过的，所以需要把这个结点移动到链表的头部，这样我们就形成了**越靠近链表头部的结点是最近访问过的结点**。

- 其次就是`put`方法了。在题目中的原型是：

  ```java
  public void put(int key, int value) {}
  ```

  这个函数是细节最多的一个函数。题目的要求中还有一点需要注意，其中`cache`的容量是有限的。所以，当我们调用`pust`的时候就要注意当前`cache`是否还有额外的空间去添加元素。这个函数的思路就是：

  - 执行`put`后无非就两种情况：添加的元素存在和添加的元素不存在。
  - 对不存在的元素来说，需要更新当前`cache`的`size`,而且还要把新添加的结点移动到链表的头部，还要在哈希表中记录这个结点的`key`以便于我们的元素的查找。而在添加的时候就要注意，如果当前的`size == capacity`,那么再执行添加的时候就没有足够的空间了，这个时候就要体现出驱逐算法，也就是把尾部那个结点从链表中删除，同时也要减小`size`和删除哈希表中的`key`。
  - 对于存在的元素(key)来说,我们只需要去更新这个结点的`value`值即可，同时还要把这个结点放到链表的首位，表示这个结点是最近访问过的

- 其它的就是一些辅助代码了，设计链表的一些操作，这里就不多叙述了，下面是完整代码

  ```java
  // 每次get和put都要把对应的结点移动到最前面
  // 在get的时候，找不到返回-1，不需要移动；找到先返回值，再移动
  // put的时候要注意容积大小，超过容积要删除尾巴，也就是最近不使用的点
  // put还要注意是否存在，存在的话要更新里面的结点值，不存在的话需要添加
  
  class LRUCache {
      class DoubleLink{
          // 双向链表
          // prev,next
          public DoubleLink prev;
          public DoubleLink next;
          public int key;
          public int val;
          // 构造函数
          public DoubleLink(){}
          public DoubleLink(int key, int val) {
              this.key = key;
              this.val = val;
          }
      }
  
      // 初始化好一些变量
      public int capacity;
      public DoubleLink head, tail;
      public HashMap<Integer, DoubleLink> cache = new HashMap<>();
      public int size = 0;
      public LRUCache(int capacity) {
          this.capacity = capacity;
          this.head = new DoubleLink();
          this.tail = new DoubleLink();
          this.head.next = this.tail;
          this.tail.prev = this.head;
      }
      
      public int get(int key) {
          // 先去查询一下存不存在
          DoubleLink ans = cache.get(key);
          if (ans == null) {
              return -1;
          }
          // 移动到头部
          move2Head(ans);
          // 然后返回
          return ans.val;
      }
      
      public void put(int key, int value) {
          // 先看看存不存在
          // 如果存在，那么就更新val并且移动到头部
          DoubleLink ans = cache.get(key);
          if (ans == null) {
              // 说明不存在，添加进去
              ans = new DoubleLink(key, value);
              // 同时也更新哈希表
              cache.put(key,ans);
              size++;
              add2Head(ans);
              if (size > capacity) {
                  // 删除末尾
                  DoubleLink retail = tail.prev;
                  cache.remove(retail.key);
                  size--;
                  removeTail();
              }
          } else {
              ans.val = value;
  	        move2Head(ans);
          }
  
      }
      private void move2Head(DoubleLink node) {
          // 处理好当前关系
          node.prev.next = node.next;
          node.next.prev = node.prev;
          // 然后移动到头部
          add2Head(node);
      }
      private void add2Head(DoubleLink node) {
          head.next.prev = node;
  	    node.next = head.next;
  	    node.prev = head;
  	    head.next = node;
      }
      private void removeTail() {
          DoubleLink retail = tail.prev;
          retail.prev.next = tail;
          tail.prev = retail.prev;
      }
  }
  ```

  

## 2.2 cache lab 杂谈

在`cache lab`里面，同样也是使用了这个算法，只不过当时没有实现的那么复杂，就是使用了简单的数组进行模拟。对于这`cache`，多添加了一个`cnt`属性，每当执行`put`或者`get`的时候，就把这个`cnt`赋值为0，然后每次在遍历`cache`的时候，给这个`cnt`遍历加一，这样，对于很长时间没有访问`cache`块，它的`cnt`就会是最大值，那么当发生驱逐的时候就替换掉这个块就行了。总的来说，还是要多多注意细节啊！