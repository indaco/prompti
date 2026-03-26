# Toggle - Default
Output "demo.gif"

Require go

Set Shell "bash"
Set FontSize 16
Set Width 800
Set Height 250
Set Theme "Catppuccin Mocha"
Set TypingSpeed 80ms
Set Padding 15

Hide
Type "cd /Users/indaco/Code/GitHub/indaco/prompti"
Enter
Sleep 500ms
Show

Type "go run _examples/toggle/default/main.go"
Sleep 500ms
Enter
Sleep 2s

# Toggle between options
Right
Sleep 800ms
Left
Sleep 800ms
Right
Sleep 800ms

# Confirm
Enter
Sleep 2s
