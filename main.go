package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

type Cover struct {
	Id   int
	Name string
}

var db *sql.DB

// server ต้นทางปิดไปแล้ว ให้หา server ใหม่มาเล่น หากต้องการ run ผ่าน
func main() {
	var err error
	db, err = sql.Open("sqlserver", "sqlserver://sa:P@ssw0rd@13.76.163.13/?database=techcoach")
	// db, err = sql.Open("mysql", "root:P@ssw0rd@tcp(13.76.163.13)/techcoach")
	if err != nil {
		panic(err)
	}
	//TEST RUN func AddCover
	// cover := Cover{44, "Fe"}
	// err = AddCover(cover)
	// if err != nil {
	// 	panic(err)
	// }

	// TEST RUN func UpdateCover
	// cover := Cover{44, "Akkawat"}
	// err = UpdateCover(cover)
	// if err != nil{
	// 	panic(err)
	// }

	//TEST RUN func DeleteCover
	err = DeleteCover(44)
	if err != nil {
		panic(err)
	}

	covers, err := GetCovers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, cover := range covers {
		fmt.Println(cover)
	}

	//TEST RUN GetCover
	// cover, err := GetCover(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(cover)
}

// แทนที่จะ rows.Next() ทีละแถว ใช้ for ทีเดียว หากถึงแถวสุดท้ายไม่สามารถสแกนผลออกมาได้ โปรแกรมก็จะหยุด
// id := 0
// name := ""
// ok := rows.Next()
//
//	if ok {
//		rows.Scan(&id, &name)
//		rows.Close()
//	}
func GetCovers() ([]Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	//ไม่ควรใช้ select * เพราะวันใดที่ database ไม่ได้มีแค่ id, name จะใช้ไม่ได้ จึงกำหนดให้ชัดเจนไปเลย
	query := "select id, name from cover"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	covers := []Cover{}
	for rows.Next() {
		cover := Cover{}
		err = rows.Scan(&cover.Id, &cover.Name)
		if err != nil {
			return nil, err
		}
		covers = append(covers, cover)

	}
	return covers, nil
}

func GetCover(id int) (*Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	//MSSQL
	// query := "select id, name from cover where id=@id"
	// row := db.QueryRow(query, sql.Named("id", id))

	//MYSQL
	query := "select id, name from cover where id=?"
	row := db.QueryRow(query, id)

	cover := Cover{}
	err = row.Scan(&cover.Id, &cover.Name)
	if err != nil {
		return nil, err
	}
	return &cover, nil
}
func AddCover(cover Cover) error {
	query := "insert into cover (id, name) value (?, ?)"
	result, err := db.Exec(query, cover.Id, cover.Name)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("cannot insert")
	}

	return nil

}
func UpdateCover(cover Cover) error {
	query := "update cover set name=? where id=?"
	result, err := db.Exec(query, cover.Name, cover.Id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("cannot update")
	}

	return nil

}
func DeleteCover(id int) error {
	query := "delete from cover where id=?"
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("cannot delete")
	}

	return nil

}
