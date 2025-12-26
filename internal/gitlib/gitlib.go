package gitlib

import (
	"bufio"
	"context"
	"os/exec"
)

func GetDiffs(ctx context.Context, head, base string) ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", head, base)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(stdout)
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if scanner.Err() != nil {
		cmd.Process.Kill()
		cmd.Wait()
		return nil, scanner.Err()
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
