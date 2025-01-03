package conf

func CreateEnv() error {
	setDefault()

	err := saveEnv()
	if err != nil {
		return err
	}

	return nil
}
