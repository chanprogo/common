package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	adminOne := "host=10.190.24.252 port=5432 user=admin password=4IeggAix@2021 dbname=creative_project sslmode=disable"
	adminTwo := "host=10.190.24.252 port=5432 user=admin password=4IeggAix@2021 dbname=postgres sslmode=disable"

	bTestOne := "host=10.190.24.252 port=5432 user=aix-creative-project-backend-test password=4IeggAix@2021 dbname=creative_project sslmode=disable"
	bTestTwo := "host=10.190.24.252 port=5432 user=aix-creative-project-backend-test password=4IeggAix@2021 dbname=postgres sslmode=disable"

	dev := "host=1.13.102.61 port=5432 user=postgres password=pgsql0922 dbname=postgres sslmode=disable"

	var strSlice []string
	strSlice = append(strSlice, adminOne, adminTwo, bTestOne, bTestTwo, dev)

	for _, v := range strSlice {
		db, err := sql.Open("postgres", v)
		if err != nil {
			fmt.Println("db: ", v)
			fmt.Printf("error: %v \n\n", err.Error())
			continue
		}

		err = db.Ping()
		if err != nil {
			fmt.Println("db: ", v)
			fmt.Printf("error: %v \n\n", err.Error())
			continue
		}

		fmt.Println("db: ", v)
		fmt.Println("connect successfully!")
		db.Close()
	}

}
