package modDatabase

import (
	"errors"

	"github.com/gatlinglab/libGatlingDatabaseModel"
	"github.com/gatlinglab/libGatlingDatabaseModel/dbModel"
)

type cDBLog struct {
	dbInst       dbModel.IWJDatabase
	modelHelper1 dbModel.IWJDBTM_Helper1
	tableName    string
	logid        int64
}

type funcDbLogCreate func(dbModel.IWJDatabase) *cDBLog

func newDBInfo(dbinst dbModel.IWJDatabase) *cDBLog {
	return &cDBLog{dbInst: dbinst, tableName: "tblInfo", logid: 10}
}
func newDBError(dbinst dbModel.IWJDatabase) *cDBLog {
	return &cDBLog{dbInst: dbinst, tableName: "tblError", logid: 10}
}

func (pInst *cDBLog) getMaxID() int64 {
	return pInst.logid
}
func (pInst *cDBLog) checkDatabaseStruct() error {
	pInst.modelHelper1 = libGatlingDatabaseModel.TDM_CreateHelper1(pInst.dbInst, pInst.tableName)
	if !pInst.modelHelper1.CheckTableExists() {
		return pInst.initTable()
	}

	sqlText := "select valueint from " + pInst.tableName + " where id = 1;"
	rows, err := pInst.dbInst.Query(sqlText)
	if err != nil {
		return err
	}

	var iID int64 = 0
	if rows.Next() {
		err = rows.Scan(&iID)
		if err != nil {
			return errors.New("load max id from database error: " + err.Error())
		}
	}
	if iID < 10 {
		return errors.New("database init error, max id value wrong")
	}

	return nil
}
func (pInst *cDBLog) initTable() error {
	err := pInst.modelHelper1.CreateTable()
	if err != nil {
		return err
	}

	sqlText := "insert into " + pInst.tableName + "(id, key, valueint) values(1, '__maxid', 10);\n"
	sqlText += "insert into " + pInst.tableName + "(id) values(2);\n"
	sqlText += "insert into " + pInst.tableName + "(id) values(3);\n"
	sqlText += "insert into " + pInst.tableName + "(id) values(4);\n"
	sqlText += "insert into " + pInst.tableName + "(id) values(5);\n"
	sqlText += "insert into " + pInst.tableName + "(id) values(6);\n"
	sqlText += "insert into " + pInst.tableName + "(id) values(7);\n"
	sqlText += "insert into " + pInst.tableName + "(id) values(8);\n"
	sqlText += "insert into " + pInst.tableName + "(id) values(9);\n"
	sqlText += "insert into " + pInst.tableName + "(id) values(10);"

	_, err = pInst.dbInst.ExecSql(sqlText)
	if err != nil {
		return errors.New("database init table id 10 failed: " + err.Error())
	}

	return nil
}
func (pInst *cDBLog) writelog(id int64, appInfo, log string) error {
	strSql := "insert into " + pInst.tableName + "(id, key, valuestr) values (?,'?','?');" +
		" update " + pInst.tableName + " set valueint=? where id=1;"
	_, err := pInst.dbInst.ExecSql(strSql, id, appInfo, log, id)
	//err := pInst.modelHelper1.InsertIDKeyValue(pInst.logid, appInfo, log)
	if err != nil {
		pInst.logid = id
	}

	return err
}
