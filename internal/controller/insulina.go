package controller

import (
	"api-golang/pkg/models"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func AddRegistry(ctx *gin.Context) {
	// Returns JSON with OK Status

	// Creates Sheets Service
	sheetsService, err := sheets.NewService(context.Background(), option.WithCredentialsFile(os.Getenv("CREDENTIALS_FILE_PATH")))
	if err != nil {
		response := models.InsulinaResponse{Msg: "ERROR", Data: fmt.Sprintf("Failed to create sheets client: %v", err)}
		ctx.JSON(500, response)
		return
	}

	// Sets needed generic values to new row

	// Create a custom fixed timezone offset for GMT-3 (UTC-3)
	zoneOffset := -3 * 60 * 60 // -3 hours in seconds
	timezone := time.FixedZone("GMT-3", zoneOffset)

	rangeValue := "A:D"
	date := time.Now().In(timezone).Format("2006-01-02 15:04:05")
	month := time.Now().Month().String()

	// Defines the values to be written
	values := [][]interface{}{
		{ctx.Param("value"), date, month, ctx.Param("dose")},
	}

	// Creates the value range object
	valueRange := &sheets.ValueRange{
		Values: values,
	}

	// Make the API call to write data to the Google Sheet
	_, err = sheetsService.Spreadsheets.Values.Append(os.Getenv("SPREADSHEET_ID"), rangeValue, valueRange).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		response := models.InsulinaResponse{Msg: "ERROR", Data: fmt.Sprintf("Failed to update data to sheet: %v", err)}
		ctx.JSON(500, response)
		return
	}

	// If everything is OK returns an success message
	response := models.InsulinaResponse{Msg: "SUCCESS", Data: fmt.Sprintf("Data added to Google Sheet successfully. Value: %s", ctx.Param("value"))}
	ctx.JSON(200, response)
}
