package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/chattes/gta-schools-info/common"
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	connectionString string
}

type schoolDb struct {
	Id           sql.NullInt64
	URL          sql.NullString
	Name         sql.NullString
	SchoolId     sql.NullString
	Type         sql.NullString
	IsCatholic   sql.NullBool
	Language     sql.NullString
	Level        sql.NullString
	City         sql.NullString
	CitySlug     sql.NullString
	Board        sql.NullString
	FraserRating sql.NullFloat64
	EQAORating   sql.NullFloat64
	Address      sql.NullString
	Grades       sql.NullString
	Website      sql.NullString
	PhoneNumber  sql.NullString
	Latitude     sql.NullFloat64
	Longitude    sql.NullFloat64
}

func NewMySql() *mysql {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	conn_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/schools", user, pass, host, port)

	return &mysql{
		connectionString: conn_string,
	}

}

func (p *mysql) Find(query string) ([]common.School, error) {

	conn, err := sql.Open("mysql", p.connectionString)

	if err != nil {
		fmt.Println("Error connecting to Database")
		return nil, err
	}

	defer conn.Close()

	prepare_statement := fmt.Sprintf("select * from school_info where id in (select id from school_info where match(name, address) against (\"%s\" in natural language mode))", query)
	fmt.Println(prepare_statement)
	res, err := conn.Query(prepare_statement)
	if err != nil {
		return nil, err
	}

	var allSchools = []common.School{}
	for res.Next() {

		var rdb schoolDb
		var rdbOut common.School
		err = res.Scan(
			&rdb.Id,
			&rdb.Name,
			&rdb.SchoolId,
			&rdb.Type,
			&rdb.IsCatholic,
			&rdb.Language,
			&rdb.Level,
			&rdb.City,
			&rdb.CitySlug,
			&rdb.Board,
			&rdb.FraserRating,
			&rdb.EQAORating,
			&rdb.Address,
			&rdb.Grades,
			&rdb.Website,
			&rdb.PhoneNumber,
			&rdb.Latitude,
			&rdb.Longitude,
		)
		if err != nil {
			return nil, err
		}

		if !rdb.Id.Valid {
			return nil, errors.New("not found")
		}

		rdbOut.Id = int(rdb.Id.Int64)

		if rdb.Address.Valid {
			rdbOut.Address = rdb.Address.String
		}

		if rdb.Board.Valid {
			rdbOut.Board = rdb.Board.String
		}

		if rdb.City.Valid {
			rdbOut.City = rdb.City.String
		}

		if rdb.CitySlug.Valid {
			rdbOut.CitySlug = rdb.CitySlug.String
		}

		if rdb.EQAORating.Valid {
			rdbOut.EQAORating = rdb.EQAORating.Float64
		}

		if rdb.FraserRating.Valid {
			rdbOut.FraserRating = rdb.FraserRating.Float64
		}

		if rdb.Grades.Valid {
			rdbOut.Grades = rdb.Grades.String
		}

		if rdb.IsCatholic.Valid {
			rdbOut.IsCatholic = rdb.IsCatholic.Bool
		}

		if rdb.Language.Valid {
			rdbOut.Language = rdb.Language.String
		}
		if rdb.Latitude.Valid {
			rdbOut.Latitude = rdb.Latitude.Float64
		}
		if rdb.Level.Valid {
			rdbOut.Level = rdb.Level.String
		}

		rdbOut.Longitude = rdb.Longitude.Float64
		rdbOut.Name = rdb.Name.String
		rdbOut.PhoneNumber = rdb.PhoneNumber.String
		rdbOut.SchoolId = rdb.SchoolId.String
		rdbOut.Type = rdb.Type.String
		rdbOut.URL = rdb.URL.String
		rdbOut.Website = rdb.Website.String

		allSchools = append(allSchools, rdbOut)

	}

	return allSchools, nil

}

func (p *mysql) ReadById(id int) (common.School, error) {

	return common.School{}, nil

}
