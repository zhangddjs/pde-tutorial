# 查找与替换

## 匹配模式

匹配模式由斜杠开头，{pattern}内容可以是字符串或正则表达式
``` vim
/{pattern}  
# 示例：
/monday # 匹配单词monday
/       # 匹配上一次的pattern(/寄存器中的内容)
```

### 与ex命令结合
{pattern}和ex命令结合，作用于下一个匹配项
``` vim
/{pattern}/[ex cmd]  
# 示例：/monday/d 删除下一个出现单词monday的行
```

### global
斜杠前结合`global`命令，作用于所有匹配项
``` vim
g[!]/{pattern}/[ex cmd]

# 示例：
g/monday/d   # 删除所有出现单词monday的行
g!/monday/d  # 只保留出现单词monday的行
```

### substitute
#### 替换当前行匹配项
`substitute`命令可将当前行第一个匹配项替换成{replace}的内容
``` vim
s/{pattern}/{replace}
# 示例：s/monday/tuesday 将当前行第一个monday替换成tuesday
```

#### 指定替换范围
通过`[range]`指定范围，可替换范围内的匹配项
```
[range]s/{pattern}/{replace}
# 示例：%s/monday/tuesday 将所有行第一个monday替换成tuesday
```

#### 使用Flags
通过指定flag，控制替换行为。常用flags：
- `g`：`g`lobal，作用于所有列。
- `c`：`c`onfirm，每次替换前需要确认。
``` vim
[range]s/{pattern}/{replace}/[flags]  

# 示例:
%s/monday/tuesday/g  # 将所有monday替换成tuesday
%s/monday/tuesday/gc # 将所有monday替换成tuesday，每个替换都需确认
```

## 案例讲解
假设目前有如下文本：
``` text
周一 monday monday
周二 tuesday
周三 wednesday
12.16 monday
12.17 tuesday
12.18 wednesday
```

### pattern+ex命令
`/monday/d`
``` diff
-- 周一 monday monday
周二 tuesday
周三 wednesday
12.16 monday
12.17 tuesday
12.18 wednesday
```

### `global`命令
`g/monday/d`
``` diff
-- 周一 monday monday
周二 tuesday
周三 wednesday
-- 12.16 monday
12.17 tuesday
12.18 wednesday
```

`g!/monday/d`
``` diff
周一 monday monday
-- 周二 tuesday
-- 周三 wednesday
12.16 monday
-- 12.17 tuesday
-- 12.18 wednesday
```

### `substitute`命令
`s/monday/tuesday`
``` diff
-- 周一 monday monday
++ 周一 tuesday monday
周二 tuesday
周三 wednesday
12.16 monday
12.17 tuesday
12.18 wednesday
```

`%s/monday/tuesday`
``` diff
-- 周一 monday monday
++ 周一 tuesday monday
周二 tuesday
周三 wednesday
-- 12.16 monday
++ 12.16 tuesday
12.17 tuesday
12.18 wednesday
```

`%s/monday/tuesday/g`
``` diff
-- 周一 monday monday
++ 周一 tuesday tuesday
周二 tuesday
周三 wednesday
-- 12.16 monday
++ 12.16 tuesday
12.17 tuesday
12.18 wednesday
```

### 正则表达式
`%s/\(.*\) \(.*\)/\2 \1/g` 交换位置
``` diff
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

## `global`和`substitute`区别与联系

- `global`默认作用于全部文本，`substitute`默认为当前行
- 加上range后，`global`和`substitute`都会作用于range选中的所有行
- `global`通常和ex命令结合使用，`substitute`和flag结合使用。
- `global`的场景通常是过滤和筛选，`substitute`的场景则是批量替换。