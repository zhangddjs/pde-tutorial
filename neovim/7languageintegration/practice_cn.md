# 课间小练习

## 温馨提示

1. 做习题时不要使用鼠标
2. 建议提前配置好golang treesitter 和 gopls
3. 虽然习题涉及到代码，但您无需理解逻辑和语法，无需任何Golang基础
4. 开源不易，如果您喜欢我的课程和习题，还望给个star⭐⭐⭐⭐⭐支持一下哦

## 习题

现在我们有代码文件`calculator.go`:

1. 请一键折叠文件里的所有方法，效果如：

TODO:
``` go
//...
func (c *Calculator) Add(x float64) {

func (c Calculator) Sum() float64 { 

func (c Calculator) Mean() float64 {
//...
```

2. 请通过Telescope，快速定位到方法`StdDev`处

3. 请一键展开方法`StdDev`，包括方法内的嵌套折叠部分也要展开

## 思考题

1. 拷贝方法`StdDev`，粘贴到文件最后，会发现此时需要我们手动输入一个空行才能和上一个方法隔开。有什么办法可以优化？
_HINT: 比如在拷贝时包含方法顶部的空行，可通过编辑treesitter-textobj实现_

2. 请借助工具为代码生成一个单元测试文件
