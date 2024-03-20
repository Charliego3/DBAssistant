package db

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charliego3/assistant/utility"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/progrium/macdriver/macos/appkit"
)

var (
	db *sqlx.DB

	//go:embed schema.sql
	create string
)

func InitializeDB() {
	path := utility.SupportPath("db")
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		response := utility.ShowAlert(
			utility.WithAlertTitle("Failed initialize Support Files"),
			utility.WithAlertMessage(err.Error()),
			utility.WithAlertShowHelp(true),
			utility.WithAlertStyle(appkit.AlertStyleCritical),
			utility.WithAlertOnHelpClicked(func(a appkit.Alert) bool {
				fmt.Println("alert help button clicked")
				return true
			}),
		)
		os.Exit(int(response))
	}

	path = filepath.Join(path, "assistant.sqlite")
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			res := utility.ShowAlert(
				utility.WithAlertMessage(err.Error()),
				utility.WithAlertStyle(appkit.AlertStyleCritical),
			)
			os.Exit(int(res))
		}
		_ = f.Close()
	}

	db, err = sqlx.Open("sqlite3", path)
	if err != nil {
		res := utility.ShowAlert(
			utility.WithAlertMessage(err.Error()),
			utility.WithAlertStyle(appkit.AlertStyleCritical),
		)
		os.Exit(int(res))
	}

	_, err = db.Exec(create)
	if err != nil {
		res := utility.ShowAlert(
			utility.WithAlertMessage(err.Error()),
			utility.WithAlertStyle(appkit.AlertStyleCritical),
		)
		os.Exit(int(res))
	}
}
