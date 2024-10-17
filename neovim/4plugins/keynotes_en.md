# Plugins

## Tree

```shell
~/.config/nvim            # Path to configuration directory
├── lua
│   ├── config            # General configurations
│   │   ├── autocmds.lua  ## Autocommands (triggers)
│   │   ├── keymaps.lua   ## Key mappings
│   │   ├── lazy.lua      ## LazyVim configurations
│   │   └── options.lua   ## Neovim global options, such as UI settings
│   └── plugins           # Plugin configurations
│       ├── example.lua   ## Example configuration
│       ├── disabled.lua  ## Disabled plugins
│       ├── plugin1.lua
│       ├── **            ## Additional plugin files
│       └── pluginx.lua
├── lazy-lock.json        # Plugin version control
└── init.lua              # Entry point
```

## Important Key Bindings

`<leader>`: After pressing, you enter operator-pending mode, allowing you to call various plugin functionalities.
`<leader>l`: Opens the LazyVim plugin manager console.
`<leader>e`: Opens the file explorer.

## Steps to Install Plugins

1. Go to the plugin repository homepage, such as `github.com/<username>/<plugin-name>`.

2. Find the Installation section, specifically the LazyVim configuration. If not available (as in some vim plugins), the plugin configuration is generally `<username>/<plugin-name>`.
``` lua
return {
  '<username>/<plugin-name>'
}
```
_Note:_
_+ Some plugins may have additional dependencies that require configuration before installation._
_+ Some plugins require calling or implementing the `setup()` method. Check the documentation for details._

3. In the `~/.config/nvim/plugins` directory, create a `<plugin-name>.lua` or `<plugin-set>.lua` file, paste the configuration, and save it.

4. Restart Neovim, and the plugin will install automatically. If installation fails, press `i` (install) to retry.

## Steps to Roll Back Plugins

1. Before updating plugins, back up the `lazy-lock.json` file.

2. If issues arise after updating, restore the backup version of the `lazy-lock.json` file.

3. Restart Neovim.

4. Press `<leader>l` to open the LazyVim console. Highlight the problematic plugin and press `r` (restore) to roll it back.

## Key Mapping Configuration

### Method 1: Configuring in keymaps.lua

1. Open `~/.config/nvim/keymaps.lua`

2. If the `map` function is not declared, declare it at the top of the file
``` lua
local map = vim.keymap.set
```

3. Configure key mappings using map("mode", "key combination", "action (key sequence or expression)", {options}).
``` lua
map("n", "<leader>e", ":lua MiniFiles.open()<cr>", { desc = "Trigger Mini Files" })
```

### Method 2: Configuring in Plugin Files

Refer to the example in `example.lua`, where key mappings are set in the `keys` field of the plugin configuration:

``` diff
  {
    "simrat39/symbols-outline.nvim",
    cmd = "SymbolsOutline",
++  keys = { { "<leader>cs", "<cmd>SymbolsOutline<cr>", desc = "Symbols Outline" } },
    config = true,
  },
```