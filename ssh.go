package containerh

import (
	"fmt"
	"net/http"
	"time"

	"github.com/loafoe/easyssh-proxy/v2"
)

func sshConfigFor(hostUser, privateIP, bastionHost string) (*easyssh.MakeConfig, error) {
	ssh := &easyssh.MakeConfig{
		User:   hostUser,
		Server: privateIP,
		Port:   "22",
		Proxy:  http.ProxyFromEnvironment,
		Bastion: easyssh.DefaultConfig{
			User:   hostUser,
			Server: bastionHost,
			Port:   "22",
		},
	}
	return ssh, nil
}

func runCommands(commands []string, ssh *easyssh.MakeConfig) (string, error) {
	var stdout, stderr string
	var done bool
	var err error

	for i := 0; i < len(commands); i++ {
		stdout, stderr, done, err = ssh.Run(commands[i], 5*time.Minute)
		if err != nil {
			return stdout, fmt.Errorf("failed [%s] [%v]: %w", stderr, done, err)
		}
	}
	return stdout, nil
}
