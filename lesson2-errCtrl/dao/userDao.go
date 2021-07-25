package dao

import (
	"database/sql"
	def "errCtrl/errors"
	"errCtrl/model"
	"github.com/pkg/errors"
)

func (d *Dao) GetUser(name string) (*model.User, error) {
	var user model.User
	err := d.engine.QueryRow("SELECT * FROM user  WHERE `name` = ?", name).Scan(&user.Id, &user.Name, &user.Age)
	if err != nil && err == sql.ErrNoRows {
		// notFound
		msg := "query not found,with name = " + name
		return nil, errors.Wrap(def.DataNotFoundErr, msg)
	}
	if err != nil {
		return nil, errors.Wrap(err, "query  from db err")
	}

	return &user, nil
}
