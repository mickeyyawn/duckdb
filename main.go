
package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	db, _ := sql.Open("duckdb", "")

	defer db.Close()

	db.Exec(`CREATE TABLE person (id INTEGER, name VARCHAR)`)
	db.Exec(`INSERT INTO person VALUES (42, 'John')`)

	var (
		id   int
		name string
	)
	row := db.QueryRow(`SELECT id, name FROM person`)
	_ = row.Scan(&id, &name)
	fmt.Println("id:", id, "name:", name)



//EvtUniqueID,OptionID,LocalCode,ExDT,OutturnStyleCD,PaytypeCD,RatioOld,RatioNew,RateValue01,DefaultOptionFlag,SourceFile
//40659958920,1,MTSUY,,NEWO,S,1,3,0,T,data/2023-12/e2020_20231207_LST_EVT_1.txt

	db.Exec(`CREATE TABLE splits (EvtUniqueID BIGINT, OptionID INTEGER, LocalCode VARCHAR, ExDT VARCHAR, OutturnStyleCD VARCHAR, PaytypeCD VARCHAR, RatioOld INTEGER, RatioNew INTEGER, RateValue01 INTEGER, DefaultOptionFlag VARCHAR, SourceFile VARCHAR)`)
	check(db.Exec(`COPY splits FROM 'splits.csv'`))


	//tx, _ := db.Begin()
	temp := db.QueryRowContext(context.Background(), "SELECT COUNT(*) FROM splits")
	var count int64
	check(temp.Scan(&count))

	//fmt.Println("count: ", count)
	fmt.Printf("%d", count)


}

func check(args ...interface{}) {
	err := args[len(args)-1]
	if err != nil {
		panic(err)
	}
}