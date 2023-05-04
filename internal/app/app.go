package app

import "github.com/weeon/weeon/contract"

var (
	defaultApp contract.App
)

func Set(app contract.App) {
	defaultApp = app
}

func Get() contract.App {
	return defaultApp
}
