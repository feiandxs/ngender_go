# NGender_go

根据中文姓名猜测其性别 go version

- 不到20行纯Python代码(核心部分)
- 无任何依赖库
- 82%的准确率
- 可用于猜测性别
- 也可用于判断名字的男性化/女性化程度

项目来自 [https://github.com/observerss/ngender](https://github.com/observerss/ngender)  
准确度可能有问题，数据也很久没更新，但好用，还好玩。  
甚至可以用于生产环境，辅助用户分类等场景。  
我只是把它翻译成了go版本。

# 使用

```go
package main

import (
    "fmt"
    "github.com/feiandxs/ngender_go"
)

func main() {

	guesser, err := ngender_go.NewGuesser()
	if err != nil {
		fmt.Println("创建猜测器失败：%v", err)
	}
	
	name := "张三"
	
	gender, rate := guesser.Guess(name)
    rateStr := fmt.Sprintf("%.2f%%", rate*100)
    fmt.Printf("姓名：%s，性别：%s， 概率： %s\n", name, gender, rateStr)
	
}
```

