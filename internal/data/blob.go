package data

import "gitlab.com/distributed_lab/kit/pgdb"

type Blobs interface {
	New() Blobs
	Transaction(fn func(q Blobs) error) error
	Create(new *Blob) (string, error)
	Select() ([]Blob, error)
	Get() (*Blob, error)
	Update(id int64, newValue *Blob) error
	Delete(id int64) error
	BlobsPage(params *pgdb.OffsetPageParams) Blobs
	GetByID(id int64) Blobs
	GetByOwner(owner string) Blobs
}

type Blob struct {
	Id      int64  `json:"id" db:"id" structs:"-"`
	Owner   string `json:"owner" db:"owner" structs:"owner"`
	Content string `json:"content" db:"content" structs:"content"`
}
