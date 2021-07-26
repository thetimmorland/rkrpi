package main

import (
	"bufio"
	"log"
	"net"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const schema = `
CREATE TABLE IF NOT EXISTS msgs (
	id INTEGER PRIMARY KEY NOT NULL,
	msg TEXT NOT NULL
);

CREATE TRIGGER IF NOT EXISTS deleteold
	AFTER INSERT ON msgs
	WHEN (SELECT count(*) FROM msgs) > 10000000
BEGIN
	DELETE FROM msgs
	WHERE id < (SELECT min(id) from msgs) + 10000;
END;
`

func readMsgs(reader *bufio.Reader, numMsgs int) []string {
	var msgs []string

	for i := 0; i < numMsgs; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		line = strings.TrimSpace(line)
		msgs = append(msgs, line)
	}

	return msgs
}

func main() {
	db, err := sqlx.Connect("sqlite3", "rkrpi.db")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)

	conn, err := net.Dial("tcp", "192.168.76.1:10110")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		msgs := readMsgs(reader, 100)

		tx := db.MustBegin()
		for _, msg := range msgs {
			tx.MustExec("INSERT INTO msgs (msg) VALUES ($1)", msg)
		}
		tx.Commit()
	}
}
