package modDatabase

type cLogManager struct {
	dbList      []*cDBLog
	fcbDbCreate funcDbLogCreate
	logid       int64
}

var g_infoLogManager *cLogManager = &cLogManager{fcbDbCreate: newDBInfo}
var g_errorLogManager *cLogManager = &cLogManager{fcbDbCreate: newDBError}

func getInfoLogManager() *cLogManager {
	return g_infoLogManager
}
func getErrorLogManager() *cLogManager {
	return g_errorLogManager
}

func (pInst *cLogManager) AddDatabase(connectstr, connectToken string) error {
	dbModel, err := getSingleDbModelManager().getDbInstance(connectstr, connectToken)
	if err != nil {
		return err
	}

	dbInst := pInst.fcbDbCreate(dbModel)

	err = dbInst.checkDatabaseStruct()

	if err != nil {
		return err
	}

	logid := dbInst.getMaxID()
	if logid > pInst.logid {
		pInst.logid = logid
	}

	pInst.dbList = append(pInst.dbList, dbInst)

	return nil
}
func (pInst *cLogManager) writeLog(appInfo, log string) error {
	var errResult error = nil
	pInst.logid++
	for _, dbInst := range pInst.dbList {
		err := dbInst.writelog(pInst.logid, appInfo, log)
		if err != nil {
			errResult = err
		}
	}

	return errResult
}
