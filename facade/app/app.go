package app

import (
	"github.com/weeon/weeon/contract"
	"github.com/weeon/weeon/internal/app"
)

func Get() contract.App {
	return app.Get()
}
