# Confirm - Default
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

Type "go run examples/confirm/default/main.go"
Sleep 500ms
Enter
Sleep 2s

# Toggle to No
Right
Sleep 800ms

# Toggle back to Yes
Left
Sleep 800ms

# Confirm
Enter
Sleep 2s
