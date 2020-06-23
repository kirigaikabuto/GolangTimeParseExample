package postgres

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"log"
	_ "github.com/lib/pq"
)
var Queries = []string{
	`CREATE TABLE IF NOT EXISTS data(
		id text,
		name text,
		date_of_birth date,
		PRIMARY KEY(id)
	);`,
}
type postgreStore struct {
	db *sql.DB
}

func NewPostgreStore(cfg Config) (DataRepo,error){
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range Queries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	return &postgreStore{db: db},err
}

func (ps *postgreStore) Create (obj *Data) (*Data,error){
	obj.Id = uuid.New().String()
	result, err := ps.db.Exec("insert into data (id,name,date_of_birth) values ($1,$2,$3)",obj.Id,obj.Name,obj.DateOfBirth.Time)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, errors.New("cannot create object")
	}
	return obj,nil
}
func (ps *postgreStore) Get(id string) (*Data,error) {
	return nil,nil
}
