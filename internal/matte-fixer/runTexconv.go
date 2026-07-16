package mattefixer

import (
	"os/exec"
)

func RunTexconv(texconvPath string, args []string) error {
	cmd := exec.Command(texconvPath, args...)
	return cmd.Run()
}
