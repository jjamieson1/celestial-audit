package services

import (
	"database/sql"

	"github.com/jjamieson1/celestial-sdk/models"

	"celestial-audit/app"

	"github.com/revel/revel"
)

func LogItemChange(itemLog models.ItemLog) error {

	query := "insert into item_log (item_id, action, activity, business_id) VALUES (?,?,?,?)"
	stmt, err := app.DB.Prepare(query)
	if err != nil {
		revel.AppLog.Errorf("LogItemChange: error preparing query for the item_log with error: %v", err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(itemLog.ItemId, itemLog.Action, itemLog.Activity, itemLog.TenantId)
	if err != nil {
		revel.AppLog.Errorf("error adding item to the item log with error: %v", err.Error())
		return err
	}
	return nil
}

func GetLogItems(itemId, businessId string) (itemLog []models.ItemLog, err error) {

	var query string

	if itemId == "0" {
		query = "SELECT item_id, action, activity, created from item_log where business_id = ?"
	} else {
		query = "SELECT item_id, action, activity, created from item_log where item_id = ? AND business_id = ?"
	}

	stmt, err := app.DB.Prepare(query)
	if err != nil {
		revel.AppLog.Errorf("Error performing query: %v, Error: %v", query, err.Error())
	}
	revel.AppLog.Debugf("using query: %v", query)
	defer stmt.Close()

	i := models.ItemLog{}
	var results *sql.Rows

	if itemId == "0" {
		results, err = stmt.Query(businessId)
	} else {
		results, err = stmt.Query(itemId, businessId)
	}

	if err != nil {
		revel.AppLog.Errorf("Error performing query: %v", err.Error())
	}

	for results.Next() {
		err := results.Scan(
			&i.ItemId,
			&i.Action,
			&i.Activity,
			&i.Created,
		)
		if err != nil {
			revel.AppLog.Errorf("Error scanning results from get all roles, error: %v", err.Error())
		}
		itemLog = append(itemLog, i)
	}
	return itemLog, err
}
