package pg

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/blobs/internal/data"
)

type BlobQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

const (
	blobsTable = "blob"
)

var blobsSelect = sq.Select("*").From(blobsTable)

func NewBlobs(q *pgdb.DB) data.Blobs {
	return &BlobQ{
		db:  q.Clone(),
		sql: blobsSelect,
	}
}

func (q *BlobQ) New() data.Blobs {
	return NewBlobs(q.db)
}

func (q *BlobQ) Transaction(fn func(q data.Blobs) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *BlobQ) Create(blob *data.Blob) (string, error) {
	clauses := structs.Map(blob)

	var id string
	stmt := sq.Insert(blobsTable).SetMap(clauses).Suffix("returning id")
	err := q.db.Get(&id, stmt)
	return id, err
}

func (q *BlobQ) Select() ([]data.Blob, error) {
	var result []data.Blob

	err := q.db.Select(&result, q.sql)

	return result, err
}

func (q *BlobQ) Get() (*data.Blob, error) {
	var result data.Blob
	err := q.db.Get(&result, q.sql)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}
	return &result, err
}

func (q *BlobQ) Update(id int64, newValue *data.Blob) error {
	clauses := structs.Map(newValue)

	stmt := sq.Update(blobsTable).Where(sq.Eq{
		"id": id,
	}).SetMap(clauses)

	return q.db.Exec(stmt)
}

func (q *BlobQ) Delete(id int64) error {
	stmt := sq.Delete(blobsTable).Where(sq.Eq{"id": id})
	return q.db.Exec(stmt)
}

func (q *BlobQ) BlobsPage(params *data.PageParams) data.Blobs {
	q.sql = params.ApplyTo(q.sql, "id")
	return q
}

func (q *BlobQ) GetByID(id int64) data.Blobs {
	q.sql = q.sql.Where(sq.Eq{"id": id})
	return q
}
