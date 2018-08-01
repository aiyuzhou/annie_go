package models

import (
  "database/sql"
	"time"

	"annie_go/models/mymysql"

	"github.com/go-sql-driver/mysql"
)

// Role model definiton.
type Role struct {
	ID       int64     `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Password string    `json:"password,omitempty"`
	RegDate  time.Time `json:"reg_date,omitempty"`
}

// NewRole alloc and initialize a role.
func NewRole(f *RolePostForm, t time.Time) *Role {
	role := Role{
		ID:       f.ID,
		Name:     f.Name,
		Password: f.Password,
		RegDate:  t}

	return &role
}

func (r *Role) FindByID(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, name, password, reg_date FROM roles WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpID sql.NullInt64
	var tmpName sql.NullString
	var tmpPassword sql.NullString
	var tmpRegDate mysql.NullTime
	if err := row.Scan(&tmpID, &tmpName, &tmpPassword, &tmpRegDate); err != nil {
		// Not found.
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		}

		return ErrDatabase, err
	}

	if tmpID.Valid {
		r.ID = tmpID.Int64
	}
	if tmpName.Valid {
		r.Name = tmpName.String
	}
	if tmpPassword.Valid {
		r.Password = tmpPassword.String
	}
	if tmpRegDate.Valid {
		r.RegDate = tmpRegDate.Time
	}

	return 0, nil
}

func (r *Role) ClearPass() {
	r.Password = ""
}

func (r *Role) Insert() (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("INSERT INTO roles(id, name, password, reg_date) VALUES(?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(r.ID, r.Name, r.Password, r.RegDate); err != nil {
		if e, ok := err.(*mysql.MySQLError); ok {
			//Duplicate key
			if e.Number == 1062 {
				return ErrDupRows, err
			}

			return ErrDatabase, err
		}

		return ErrDatabase, err
	}

	//r.ID, _ = result.LastInsertId()

	return 0, nil
}

func (r *Role) UpdateByID(id int64, f *RolePutForm) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("UPDATE roles SET name = ?, password = ? WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	result, err := st.Exec(f.Name, f.Password, id)
	if err != nil {
		return ErrDatabase, err
	}

	num, _ := result.RowsAffected()
	if num > 0 {
		return 0, nil
	}

	return ErrNotFound, nil
}

func (r *Role) DeleteByID(id int64) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("DELETE FROM roles WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	result, err := st.Exec(id)
	if err != nil {
		return ErrDatabase, err
	}

	num, _ := result.RowsAffected()
	if num > 0 {
		return 0, nil
	}

	return ErrNotFound, nil
}
