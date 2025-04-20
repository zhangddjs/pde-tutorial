# 编程环境搭建

## 涉及插件

- nvim-treesitter & language parser
- nvim-lspconfig & language server
- nvim-treesitter-textobjects
- mason
- [lspsaga](https://github.com/nvimdev/lspsaga.nvim)
- [darcula-dark.nvim](https://github.com/xiantang/darcula-dark.nvim)

## 重要快捷键

`gd`: go to definition 跳转到定义
`gr`: go to reference 跳转到引用
`gI`: go to implementation 跳转到实现
`gy`: 跳转到声明/签名
`k`: 查看方法或变量信息，类似vscode hover，再次按下可focus
`zM`: 折叠所有方法
`zo`: 打开光标下的折叠内容（不包含嵌套折叠）
`zO`: 打开光标下的所有折叠内容

## 文本对象
`f` function: 方法，可结合操作符使用，比如`yaf`, `vaf`, `yif`
`a` argument: 参数，`daa` 删除参数, `]a` 跳转下一个参数
`s` struct: 结构体，`yas`, `das`
`c` class: 类，`yac`
`i` indent: 缩进体，`yii`, `yai`
`o` conditional: if或for循环, `yac`
`q` quote: 引号内容，`yaq`, `diq`，等价于 `"`, `ya"`
`b` block: 括号内容，`yab`, `cib`
`t` tag: html标签，`yat`