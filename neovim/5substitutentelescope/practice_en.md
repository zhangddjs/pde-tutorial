# Practice

## Reminder

1. Do not use the mouse when solving exercises.  
2. Try using `global`, `substitute`, and regular expressions to solve the problems.  
3. Open-source projects are not easy, if you enjoy my course and exercises, please consider giving a ★★★★★ star to support!

## Exercises

### Pattern

We have a text file named `pattern`:  

1. Use Ex mode to match the word `this` in `practice1` (excluding occurrences within `thisisalongword`).  

2. Given three URLs in `practice2`, use Ex mode to:  
   - (1) Match the URL `https://github.com/zhangddjs/pde-tutorial`.  
   - (2) Match the URL `https://github.com/zhangddjs?repository=pde-tutorial&link=https%3a%2f%2foutlook`.  
   - (3) Match the file path `C:\Windows\Users\Documents\pde-tutorial`.  

3. Filter `practice3` to keep only lines that end with a line number.  

### substitute

We have a text file named `substitute`:  

1. Add a space before and after each `+` in `practice1`.  

2. Surround every word in `practice2` with parentheses.  
   _HINT: The regular expression for matching a word is `\w+`, and escape characters (`\`) may be needed._  

3. Surround each occurrence of the word `this` in `practice3` with parentheses.  

4. Correct the misspelled word in `practice4` (the correct word is `line`).  

5. Fix the column alignment in `practice5` to the correct order: `column1 column2 column3 column4`.  
   _HINT: This requires four capture groups. Too many escape characters, any way to simplify?_  

6. In `practice6`, modify the text by:  
   - (1) Adding a colon (`:`) at the end of each line.  
   - (2) Appending the line number at the end of each line.  
   _HINT1: Use `\=` to call an expression in `{replace}`._  
   _HINT2: The expression to get the current line number is `line('.')`._  

7. In `practice7`, increment each date by 3 days.  
   _HINT1: The expression to get the matched value is `submatch(0)`._  
   _HINT2: The regular expression for matching a number is `\d+`, and escape characters (`\`) may be needed._  

---

### telescope

We have folders `subfolder1` and `subfolder2`, containing various text file formats:  

1. Use `telescope` to open the file `split.md` in a split window.  

2. Use `telescope` to:  
   - (1) Search for all files containing the word `txt` and send the results to the `quickfix` list.  
   - (2) Search for all files containing the word `markdown` with the file extension `.md`.  

   _HINT: Check the `telescope` keybindings documentation or run `:h telescope` for help._  
