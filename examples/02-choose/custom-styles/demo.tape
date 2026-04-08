# Choose - Custom Styles
Output "demo.gif"

Require go

Set Shell "bash"
Set FontSize 16
Set Width 800
Set Height 400
Set Theme "Catppuccin Mocha"
Set TypingSpeed 80ms
Set Padding 15

Hide
Type "cd /Users/indaco/Code/GitHub/indaco/prompti"
Enter
Sleep 500ms
Show

Type "go run examples/choose/custom-styles/main.go"
Sleep 500ms
Enter
Sleep 2s

# Browse items
Down
Sleep 800ms
Down
Sleep 800ms

# Select carbonara
Enter
Sleep 2s
