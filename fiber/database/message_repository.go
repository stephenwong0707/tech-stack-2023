package database

import (
	"store/entity"

	"github.com/scylladb/gocqlx/v2"
)

type MessageRepository struct {
	session gocqlx.Session
}

func (r *MessageRepository) Init() {
	newSession, err := gocqlx.WrapSession(cluster().CreateSession())
	if err != nil {
		panic(err)
	}
	r.session = newSession
}

func (r *MessageRepository) Insert(message *entity.Message) error {
	q := r.session.Query(messageTable.Insert()).BindStruct(message)
	return q.ExecRelease()
}
