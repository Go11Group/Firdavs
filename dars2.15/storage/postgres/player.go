package postgres

import (
	"database/sql"
	"github.com/Go11Group/at_lesson/lesson35/model"
)

type PlayerRepo struct {
	Db *sql.DB
}

func (p *PlayerRepo) Create(player model.Player) error {
	_, err := p.Db.Exec("insert into player (id, name, number, birthday) values ($1, $2,$3,$4)",
		player.Id, player.Name, player.Number, player.Birthday)
	if err != nil {
		return err
	}

	return nil
}
