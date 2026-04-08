# Input - Password
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

Type "go run examples/input/password/main.go"
Sleep 500ms
Enter
Sleep 2s

Type "s3cr3tP@ss"
Sleep 1s
Enter
Sleep 2s
