Output docs/demo.gif
Set FontSize 16
Set Theme Dracula
Set Width 900
Set Height 500
Set Padding 20
Set Framerate 12

# Step 1: Show demo dir
Type "mkdir flowkit-demo && cd flowkit-demo"
Sleep 500ms
Enter
Sleep 1s

# Step 2: Init git
Type "git init"
Sleep 300ms
Enter
Sleep 1s

# Step 3: Create package.json
Type "echo '{\"dependencies\":{\"next\":\"14\"}}' > package.json"
Sleep 300ms
Enter
Sleep 1s

# Step 4: Run flowkit init (non-interactive)
Type "flowkit init -n my-app -w gitflow"
Sleep 500ms
Enter
Sleep 4s

# Step 5: Show generated files
Type "ls -la"
Sleep 300ms
Enter
Sleep 2s

# End
Sleep 1s
