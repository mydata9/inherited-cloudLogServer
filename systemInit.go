package main

import (
	"cloudLogServer/modDatabase"
	"cloudLogServer/modUtility"
	"errors"
	"fmt"
	"strconv"
)

func systemInit() error {

	var err error = nil
	var iCount int = 0
	for iLoop := 1; iLoop < 10; iLoop++ {
		err = databaseInit(iLoop)
		if err != nil {
			iCount++
		} else {
			fmt.Println("database init failed:", iLoop, err)
			break
		}
	}

	return err
}

func databaseInit(index int) error {
	url := modUtility.Config_ReadString("URL" + strconv.Itoa(index))
	token := modUtility.Config_ReadString("TOKEN" + strconv.Itoa(index))
	logType := modUtility.Config_ReadString("LOGTYPE" + strconv.Itoa(index)) // 1: info; 2: error

	if logType == "1" {
		return modDatabase.DB_AddInfoDatabase(url, token)
	} else if logType == "2" {
		return modDatabase.DB_AddErrorDatabase(url, token)
	} else {
		return errors.New("log type error")
	}
}
