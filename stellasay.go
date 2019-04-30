package stella

import (
	"io"
	"os/exec"
)

func Stellasay(text string) (string, error) {
	cmd := exec.Command("/Users/emaas/go/bin", "-n")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	io.WriteString(stdin, text)
	stdin.Close()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
