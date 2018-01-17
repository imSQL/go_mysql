package imSQL

import (
	"testing"
)

func TestFindAllUsers(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	conn.SetCharset("utf8")
	conn.SetCollation("utf8_general_ci")
	conn.MakeDBI()

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	allusers, err := FindAllUserInfo(db, 1, 0)
	if err != nil {
		t.Error(allusers, err)
	}

}
