package homebrew

func GetHomebrewCommands() []string {
	commands := []string{
		"/bin/bash -c \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\"",
	}
	return commands
}
