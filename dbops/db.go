package dbops

import (
	"GoProxyPoll/GoProxyPoll/defs"
	"database/sql"
	"errors"
	"log"
	"math/rand"
	"time"
)

func AddIp(ip *defs.Ip) error {
	stmtIns, err := dbConn.Prepare(`INSERT INTO proxypoll (ip, protocol, score)
									VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(ip.Data, ip.Type, MAX_SCORE)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func NotExist(ip *defs.Ip) (bool, error) {
	stmtOut, err := dbConn.Prepare(`SELECT id FROM proxypoll WHERE ip = ?`)
	if err != nil {
		return false, err
	}

	var id int
	err = stmtOut.QueryRow(ip.Data).Scan(&id)
	if err == sql.ErrNoRows {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	defer stmtOut.Close()
	return false, nil
}

func DeleteIp(ip *defs.Ip) error {
	stmtDel, err := dbConn.Prepare(`DELETE FROM proxypoll WHERE ip = ?`)
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(ip.Data)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func UpdateIp(ip *defs.Ip, score int) error {
	stmtUp, err := dbConn.Prepare(`UPDATE proxypoll SET score = ? 
									WHERE ip = ?`)
	if err != nil {
		return err
	}
	_, err = stmtUp.Exec(score, ip.Data)
	if err != nil {
		return err
	}
	defer stmtUp.Close()
	return nil
}

func Decrease(ip *defs.Ip) error {
	stmtOut, err := dbConn.Prepare(`SELECT score FROM proxypoll WHERE ip = ?`)
	if err != nil {
		return err
	}
	var score int
	err = stmtOut.QueryRow(ip.Data).Scan(&score)
	if err != nil {
		return nil
	}
	if score == MIN_SCORE {
		err = DeleteIp(ip)
	} else {
		err = UpdateIp(ip, score-1)
	}
	defer stmtOut.Close()
	return err
}

func Count() (int, error) {
	stmtOut, err := dbConn.Prepare(`SELECT count(*) FROM proxypoll`)
	if err != nil {
		return 0, err
	}
	var cnt int
	err = stmtOut.QueryRow().Scan(&cnt)
	if err != nil {
		return 0, err
	}
	defer stmtOut.Close()
	return cnt, nil
}

func Random() (*defs.Ip, error) {
	stmtOut, err := dbConn.Prepare(`SELECT ip, protocol FROM proxypoll WHERE score = ?`)
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.Query(MAX_SCORE)
	if err != nil {
		return nil, err
	}
	var ips []*defs.Ip

	if columns, _ := rows.Columns(); len(columns) == 0 {
		stmtOut, err = dbConn.Prepare(`SELECT ip, protocol FROM proxypoll`)
		if err != nil {
			return nil, err
		}
		rows, err = stmtOut.Query()
		if err != nil {
			return nil, err
		}
		if columns, _ := rows.Columns(); len(columns) == 0 {
			err = errors.New("There is no proxy in database.")
			return nil, err
		}
	}

	var cnt int
	for rows.Next() {
		var ip, protocol string
		if err = rows.Scan(&ip, &protocol); err != nil {
			return nil, err
		}
		c := &defs.Ip{Data:ip, Type:protocol}
		ips = append(ips, c)
		cnt++
	}
	defer stmtOut.Close()
	rand.Seed(time.Now().UnixNano())
	log.Println(cnt)
	num := rand.Intn(cnt)
	return ips[num], nil
}

func AllIps() ([]*defs.Ip, error) {
	stmtOut, err := dbConn.Prepare(`SELECT ip, protocol FROM proxypoll`)
	if err != nil {
		return nil, err
	}

	var ips []*defs.Ip
	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ip, protocol string
		if err = rows.Scan(&ip, &protocol); err != nil {
			return nil, err
		}
		c := &defs.Ip{Data:ip, Type:protocol}
		ips = append(ips, c)
	}
	defer stmtOut.Close()
	return ips, nil
}