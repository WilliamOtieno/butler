package docker

func GetFedoraCommands() []string {
	commands := []string{
		"sudo dnf remove docker docker-client docker-client-latest docker-common docker-latest docker-latest-logrotate docker-logrotate docker-selinux docker-engine-selinux docker-engine",
		"sudo dnf -y install dnf-plugins-core",
		"sudo dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo",
		"sudo dnf -y install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin",
		"sudo systemctl enable docker",
		"sudo systemctl start docker",
		"sudo groupadd docker",
		"sudo usermod -aG docker $USER",
		"newgrp docker",
		"docker run hello-world",
	}
	return commands
}
