# 版本控制十倍速

## 涉及插件

- [gitsigns.nvim](https://github.com/lewis6991/gitsigns.nvim)
- [vim-fugitive](https://github.com/tpope/vim-fugitive)
- [lazygit](https://github.com/jesseduffield/lazygit)

## 重要快捷键

### Neovim

`<leader>gg`: 唤醒lazygit

### Lazygit

`hljk`: 光标导航，切换菜单或菜单内选项
`[`, `]`: 切换菜单内子菜单
`<c-o>`: 复制选中文件名、分支名、commit hash...
`/`: 按条件过滤和搜索匹配菜单项
`z`, `<c-z>`: undo, redo

#### 文件菜单

`<space>`: 暂存光标下文件、目录
`<cr>`: 选择要暂存的更改块(hunk)
`d`: 丢弃光标下文件、目录
`a`: 一键暂存所有文件
`s`: 创建stash
`c`: 创建commit

#### 分支菜单

`n`: 新建分支
`d`: 删除分支
`p`, `f`: pull / fetch 远程分支
`P`: push修改
`o`: 一键创建Pull Request
`<space>`: checkout到选中分支
`<cr>`: 查看选中分支的commits
`r`: rebase选中分支

#### Commit菜单

`s`: 向下squash选中commit
`d`: 删除commit
`C`: cherry-pick 选中commit
`V`: 粘贴cherry-pick的commit
`i`: 交互式编辑commit
`<c-j>`, `<c-k>`: 将选中commit向下或向上挪
`<cr>`: 查看和编辑commit改动的文件
`b`: 开启Bisect
`+`: commit graph