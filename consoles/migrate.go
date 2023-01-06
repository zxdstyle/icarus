package consoles

import "github.com/zxdstyle/icarus"

type MigrateProvider struct {
	models []any
}

func NewMigrateProvider(models []any) MigrateProvider {
	return MigrateProvider{models}
}

func (s MigrateProvider) Signature() string {
	return "migrate"
}

func (s MigrateProvider) Description() string {
	return "Migrate models"
}

func (s MigrateProvider) Handle(args ...string) error {
	return icarus.DB().AutoMigrate(s.models...)
}
