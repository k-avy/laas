package license

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgconn"
)

type PostgreSQLRepository struct{
	db *sql.DB
}

func NewPostgreSQLRepository( db *sql.DB) *PostgreSQLRepository {
	return &PostgreSQLRepository{
        db: db,
    }
}

func (r* PostgreSQLRepository) Migrate(ctx context.Context) error {
	query :=`
	CREATE TABLE IF NOT EXISTS licenses(
		id SERIAL PRIMARY KEY,
		shortname TEXT NOT NULL UNIQUE,
		longname TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL UNIQUE
	 );
	`

	 _,err := r.db.ExecContext(ctx, query)
	 return err
}

func (r* PostgreSQLRepository) Create(ctx context.Context, license License) (*License, error) {
	var id int64
	err:= r.db.QueryRowContext(ctx, "INSERT INTO license(shortname, longname, url) values($1, $2, $3) RETURNING id", license.Shortname,license.Longname, license.Url).Scan(&id)
	if err!= nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
			return nil, ErrCheck
			}
		}
		return  nil, err
	}
	license.ID=id

	return &license,nil
}

func (r *PostgreSQLRepository) All(ctx context.Context) ([]License, error) {
	rows,err:= r.db.QueryContext(ctx, "SELECT * FROM licenses")
	if err!= nil {
		return nil,err
	}
	defer rows.Close()

	var all []License
	for rows.Next() {
		var license License
		if err:= rows.Scan(&license.ID, &license.Shortname, &license.Longname, &license.Url) ; err!=nil{
			return nil, err
		} 
		all= append(all, license)
	}
	return all,nil
}
func (r*PostgreSQLRepository) GetByName(ctx context.Context, shortname string) (*License,error){
	row:= r.db.QueryRowContext(ctx, "SELECT * FROM licenses WHERE shortname=$1",shortname)
	var license License
	if err:=row.Scan(&license.ID,&license.Shortname,&license.Longname,&license.Url); err!=nil {
		if errors.Is(err,sql.ErrNoRows){
			return nil,ErrNotExist
		}
		return nil, err
	}
	return &license, nil
}

func (r *PostgreSQLRepository) Update(ctx context.Context, id, int64, updated License) (*License, error){
	res, err:= r.db.ExecContext(ctx,"UPDATE licenses SET shortname = $1, longname=$2, url=$3 WHERE id=$4",updated.Shortname,updated.Longname, updated.Url, id)
	if err !=nil {
		var pgxError *pgconn.PgError
		if errors.As(err,& pgxError){
			if pgxError.Code=="23505" {
				return nil, ErrCheck
			}
		}
		return nil, err
	}
	 
	rowsAffected, err:=res.RowsAffected()
	if err!=nil{
		return nil,err
	}
	if rowsAffected==0 {
		return nil,ErrUpdateFailed
	}
	return &updated,nil

}

func (r *PostgreSQLRepository) Delete(ctx context.Context, id int64) error{
	res, err := r.db.ExecContext(ctx, "DELETE FROM licenses WHERE id= $1", id)
	if err!= nil{
		return err
	}
	rowsAffected, err:= res.RowsAffected()
	if err!= nil {
		return err
	}
	if rowsAffected==0 {
		return ErrDeleteFailed
	}
	return err
}
