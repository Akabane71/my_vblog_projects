package mysql

import (
	"bluebell/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (community []*models.Community, err error) {
	sqlStr := `select community_id, community_name from community`
	err = db.Select(&community, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no community in db")
		err = nil
	}
	return community, err
}

func GetCommunityDetailByID(id int64) (*models.Community, error) {
	community := new(models.Community)
	sqlStr := `select community_id,community_name,introduction,create_time from community where community_id = ?`
	err := db.Get(community, sqlStr, id)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no community in db")
		err = ErrorInvalidID
	}
	return community, err
}
