package data

func Migrate() error {
	err := DB.AutoMigrate()
	return err
}
