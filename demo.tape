# Generates a demo using https://github.com/charmbracelet/vhs

Output demo.gif

Require go

Set Shell "bash"
Set FontSize 32
Set Width 660
Set Height 280
Set Framerate 30
Set WindowBar Colorful
Set Padding 40
Set BorderRadius 10
Set Margin 20
Set MarginFill "#607D8B"

Type "go run ./examples/demo"
Sleep 500ms
Enter

Sleep 34s
