# 课间小练习

## 温馨提示

1. 做习题时不要使用鼠标
2. 提前配置好git和lazygit环境
3. 开源不易，如果您喜欢我的课程和习题，还望给个star⭐⭐⭐⭐⭐支持一下哦

## 习题

现在我们有:

- 文本文件 `text.md`
- 分支 `gitintegration/main`
- 分支 `gitintegration/feat`

1. 请通过lazygit切换到`gitintegration/feat`分支  
_HINT: 分支不在local，可以从Remotes子菜单获取_

2. 查看`text.md`第一行的commit message和Author  
_HINT: gitsigns 集成了丰富的可视化功能，可以试试看_

3. 查看`text.md`的commit历史  
_HINT: 2种方法，Telescope和vim-fugitive插件_

## 思考题

通过lazygit将`gitintegration/feat`rebase到`gitintegration/main`

步骤：
1. checkout到`gitintegration/feat`
2. 光标移动到`gitintegration/main`，按r键rebase

问题：
1. 在rebase时会遇到冲突，这是因为有其他分支修改了相同文件的相同代码块并被提前merge进了`gitintegration/main`，请解决它  
_HINT: 2种方法，Lazygit和Merge Tool_

2. 解决完冲突continue后发现一直有类似的重复冲突要解决，每个commit都要解决一次，那么有没有方法减少这种重复工作？
_HINT: Squash_
