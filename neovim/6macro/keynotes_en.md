# Macros

## Important Shortcuts

- `q{register}`: Start recording  
- `q{Register}`: Append commands  
- `q`: Stop recording  
- `@{register}`: Play back the stored macro  
- `@@`: Replay the last played macro  
- `[n]@{register}`, `[n]@@`: Execute serially `n` times  
- `:[range]normal @{register}`: Execute in parallel, applying the macro to each selected line  

## Usage Tips

1. **Use register 'q'**: Quickly start recording by double-tapping `q` for convenience.  
2. **End recording with a movement command**: This allows macros to be executed efficiently with numerical repetitions.  
3. **Normalize cursor position—move to the beginning of the line before starting and stopping recording**: Ensuring a consistent cursor position prevents unintended misalignment.  

## Golden Rule

When recording a macro, ensure every command can be **repeated reliably** for consistent execution.
