# 课间小练习

## 温馨提示

1. 做习题时不要使用鼠标。
2. 请尝试使用global，substitute和正则表达式来解题
3. 开源不易，如果您喜欢我的课程和习题，还望给个star★★★★★支持一下哦

## 习题

### pattern

现在我们有文本文件`pattern`:

1. 通过ex mode匹配practice1中的单词`this`（thisisalongword中的this不算）

2. 现有practice2的3个url，请通过ex mode：
- (1) 匹配url `https://github.com/zhangddjs/pde-tutorial`
- (2) 匹配url `https://github.com/zhangddjs?repository=pde-tutorial&link=https%3a%2f%2foutlook`
- (3) 匹配url `C:\Windows\Users\Documents\pde-tutorial`

3. 筛选practice3，只保留末尾有行号文本行

### substitute

现在我们有文本文件`substitute`:

1. 请给practice1的每个加号前后加上空格

2. 请给practice2的每个单词前后都加上小括号
_HINT: 已知匹配单词的正则表达式是`\w+`，必要时需加上转义符`\`_

3. 请给practice3的每个`this`单词前后加上小括号

4. 请给practice4的错误单词改正确（正确单词是`line`）

5. 请给practice5的4列数据修正位置（column1 column2 column3 column4）
_HINT: 这里需要用到4个捕获组，转义符太多，有什么办法可以简化？_

6. 现有practice6，请给：
- (1) 每行末尾加上冒号
- (2) 每行末尾加上行号
_HINT1: 已知在{replace}中调用表达式使用`\=`_
_HINT2: 已知获取行号的表达式为`line('.')`_

7. 请给practice7的3个日期往后推迟3天
_HINT1: 已知获取当前匹配项的表达式为`submatch(0)`_
_HINT2: 匹配数字的正则表达式为`\d+`，必要时需加上转义符`\`_

### telescope
TODO: 有待验证

现有文件夹`subfolder1`, `subfolder2`，里面包含了多种格式的文本文件：

1. 请使用telescope分屏打开文件`split.md`
2. 请使用telescope：
- (1) 查找所有包含单词`txt`的文件，将结果放入`quickfix`
- (2) 查找所有包含单词`markdown`且后缀名为`md`的文件

_HINT: 可以查看telescope的快捷键文档，或者:h telescope来获取帮助_