# 课间小练习

## 温馨提示

1. 做习题时不要使用鼠标
2. 建议提前配置好golang treesitter 和 gopls
3. 虽然习题涉及到代码，但您无需理解逻辑和语法，无需任何Golang基础
4. 开源不易，如果您喜欢我的课程和习题，还望给个star⭐⭐⭐⭐⭐支持一下哦

## 习题

现在我们有代码文件`calculator.go`:

1. 请一键折叠文件里的所有方法，效果如：

``` go
//...
> func (c *Calculator) Add(x float64) {

> func (c Calculator) Sum() float64 { 

> func (c Calculator) Mean() float64 {
//...
```

2. 请通过Telescope，快速定位到方法`StdDev`处
_HINT: Telescope对treesitter进行了支持，可查阅资料并尝试在当前文件搜索treesitter object_

3. 将光标一键跳转回原位
_HINT: 2种方法，其中之一是一对很有用的快捷键_

4. 跳转回`StdDev`方法，将其一键展开，包括方法内的嵌套折叠部分也要展开

## 思考题

请借助工具为方法`StdDev`生成一个单元测试，并尝试解决lsp的报错信息
_HINT: lsp 报错sort未定义，意味着漏导了sort包，手动解决的方式是在顶部 `import{}` 内加上一行 `"sort"`，可以自动完成吗_
_Follow Up: 测试用例也可以生成吗_
