# Practice

## Reminder

1. Do not use the mouse when doing exercises.
2. Before starting the exercise, back up your configuration directory with: `cp -a ~/.config/nvim{,.bak}`.
3. Prepare the MiniFiles plugin in advance (for practice only, no need to replace your current preferred file explorer).
4. If you like my exercise design and notes, please consider giving a star★★★★★ to show your support!

## Exercises

We have a configuration located at `~/.config/nvim`. First, back it up using: `cp -a ~/.config/nvim{,_bak}`.

1. Start Neovim in the `$HOME` directory (`cd ~ && nvim`), and open the `init.lua` configuration file.

2. Use the `<leader>` key to split the screen into three sections: left (Left), right-up (R-UP), and right-down (R-DN).
```
|‾‾‾‾‾‾|‾‾‾‾‾‾|
|      | R-UP |
| Left |──────|
|      | R-DN |
|______|______|
```

3. In the right-up (R-UP) section, use MiniFiles to open the `keymaps.lua` file, and then in the right-down (R-DN) section, open the `minifiles.lua` file.
```
|‾‾‾‾‾‾‾‾‾‾|‾‾‾‾‾‾‾‾‾‾‾|
|          |  keymaps  |
|          |    .lua   |
| init.lua |───────────|
|          | minifiles |
|          |    .lua   |
|__________|___________|
```

4. After opening the `minifiles.lua` file, move the cursor to the right-up (`keymaps.lua`) section and activate MiniFiles. You will find that the MiniFiles cursor still points to `minifiles.lua`. We need MiniFiles to smartly jump to the current file `keymaps.lua` when activated.
It’s known that smart positioning requires calling these two Lua methods:
`MiniFiles.open(vim.api.nvim_buf_get_name(0))`
`MiniFiles.reveal_cwd()`
Please:
(1) Bind this function to the `<leader>e` key (or another key if it conflicts).
(2) After restarting Neovim and entering the welcome page, restore the previous split screen state with single key.

5. Move the cursor to the right-down (`minifiles.lua`) section, activate MiniFiles, and try to open `lazy-lock.json` in a vertical split (similar to Neo-Tree's "s" key functionality).
It is known that MiniFiles uses the following configuration to open a file in a split when pressing "s" (Powered by: [@xiantang](https://github.com/xiantang)):
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
_HINT: Try referring to the format used in example.lua, or check the documentation for the minifiles plugin to see how to configure this code._

```
|‾‾‾‾‾‾‾‾‾‾|‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾|
|          |     keymaps.lua       |
|          |                       |
| init.lua |───────────┬───────────|
|          | minifiles | lazy-lock |
|          |    .lua   |   .json   |
|__________|___________|___________|
```

6. After implementing the split-open functionality, your classmate wants to study the code in `minifiles.lua`. Please maximize that split screen for easier viewing.
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
_HINT: Try exploring some functionalities and shortcuts like <leader>, <c-w>, etc._

## Challenge

Most Neovim plugins provide LazyVim installation configurations on their repository homepage, but some (like Vim plugins) do not. Explore the patterns and try installing and uninstalling the [vim-peekaboo plugin](https://github.com/junegunn/vim-peekaboo).
