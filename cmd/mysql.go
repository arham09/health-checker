package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/cobra"
)

// mysqlCmd represents the mysql command
var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Check if MySQL is alive or not",
	Long:  `Check ping for MySQL.`,
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		db, _ := cmd.Flags().GetString("db")

		pingMysql(address, user, password, db)
	},
}

func init() {
	rootCmd.AddCommand(mysqlCmd)
	mysqlCmd.Flags().StringP("address", "a", "127.0.0.1:3306", "Host for MySQL")
	mysqlCmd.Flags().StringP("user", "u", "root", "User for MySQL")
	mysqlCmd.Flags().StringP("password", "p", "root", "Password for MySQL")
	mysqlCmd.Flags().StringP("database", "d", "root", "Database for MySQL")
}

func pingMysql(address, user, password, db string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, address, db)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = conn.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Ping sucessful")
}
