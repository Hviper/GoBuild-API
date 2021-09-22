**笔记**

1.Easy and safe casting from one type to another in Go
https://github.com/spf13/cast


go中string与[]byte的互换
```go
   // string to []byte
    s1 := "hello"
    b := []byte(s1)
    
    // []byte to string
    s2 := string(b)
```

go断言：

语法：

    value, ok := x.(T)
    //value, _ := s.(Person)  判断s是否为Person类型，是就返回value

    //类型断言还可以配合 switch 使用，示例代码如下：
    switch a.(type) {
        case int:
            fmt.Println("the type of a is int")
        case string:
            fmt.Println("the type of a is string")
        case float64:
            fmt.Println("the type of a is float")
        default:
            fmt.Println("unknown type")
    }

---
