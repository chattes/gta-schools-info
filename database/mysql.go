package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/chattes/gta-schools-info/common"
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	connectionString string
}

func NewMySql() *mysql {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	conn_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, pass, host, port)

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

	conn.Exec("USE schools")
	prepare_statement := fmt.Sprintf("select * from school_info")
	res, err := conn.Exec(prepare_statement)
	if err != nil {
		return nil, err
	}
	fmt.Print(res)

	return make([]common.School, 0), nil

}

func (p *mysql) ReadById(id int) (common.School, error) {

	return common.School{}, nil

}
