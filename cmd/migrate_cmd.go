package cmd

import (
	"github.com/netlify/git-gateway/conf"
	"github.com/netlify/git-gateway/storage/dial"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var migrateCmd = cobra.Command{
	Use:  "migrate",
	Long: "Migrate database structures. This should create new tables and add missing columns and indexes. Currently it does nothing.",
	Run: func(cmd *cobra.Command, args []string) {
		execWithConfig(cmd, migrate)
	},
}

func migrate(globalConfig *conf.GlobalConfiguration, config *conf.Configuration) {
	db, err := dial.Dial(globalConfig)
	if err != nil {
		logrus.Fatalf("Error opening database: %+v", err)
	}
	defer db.Close()
}
