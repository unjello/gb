package core

func VerifyConanExists() error {
	runner := NewOsCommandRunner()
	return runner.Run([]string{"conan", "--version"})
}
