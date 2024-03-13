package controllers

import (
	"celestial-audit/app/services"

	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/jjamieson1/celestial-sdk/utilities"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) LogItemLogEvent() revel.Result {
	var log models.ItemLog
	err := c.Params.BindJSON(&log)
	if err != nil {
		return utilities.HandleBadRequestError("itemLog", err.Error(), 400, c.Controller)
	}
	err = services.LogItemChange(log)
	return c.Result
}
