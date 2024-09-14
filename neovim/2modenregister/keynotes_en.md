# Mode & Register

## Normal Mode

### navigation (motion)

`j`,`k`,`l`,`h`
`w`,`b`,`e`
`f`,`t`
`^`,`0`,`$`
`gg`, `G`, `200G`, `<C-u>`, `<C-d>`

### undo, redo

`u`, `<C-r>`

undo unit: 
1. a content operation
2. a change between leave normal mode and reenter normal mode;

### operator-pending mode

**operator + motion = action**

`d{motion}`, `y{motion}`, `c{motion}` ...

### yank

`yy`, `yj`, `yh`
`yt,`, `yf,`
`yaw`, `yiw`, `y2w`, `y2aw`, `2yaw`

### content operation

#### update

`d{motion}`, `c{motion}`, `x`, `r`, `R`, `gR`

#### indent

`>{motion}`, `<{motion}`

#### switch case

`gu{motion}`, `gU{motion}`, `~`

#### number

`<C-a>`, `78<C-x>`

## Insert Mode

`<C-w>` delete a word
`<C-u>` delete cpntent before cursor
`<C-a>` navigate to start of line (need config)
`<C-e>` navigate to end of line (need config)

`<C-r>` paste register, 0, =.
`<C-o>` enter insert-normal mode

## Visual Mode

`o` switch edge
`gv` reselect
`viw` select a word
`V` select lines
`<C-v>` block select, `A` or `I` add or remove content simultaneously for selected lines
`<C-g>` enter select mode. 


## Ex Mode

### how to trigger

`:`, `/`, `?` under normal mode or visual mode, `<C-r>=` under insert mode.
`q:` under normal mode or `<C-f>` under Ex-mode to trigger history, 

### commands

`w` save, `q` quit, `q!` compulsory quit,  `wq` save and quit
`!ls` shell command ls, `r!date` insert current date, `r! date "+\%s"` insert timestamp

## Register

`""` Unnamed register
`"0` Yank register
`"a` - `"z` Named registers
`"=` Expression register
...

`<C-r>{register}` insert register's content under insert mode
`"0p` paste yank register's content
`""p` same with `p`
`"adaw` delete a word and store into named register a