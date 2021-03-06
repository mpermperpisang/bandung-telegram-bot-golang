package helper

import (
	"context"
	"io/ioutil"
	"strconv"
	"time"

	"golang.org/x/oauth2/google"
	spreadsheet "gopkg.in/Iwark/spreadsheet.v2"
)

var sheet *spreadsheet.Sheet

func GoogleSheet(gsheet string) (file *spreadsheet.Sheet) {
	t := time.Now()
	year, _ := strconv.Atoi(t.Format("2006"))
	data, _ := ioutil.ReadFile("client_secret.json")
	conf, _ := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, _ := service.FetchSpreadsheet(gsheet)
	sheet, _ := spreadsheet.SheetByTitle(strconv.Itoa(year))

	return sheet
}
