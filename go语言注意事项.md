# 注意事项

在刷题的过程中，常常会遇到一些不太熟悉的go语言的用法，所以在这里记录一下。

## 循环

- go里面只有`for` 循环,没有`while`循环。同时，还提供了一种新的`range`方法。第一种

  ```go
  for i := 0; i < n; i++
  ```

- 还可以直接基于`range`的方法去遍历，第一个返回的元素是索引，第二个返回的元素就是数组的值，例如我们要遍历一个数组
  ```go
  for index, val := range nums
  ```

  同时也可以忽略某些值，不去使用

  ```go
  for _, val := range nums // 忽略索引
  
  for index, _ := range nums // 忽略值
  ```

  

## make的声明

- 可以使用make来提前声明一个`slice`
  ```go
  ans := make([]int)
  ```

- 还可以指定长度
  ```go
  ans := make([]int, length) // 指定长度
  ```

- 还可以指定容量
  什么时候去指定容量呢，也就是在提前知道长度的情况下，我们就可以直接声明好长度。这样在使用append的时候就不会再去再分配空间上面消耗时间

  ```go
  ans := make([]int, length, capacity)
  ```

- 声明哈希表

  ```go
  ans := make(map[byte]int) // 创建一个键为byte, 值为int类型的哈希表
  ```

  

## 变量的交换

- go语言里面的交换两个元素简直和python一样
  ```go
  // 在c语言里面可能是
  int temp = a
  a = b
  b = temp
  
  // 但是在go语言里面只需要简单的一行
  a, b = b, a
  ```



## if语句的判断

在映射中查找键值对时，可以使用这种语法来判断键是否存在：

```go
package main

import "fmt"

func main() {
    nums := map[string]int{"one": 1, "two": 2, "three": 3}

    if val, ok := nums["two"]; ok {
        fmt.Println("Found:", val) // 输出: Found: 2
    } else {
        fmt.Println("Not found")
    }

    if val, ok := nums["four"]; ok {
        fmt.Println("Found:", val)
    } else {
        fmt.Println("Not found") // 输出: Not found
    }
}

```



## 切片

目前来说，字符串、数组切片都是支持切片的，而且切片的时间复杂度也是`O(1)`的。例如，我们想要删除第一个元素，我们可以这样写
```go
q = q[1:]
```

在一些数据结构，比如说使用stack的时候，就可以很方便的使用数组去模拟栈，这个时候也可以使用切片的性质。比如说要删除最后一个元素，就可以这样写

```go
stack = stack[ : len(stack) - 1]
```

最后那个区间刚好是开区间，所以最后那个元素不会保留下来





## 面向对象

在go语言里面初始化一个变量写法与其它语言不太一样，我觉得可以直接使用`key:value`的形式去构造成员变量。

```go
type book struct {
    name string
    price float32
}

// 直接使用构造函数的方式去构造可能会更好一些
book1 := book{name:"haha", price:12.5}
```

