package controller

import (
	"fmt"

	"github.com/chattes/gta-schools-info/common"
	"github.com/chattes/gta-schools-info/database"
)

type SchoolController struct {
	Db database.Persistence
}

func (c *SchoolController) Search(name string) (resp []common.School, err error) {

	fmt.Println("In Search Controller")

	if c.Db == nil {
		panic("DB not setup")
	}

	school, err := c.Db.Find(name)
	if err != nil {
		return nil, err
	}
	return school, nil
}
func (c *SchoolController) ReadById(id int) (resp common.School, err error) {

	fmt.Println("In Read By Id Controller")
	if c.Db == nil {
		panic("DB not setup")
	}
	school, err := c.Db.ReadById(id)
	if err != nil {
		return school, err
	}
	return school, nil
}
