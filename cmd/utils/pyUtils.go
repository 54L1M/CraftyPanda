package utils

func CreatePyEnv(projectName string, appDir string) error {
	if err := ExecuteCmd("python3",
		[]string{"-m", "venv", "env"},
		appDir); err != nil {
		return err
	}

	return nil
}
