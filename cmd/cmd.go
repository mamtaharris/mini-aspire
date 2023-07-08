package cmd

import (
	"context"

	"github.com/mamtaharris/mini-aspire/pkg/database"
	"github.com/mamtaharris/mini-aspire/pkg/logger"
	"github.com/mamtaharris/mini-aspire/pkg/server"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var cmd = &cobra.Command{
	Use:   "mini-aspire",
	Short: "Aspire Loan Service",
}

func Execute(ctx context.Context, db *gorm.DB) error {
	cmd.AddCommand(startServer(ctx, db))
	cmd.AddCommand(runMigration(db))
	cmd.AddCommand(runSeeder(db))
	if err := cmd.Execute(); err != nil {
		return err
	}

	return nil
}

func startServer(ctx context.Context, db *gorm.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.Start(ctx, db)
		},
	}
}

func runMigration(db *gorm.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Run Migration",
		RunE: func(cmd *cobra.Command, args []string) error {
			tx := db.Begin()
			for _, migrate := range database.AutoMigrate(tx) {
				if err := migrate.Run(tx); err != nil {
					tx.Rollback()
					logger.Log.Error(err.Error())
					return nil
				}
			}
			tx.Commit()
			logger.Log.Info("Migration completed!")
			return nil
		},
	}
}

func runSeeder(db *gorm.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "seed",
		Short: "Run Lending Service Seeds",
		RunE: func(cmd *cobra.Command, args []string) error {
			tx := db.Begin()
			for _, seed := range database.AllSeeds(tx) {
				if err := seed.Run(tx); err != nil {
					tx.Rollback()
					logger.Log.Error(err.Error())
					return nil
				}
			}
			tx.Commit()
			logger.Log.Info("Seed completed!")
			return nil
		},
	}
}
