package utils_run_tsx_script

import "os/exec"

func RunTsxScript(filePath string) error {
	cmd := exec.Command("tsx", filePath)
	return cmd.Run()
}
