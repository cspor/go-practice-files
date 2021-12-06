package row

import (
	"github.com/cspor/go-practice-files/services/errorHandler"
	"github.com/google/uuid"
)

type Row struct {
	Id string `json:"id"`
}

func (row *Row) GenerateId() *Row {
	id, err := uuid.NewRandom()
	errorHandler.Check(err)

	row.Id = id.String()

	return row
}

func NewRow() *Row {
	row := new(Row)
	row.GenerateId()
	return row
}
