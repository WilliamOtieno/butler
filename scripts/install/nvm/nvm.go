package nvm

func GetNVMCommands() []string {
	commands := []string{
		"wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash",
	}
	return commands
}
