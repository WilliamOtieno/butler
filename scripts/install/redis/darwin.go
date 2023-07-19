package redis

func GetDarwinCommands() []string {
	commands := []string{
		"brew --version",
		"brew install redis",
		"brew services start redis",
		"brew services info redis",
	}
	return commands
}
