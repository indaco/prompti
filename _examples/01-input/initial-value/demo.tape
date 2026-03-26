# Input - Initial Value and Validation
Output "demo.gif"

Require go

Set Shell "bash"
Set FontSize 16
Set Width 800
Set Height 300
Set Theme "Catppuccin Mocha"
Set TypingSpeed 80ms
Set Padding 15

Hide
Type "cd /Users/indaco/Code/GitHub/indaco/prompti"
Enter
Sleep 500ms
Show

Type "go run _examples/input/initial-value/main.go"
Sleep 500ms
Enter
Sleep 2s

# First try an invalid value
Type "not-a-number"
Sleep 500ms
Enter
Sleep 1s

# Clear and type a valid integer
Backspace 14
Sleep 500ms
Type "42"
Sleep 1s
Enter
Sleep 2s
