# 课间小练习

## 温馨提示

1. 做习题时不要使用鼠标。
2. 做习题前先备份配置目录 `cp -a ~/.config/nvim{,.bak}`
3. 提前准备好MiniFiles插件（练习用，无需替换您当前习惯使用的文件树插件）
4. 如果您喜欢我的习题设计和笔记总结，还望给个star★★★★★支持一下哦

## 习题

我们现在有配置路径`~/.config/nvim`，先备份它 `cp -a ~/.config/nvim{,_bak}`。

1. 在`$HOME`目录下启动neovim（`cd ~ && nvim`），打开`init.lua`配置文件。

2. 使用`<leader>`键将屏幕分割成 左（Left），右上（R-UP），右下（R-DN） 共3块分屏。 
```
|‾‾‾‾‾‾|‾‾‾‾‾‾|
|      | R-UP |
| Left |──────|
|      | R-DN |
|______|______|
```

3. 先在右上（R-UP）分屏中使用MiniFiles打开`keymaps.lua`文件，然后再在右下（R-DN）分屏中打开`minifiles.lua`文件
```
|‾‾‾‾‾‾‾‾‾‾|‾‾‾‾‾‾‾‾‾‾‾|
|          |  keymaps  |
|          |    .lua   |
| init.lua |───────────|
|          | minifiles |
|          |    .lua   |
|__________|___________|
```

4. 打开了`minifiles.lua`文件后，将光标移动到右上（`keymaps.lua`）分屏，唤醒MiniFiles。会发现MiniFiles中的光标仍然定位在`minifiles.lua`。我们需要让MiniFiles在唤醒时智能定位到当前文件`keymaps.lua`。
已知，MiniFiles智能定位需要调用以下两个lua方法:
`MiniFiles.open(vim.api.nvim_buf_get_name(0))`
`MiniFiles.reveal_cwd()`
请:
(1) 将其绑定到快捷键`<leader>e`上（如果冲突可以绑定别的快捷键）
(2) 重启Neovim进入欢迎页后，一键恢复先前的分屏状态

5. 光标来到右下（`minifiles.lua`）分屏，唤醒MiniFiles，请尝试分屏打开`lazy-lock.json`（类似Neo-Tree的s键）。
已知，MiniFiles按s键分屏打开文件的配置如下(Powered by: [@xiantang](https://github.com/xiantang))：
```lua
-- Mini Files Configuration Script
-- Author: @xiantang https://github.com/xiantang
-- Description: It introduces the ability to open files in a vertical split 
-- directly from mini.files, triggered by the 's' key. 

  config = function()
    local window_handle
    local mini_files = require("mini.files")

    local group = vim.api.nvim_create_augroup("mini.files", { clear = true })
    vim.api.nvim_create_autocmd("User", {
      pattern = "MiniFilesBufferCreate",
      desc = "define commands/keymaps for mini.files",
      group = group,
      callback = function(args)
        window_handle = vim.api.nvim_get_current_win()
        -- mini_files.open(vim.api.nvim_buf_get_name(0))
        local buf_id = args.data.buf_id
        vim.keymap.set("n", "s", function()
          local entry = mini_files.get_fs_entry()
          if not entry then
            return
          end
          vim.api.nvim_win_call(window_handle, function()
            mini_files.close()
            vim.cmd.vsplit({ args = { entry.path } })
          end)
        end, {
          buffer = buf_id,
          desc = "open from mini.files in vertical split",
        })
      end,
    })
  end,
```
_HINT: 可参考example.lua的格式，或者查看minifiles插件的文档，看看如何配置这段代码_

```
|‾‾‾‾‾‾‾‾‾‾|‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾|
|          |     keymaps.lua       |
|          |                       |
| init.lua |───────────┬───────────|
|          | minifiles | lazy-lock |
|          |    .lua   |   .json   |
|__________|___________|___________|
```

6. 实现了分屏打开文件的功能后，您的同学想要学习`minifiles.lua`的代码，请将该分屏最大化方便查看。
```
|‾‾‾‾‾|‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾|
|     |        keymaps.lua         |
|     |──────────────────────┬─────|
|init |                      |     |
|.lua |                      |la...|
|     |    minifiles.lua     |.json|
|     |                      |     |
|_____|______________________|_____|
```
_HINT: 尝试探索一些功能和快捷键，比如<leader>，<c-w>等_

7. 按思考题2的要求修改配置，实现MiniFiles的懒加载。

## 思考题

1. 目前大多数Neovim插件都在仓库主页提供了通过LazyVim进行安装的配置，但仍然有一些插件（比如Vim插件）没有提供，请大家挖掘规律，并尝试安装和卸载[vim-peekaboo插件](https://github.com/junegunn/vim-peekaboo)

2. 当前配置的MiniFiles插件在Neovim启动时会自动加载（控制台中loaded状态），为了保障Neovim的启动效率，请修改配置，让MiniFiles懒加载，只有用户第一次按下<leader>e时才会被加载。
已知懒加载三要素：
(1) `require('mini.files').setup()`会初始化插件，想要实现懒加载，就不能在`init.lua`中调用，需要挪到`minifiles.lua`中
(2) 懒加载配置项`lazy = true`
(3) 插件加载的常用触发条件：Ex命令`cmd`配置项；快捷键`keys`配置项