# Neovim Philosophy: Make Changes Repeatable

## Definition of the Dot Command

Repeat the last change.

## Definition of Change

1. A content operation.
2. Changes made during the time between leaving and re-entering normal mode.
3. In insert mode, moving the cursor resets the change state.

## Tips for Making Changes Repeatable

### Tip 1: Reduce Unnecessary Movements

`A` = `$a`
`I` = `^i`
`O` = `ko`
`C` = `d$a` = `Da`

### Tip 2: Use `a` and `i` Wisely

`daw`, `ciq`, `dia` ...

`daw` > `bdw` > `dbx`

### Tip 3: Use Visual Mode Only When Necessary

`daw` > `vawd`, `ciw` > `viwc`, `ciq` > `viqc`

### Tip 4: Step Back to Move Forward

When you have to move the cursor in insert mode, consider using backspace to delete the text and then reinsert it.

### Tip 5: Make Movements Repeatable (Dot Paradigm)

`j.`, `w.`, `;.`, `*.`, `n.`
