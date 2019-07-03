package dbops

import (
	"GoProxyPoll/GoProxyPoll/defs"
	"testing"
)

var (
	testIp = &defs.Ip{Data:"127.0.0.1", Type:"HTTP"}
)

func clearTable() {
	dbConn.Exec(`TRUNCATE proxypoll`)
}

func TestMain(m *testing.M) {
	clearTable()
	m.Run()
	clearTable()
}

func TestAll(t *testing.T) {
	t.Run("AddIp", testAddIp)
	t.Run("NotExist", testNotExist)
	t.Run("Count", testCount)
	t.Run("UpdateIp", testUpdateIp)
	t.Run("Decrease", testDecrease)
	t.Run("AddAgain", testAddIp)
	t.Run("Random", testRandom)
	t.Run("AllIps", testAllIps)
	t.Run("DeleteIp", testDeleteIp)
}

func testAddIp(t *testing.T) {
	err := AddIp(testIp)
	if err != nil {
		t.Errorf("Error of AddIp: %v\n", err)
	}
}

func testNotExist(t *testing.T) {
	res, err := NotExist(testIp)
	if err != nil {
		t.Errorf("Error of NotExist: %v\n", err)
	}
	if res != false {
		t.Error("Res should be false")
	}
}

func testDeleteIp(t *testing.T) {
	err := DeleteIp(testIp)
	if err != nil {
		t.Errorf("Error of DeleteIp: %v\n", err)
	}
	if ok, _ := NotExist(testIp); !ok {
		t.Errorf("Delete ip failed")
	}
}

func testCount(t *testing.T) {
	res, err := Count()
	if err != nil {
		t.Errorf("Error of count: %v\n", err)
	}
	if res != 1 {
		t.Errorf("Wrong answer of Count")
	}
}

func testUpdateIp(t *testing.T) {
	err := UpdateIp(testIp, MIN_SCORE)
	if err != nil {
		t.Errorf("Error of UpdateIp: %v\n", err)
	}
	var score int
	stmtOut, _ := dbConn.Prepare(`SELECT  score FROM proxypoll WHERE ip = ?`)
	stmtOut.QueryRow(testIp.Data).Scan(&score)
	if score != MIN_SCORE {
		t.Errorf("Update ip failed")
	}
}

func testDecrease(t *testing.T) {
	err := Decrease(testIp)
	if err != nil {
		t.Errorf("Error of Decrease: %v\n", err)
	}
	if ok, _ := NotExist(testIp); !ok {
		t.Errorf("Decrease failed")
	}
}

func testRandom(t *testing.T) {
	_, err := Random()
	if err != nil {
		t.Errorf("Error of Random: %v\n", err)
	}
}

func testAllIps(t *testing.T) {
	_, err := AllIps()
	if err != nil {
		t.Errorf("Error of AllIps: %v\n", err)
	}
}