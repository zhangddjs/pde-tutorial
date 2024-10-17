# 插件

## 配置目录结构

``` shell
~/.config/nvim            # 配置目录的路径
├── lua
│   ├── config            # 通用配置
│   │   ├── autocmds.lua  ## 触发器
│   │   ├── keymaps.lua   ## 快捷键映射
│   │   ├── lazy.lua      ## LazyVim配置
│   │   └── options.lua   ## Neovim全局配置，如UI等
│   └── plugins           # 插件配置
│       ├── example.lua   ## 样例
│       ├── disabled.lua  ## 禁用插件
│       ├── plugin1.lua
│       ├── **
│       └── pluginx.lua
├── lazy-lock.json        # 插件版本控制
└── init.lua              # 入口文件
```

## 重要快捷键

`<leader>`: 按下后进入操作符待决模式，和其它按键组合可调用不同插件的功能
`<leader>l`: LazyVim插件管理控制台
`<leader>e`: 启动文件树

## 安装插件步骤

1. 进入插件仓库主页，比如 `github.com/<用户名>/<插件名>`
2. 找到Installation部分的LazyVim配置，如果没有（比如某些vim插件），则安装插件的配置通常是`<用户名>/<插件名>`
``` lua
return {
  '<用户名>/<插件名>'
}
```
_注意：_
_+ 个别插件会有额外依赖，安装前注意进行配置_
_+ 个别插件需要调用或实现setup()方法，注意看文档_

3. 在`~/.config/nvim/plugins`目录下创建`<插件名>.lua`或`<插件集>.lua`文件，将配置粘贴进去并保存
4. 重启Neovim，插件就会自动安装完毕，如果安装失败，可以按i键(install)再次尝试。

## 回滚插件步骤

1. 更新插件前先备份`lazy-lock.json`文件
2. 更新完成后如果有问题，可以将`lazy-lock.json`文件回滚
3. 重启Neovim
4. `<leader>l`打开控制台，将问题插件按r键(restore)回滚。

## 配置快捷键映射

### 方式一：在Keymaps.lua中进行配置

1. 打开`~/.config/nvim/keymaps.lua`
2. 如果`map`方法没有声明，就在顶部先声明
``` lua
local map = vim.keymap.set
```
3. 配置快捷键 map("模式", "按键组合", "触发效果(按键序列或表达式)", {其它选项})
``` lua
map("n", "<leader>e", ":lua MiniFiles.open()<cr>", { desc = "Trigger Mini Files" })
```

### 方式二：在插件配置文件中进行配置

参考example.lua中的样例，在`keys`字段中配置快捷键：

``` diff
  {
    "simrat39/symbols-outline.nvim",
    cmd = "SymbolsOutline",
++  keys = { { "<leader>cs", "<cmd>SymbolsOutline<cr>", desc = "Symbols Outline" } },
    config = true,
  },
```
