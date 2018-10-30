package model

import (
	"time"
)

// RecruitmentPrefecture database model
type RecruitmentPrefecture struct {
	ID            int64     `xorm:"id pk autoincr" json:"id"`
	RecruitmentID int64     `xorm:"recruitment_id"`
	PrefectureID  int64     `xorm:"prefecture_id"`
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`
}

// NewRecruitmentPrefecture constructor
func NewRecruitmentPrefecture() *RecruitmentPrefecture {
	return new(RecruitmentPrefecture)
}

// Create create new recruitment Prefecture
func (rp *RecruitmentPrefecture) Create(recruitmentID int64, prefectureID int64) error {
	rp.RecruitmentID = recruitmentID
	rp.PrefectureID = prefectureID
	_, err := engine.Insert(rp)
	if err != nil {
		return err
	}
	return nil
}

func (rp *RecruitmentPrefecture) BulkCreate(recruitmentID int64, prefectureIDs []int64) error {
	var recruitmentPrefectures []RecruitmentPrefecture
	for _, value := range prefectureIDs {
		item := RecruitmentPrefecture{RecruitmentID: recruitmentID, PrefectureID: value}
		recruitmentPrefectures = append(recruitmentPrefectures, item)
	}
	_, err := engine.Insert(recruitmentPrefectures)
	if err != nil {
		return err
	}
	return nil
}
