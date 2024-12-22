package migrations

import (
	"os"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		// add up queries...

		// get environment variable
		email := os.Getenv("SUPERUSER_EMAIL")
		password := os.Getenv("SUPERUSER_PASSWORD")
		if email == "" || password == "" {
			return nil
		}

		superusers, err := app.FindCollectionByNameOrId(core.CollectionNameSuperusers)
		if err != nil {
			return err
		}

		record := core.NewRecord(superusers)

		// note: the values can be eventually loaded via os.Getenv(key)
		// or from a special local config file
		record.Set("email", email)
		record.Set("password", password)

		return app.Save(record)
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
