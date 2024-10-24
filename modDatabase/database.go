package modDatabase

func DB_AddInfoDatabase(connectstr, connectToken string) error {
	return getInfoLogManager().AddDatabase(connectstr, connectToken)
}
func DB_AddErrorDatabase(connectstr, connectToken string) error {
	return getErrorLogManager().AddDatabase(connectstr, connectToken)
}

func DB_WriteLogInfo(appInfo, log string) error {
	return getInfoLogManager().writeLog(appInfo, log)
}

func DB_WriteLogError(appInfo, log string) error {
	return getErrorLogManager().writeLog(appInfo, log)
}
