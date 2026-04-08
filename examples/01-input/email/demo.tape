# Input - Email Validation
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

Type "go run examples/input/email/main.go"
Sleep 500ms
Enter
Sleep 2s

# Try an invalid email first
Type "not-an-email"
Sleep 500ms
Enter
Sleep 1s

# Clear and type a valid email
Backspace 12
Sleep 500ms
Type "user@example.com"
Sleep 1s
Enter
Sleep 2s
