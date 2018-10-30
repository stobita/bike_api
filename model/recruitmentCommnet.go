package model

import "time"

// RecruitmentComment database model
type RecruitmentComment struct {
	ID            int64     `xorm:"id pk" json:"id"`
	RecruitmentID int64     `xorm:"recruitment_id"`
	UserID        int64     `xorm:"user_id"`
	Content       string    `xorm:"content"`
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`
}

// NewRecruitmentComment constructor for RecruitmentComment
func NewRecruitmentComment() *RecruitmentComment {
	return new(RecruitmentComment)
}

// Create create new recruitment comment
func (rc RecruitmentComment) Create(recruitmentID int64, userID int64, content string) error {
	rc.RecruitmentID = recruitmentID
	rc.UserID = userID
	rc.Content = content
	_, err := engine.Insert(&rc)
	if err != nil {
		return err
	}
	return nil
}

// GetByRecruitmentID get comments by RecruitmentID
func (rc RecruitmentComment) GetByRecruitmentID(recruitmentID int64) *[]RecruitmentComment {
	var recruitmentComments []RecruitmentComment
	engine.Find(&recruitmentComments, &RecruitmentComment{RecruitmentID: recruitmentID})
	return &recruitmentComments
}
