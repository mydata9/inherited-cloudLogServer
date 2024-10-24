package modDatabase

import (
	"errors"

	"github.com/gatlinglab/libGatlingDatabaseModel"
	"github.com/gatlinglab/libGatlingDatabaseModel/dbModel"
)

type cDbModelManager struct {
	dbMap map[string]dbModel.IWJDatabase
}

var g_singleDbModelManager *cDbModelManager = &cDbModelManager{dbMap: make(map[string]dbModel.IWJDatabase)}

func getSingleDbModelManager() *cDbModelManager {
	return g_singleDbModelManager
}

func (pInst *cDbModelManager) getDbInstance(connectstr, connecttoken string) (dbModel.IWJDatabase, error) {
	dbInst, exists := pInst.dbMap[connectstr]
	if exists {
		return dbInst, nil
	}
	dbInst, err := pInst.connect(connectstr, connecttoken)

	if dbInst == nil {
		return nil, err
	}

	pInst.dbMap[connectstr] = dbInst

	return dbInst, nil
}
func (pInst *cDbModelManager) connect(connectstr, connecttoken string) (dbModel.IWJDatabase, error) {
	dbInst := libGatlingDatabaseModel.GDM_CreateSqlDB(connectstr, connecttoken)
	if dbInst == nil {
		return nil, errors.New("database connect failed")
	}

	err := dbInst.Connect()
	if err != nil {
		return nil, err
	}

	if dbInst.GetDatabaseVersion() == "" {
		return nil, errors.New("database connected, but no version info")
	}

	return dbInst, nil
}
