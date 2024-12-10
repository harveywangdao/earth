package main

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
	"log"
	"strconv"
	"time"
)

func QueryRow() error {
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"127.0.0.1:19000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
	})
	defer func() {
		if _, err := conn.Exec("DROP TABLE example"); err != nil {
			log.Println(err)
		}
	}()
	if _, err := conn.Exec("DROP TABLE IF EXISTS example"); err != nil {
		log.Println(err)
	}
	if _, err := conn.Exec(`
		CREATE TABLE example (
			  Col1 UInt8
			, Col2 String
			, Col3 FixedString(3)
			, Col4 UUID
			, Col5 Map(String, UInt8)
			, Col6 Array(String)
			, Col7 Tuple(String, UInt8, Array(Map(String, String)))
			, Col8 DateTime
		) Engine = Memory
	`); err != nil {
		log.Fatal(err)
		return err
	}

	scope, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	batch, err := scope.Prepare("INSERT INTO example")
	if err != nil {
		log.Fatal(err)
		return err
	}
	for i := 0; i < 1000; i++ {
		if _, err := batch.Exec(
			uint8(i),
			"ClickHouse",
			fmt.Sprintf("%03d", uint8(i)),
			uuid.New(),
			map[string]uint8{"key": uint8(i)},
			[]string{strconv.Itoa(i), strconv.Itoa(i + 1), strconv.Itoa(i + 2), strconv.Itoa(i + 3), strconv.Itoa(i + 4), strconv.Itoa(i + 5)},
			[]any{
				strconv.Itoa(i), uint8(i), []map[string]string{
					{"key": strconv.Itoa(i)},
					{"key": strconv.Itoa(i + 1)},
					{"key": strconv.Itoa(i + 2)},
				},
			},
			time.Now()); err != nil {
			log.Fatal(err)
			return err
		}
	}
	if err := scope.Commit(); err != nil {
		log.Fatal(err)
		return err
	}
	row := conn.QueryRow("SELECT * FROM example")
	var (
		col1             uint8
		col2, col3, col4 string
		col5             map[string]uint8
		col6             []string
		col7             any
		col8             time.Time
	)
	if err := row.Scan(&col1, &col2, &col3, &col4, &col5, &col6, &col7, &col8); err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("row: col1=%d, col2=%s, col3=%s, col4=%s, col5=%v, col6=%v, col7=%v, col8=%v\n", col1, col2, col3, col4, col5, col6, col7, col8)

	rows, err := conn.Query("SELECT * FROM example")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&col1, &col2, &col3, &col4, &col5, &col6, &col7, &col8); err != nil {
			log.Fatal(err)
			return err
		}
		log.Printf("row: col1=%d, col2=%s, col3=%s, col4=%s, col5=%v, col6=%v, col7=%v, col8=%v\n", col1, col2, col3, col4, col5, col6, col7, col8)
	}
	return nil
}

func test1() {
	if err := QueryRow(); err != nil {
		log.Fatal(err)
		return
	}
}

func test2() {
	conn, err := connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	rows, err := conn.Query(context.Background(), "SELECT name,toString(uuid) as uuid_str FROM system.tables LIMIT 5")
	if err != nil {
		log.Fatal(err)
		return
	}
	for rows.Next() {
		var (
			name, uuid string
		)
		if err := rows.Scan(
			&name,
			&uuid,
		); err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("name: %s, uuid: %s", name, uuid)
	}
}

func connect() (driver.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:19000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := conn.Ping(context.Background()); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		log.Fatal(err)
		return nil, err
	}
	return conn, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	test1()
}
