package digcloudlog_dbpostgresql

import "github.com/digautos-library/digcloudlog"

// dig cloud log database postgresql

func DCLDP_AddPostgresqlToLogSystem(flag, dburl string) error {
	dbInst, err := DCLDP_NewPostgresql(flag, dburl)
	if err != nil {
		return err
	}
	err = digcloudlog.DCL_AddNewLogService(dbInst)
	return err
}

func DCLDP_NewPostgresql(flag, dburl string) (*CDbPostgres, error) {
	dbInst := newDbPostgres()
	err := dbInst.Initialize(flag, dburl)
	if err != nil {
		return nil, err
	}
	return dbInst, nil
}
