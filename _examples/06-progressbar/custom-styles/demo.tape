# ProgressBar - Custom Styles
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

Type "go run _examples/progressbar/custom-styles/main.go"
Sleep 500ms
Enter

# Wait for progressbar to complete
Sleep 15s
