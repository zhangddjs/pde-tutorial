# Substitute and Telescope

## Pattern  

patterns start with a forward slash, and `{pattern}` can be either a string or a regular expression.  
```vim
/{pattern}  
# Example:
  /monday  # Matches the word "monday"
  /        # Matches the last searched pattern (content stored in the `/` register)
```

### Using Patterns with Ex Commands  
When combined with an Ex command, `{pattern}` applies the command to the next match.  
```vim
/{pattern}/[ex cmd]  
# Example:  
  /monday/d  # Deletes the next line containing the word "monday"
```

### `global` Command  
When combined with the `global` command before the slash, `{pattern}` applies to all matches in the file.
```vim
g[!]/{pattern}/[ex cmd]  

# Examples:  
  g/monday/d   # Deletes all lines containing the word "monday"
  g!/monday/d  # Keeps only the lines containing the word "monday"
```  

### `substitute` Command  

#### Replace First Match in the Current Line  
The `substitute` command replaces the first occurrence of `{pattern}` in the current line with `{replace}`.
```vim
s/{pattern}/{replace}  
# Example:  
  s/monday/tuesday  # Replaces the first "monday" in the current line with "tuesday"
```

#### Replace in a Specified Range  
Use `[range]` to specify the lines where replacements should occur.
```
[range]s/{pattern}/{replace}  
# Example:  
  %s/monday/tuesday  # Replaces the first "monday" in each line of the entire file with "tuesday"
```

#### Using Flags  
Flags modify the behavior of the `substitute` command. Common flags include:  
- `g`: global, replaces all occurrences in each line.  
- `c`: confirm, asks for confirmation before each replacement.  
```vim
[range]s/{pattern}/{replace}/[flags]  

# Examples:  
  %s/monday/tuesday/g   # Replaces all occurrences of "monday" with "tuesday" in the entire file.  
  %s/monday/tuesday/gc  # Replaces all occurrences of "monday" with "tuesday" but asks for confirmation before each change.  
```

## Case Study  

Given the following text:  
```text
周一 monday monday
周二 tuesday
周三 wednesday
12.16 monday
12.17 tuesday
12.18 wednesday
```

### Pattern + Ex Command  
Command: `/monday/d`  
```diff
-- 周一 monday monday
周二 tuesday
周三 wednesday
12.16 monday
12.17 tuesday
12.18 wednesday
```

### `global` Command  
Command: `g/monday/d`  
```diff
-- 周一 monday monday
周二 tuesday
周三 wednesday
-- 12.16 monday
12.17 tuesday
12.18 wednesday
```  

Command: `g!/monday/d`  
```diff
周一 monday monday
-- 周二 tuesday
-- 周三 wednesday
12.16 monday
-- 12.17 tuesday
-- 12.18 wednesday
```  

### `substitute` Command  
Command: `s/monday/tuesday`  
```diff
-- 周一 monday monday
++ 周一 tuesday monday
周二 tuesday
周三 wednesday
12.16 monday
12.17 tuesday
12.18 wednesday
```  

Command: `%s/monday/tuesday`  
```diff
-- 周一 monday monday
++ 周一 tuesday monday
周二 tuesday
周三 wednesday
-- 12.16 monday
++ 12.16 tuesday
12.17 tuesday
12.18 wednesday
```  

Command: `%s/monday/tuesday/g`  
```diff
-- 周一 monday monday
++ 周一 tuesday tuesday
周二 tuesday
周三 wednesday
-- 12.16 monday
++ 12.16 tuesday
12.17 tuesday
12.18 wednesday
```  

### Using Regular Expressions  

Command: `%s/\(.*\) \(.*\)/\2 \1/g` (Swap positions)  
```diff
-- 周一 monday monday
-- 周二 tuesday
-- 周三 wednesday
-- 12.16 monday
-- 12.17 tuesday
-- 12.18 wednesday
++ monday 周一 monday
++ tuesday 周二
++ wednesday 周三
++ monday 12.16
++ tuesday 12.17
++ wednesday 12.18
```  

## Differences and Relationship Between `global` and `substitute`  

- `global` applies to the entire text by default, while `substitute` applies only to the current line.  
- With a range, both `global` and `substitute` apply to the selected lines.  
- `global` is typically used with Ex commands, while `substitute` is combined with flags.  
- `global` is mainly used for filtering and selection, while `substitute` is used for batch replacement.  
