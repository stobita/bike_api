package gateway

import (
	"log"

	"github.com/stobita/bike_api/entity"
)

type RecruitmentRepository struct {
	SqlHandler
}

func (rr *RecruitmentRepository) GetRecruitmentByUserID(userID int) ([]entity.Recruitment, error) {
	rows, err := rr.Query("SELECT title FROM recruitment WHERE user_id = ?", userID)
	if err != nil {
		log.Fatalln(err)
	}
	var recruitments []entity.Recruitment
	for rows.Next() {
		recruitment := entity.Recruitment{}
		err := rows.StructScan(&recruitment)
		if err != nil {
			log.Fatalln(err)

		}
		recruitments = append(recruitments, recruitment)
	}
	return recruitments, nil
}
