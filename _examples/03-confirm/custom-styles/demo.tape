# Confirm - Custom Styles
Output "demo.gif"

Require go

Set Shell "bash"
Set FontSize 16
Set Width 800
Set Height 350
Set Theme "Catppuccin Mocha"
Set TypingSpeed 80ms
Set Padding 15

Hide
Type "cd /Users/indaco/Code/GitHub/indaco/prompti"
Enter
Sleep 500ms
Show

Type "go run _examples/confirm/custom-styles/main.go"
Sleep 500ms
Enter
Sleep 2s

# Toggle to No and back
Right
Sleep 800ms
Left
Sleep 800ms

# Confirm
Enter
Sleep 2s
