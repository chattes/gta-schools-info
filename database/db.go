package database

import "github.com/chattes/gta-schools-info/common"

// Delete , Insert etc for future

type Persistence interface {
	Find(query string) ([]common.School, error)
	ReadById(id int) (common.School, error)
}
