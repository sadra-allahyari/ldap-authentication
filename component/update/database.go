package update

import (
	"os/exec"
)

func RunUpdateFile() error {

	cmd := exec.Command("sh", "update.sh")

	cmd.Dir = "config"

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
