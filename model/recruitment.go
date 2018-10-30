package model

import (
	"log"
	"time"

	"github.com/go-xorm/xorm"
)

// Recruitment database model
type Recruitment struct {
	ID                int64     `xorm:"id pk autoincr" json:"id"`
	UserID            int64     `xorm:"user_id" json:"userId"`
	Title             string    `xorm:"title" json:"title"`
	Content           string    `xorm:"content" json:"content"`
	Image             string    `xorm:"image" json:"image"`
	MinEngineCapacity int       `xorm:"min_engine_capacity" json:"minEngineCapacity"`
	MaxEngineCapacity int       `xorm:"max_engine_capacity" json:"maxEngineCapacity"`
	MinAge            int       `xorm:"min_age" json:"minAge"`
	MaxAge            int       `xorm:"max_age" json:"maxAge"`
	CreatedAt         time.Time `xorm:"created" json:"createdAt"`
	UpdatedAt         time.Time `xorm:"updated" json:"updatedAt"`
}

// RecruitmentRelation recruitment join user
type RecruitmentRelation struct {
	Recruitment           `xorm:"extends"`
	User                  `xorm:"extends"`
	RecruitmentComment    `xorm:"extends"`
	RecruitmentPrefecture `xorm:"extends"`
	Prefecture            `xorm:"extends"`
}

// RecruitmentRelationModel recruitment relation model
type RecruitmentRelationModel struct {
	Recruitment
	User
	Comments      []RecruitmentComment
	PrefectureIDs []int64
}

// TableName for relation table name
func (rr *RecruitmentRelation) TableName() string {
	return "recruitment"
}

// Poster recruitment create user
type Poster struct {
	Name  string
	Image string
}

// func (p *Poster) TableName() string {
// 	return "user"
// }

// NewRecruitment constructot of recruitment
func NewRecruitment() *Recruitment {
	return new(Recruitment)
}

// Create create new recruitment
func (r *Recruitment) Create(
	userID int64,
	title string,
	content string,
	image string,
	minEngineCapacity int,
	maxEngineCapacity int,
	minAge int,
	maxAge int,
	prefectureIDs []int64,
) error {
	r.UserID = userID
	r.Title = title
	r.Content = content
	r.Image = image
	r.MinEngineCapacity = minEngineCapacity
	r.MaxEngineCapacity = maxEngineCapacity
	r.MinAge = minAge
	r.MaxAge = maxAge
	_, err := engine.Insert(r)
	if err != nil {
		return err
	}
	err = NewRecruitmentPrefecture().BulkCreate(r.ID, prefectureIDs)
	if err != nil {
		return err
	}
	return nil
}

func RecruitmentGeneralRelation() *xorm.Session {
	return engine.
		Join("INNER", "user", "user.id = recruitment.user_id").
		Join("LEFT OUTER", "recruitment_prefecture", "recruitment.id = recruitment_prefecture.recruitment_id").
		Join("LEFT OUTER", "recruitment_comment", "recruitment.id = recruitment_comment.recruitment_id").
		Join("LEFT OUTER", "prefecture", "prefecture.id = recruitment_prefecture.prefecture_id")
}

func RecruitmentGeneralRelationByUser(userID int64) *xorm.Session {
	return engine.
		Join("INNER", "user", "user.id = recruitment.user_id").
		Join("LEFT OUTER", "recruitment_prefecture", "recruitment.id = recruitment_prefecture.recruitment_id").
		Join("LEFT OUTER", "recruitment_comment", "recruitment.id = recruitment_comment.recruitment_id").
		Join("LEFT OUTER", "prefecture", "prefecture.id = recruitment_prefecture.prefecture_id").
		Where("recruitment.user_id = ?", userID)
}

// GetAllWithRelation get all recruitments with user
func (r Recruitment) GetAllWithRelation() *[]RecruitmentRelationModel {
	var recruitments []RecruitmentRelation
	RecruitmentGeneralRelation().Find(&recruitments)
	return r.Shape(recruitments)
}

// GetUserRecruitment get all recruitments by userid
func (r Recruitment) GetUserRecruitment(userID int64) *[]RecruitmentRelationModel {
	var recruitments []RecruitmentRelation
	RecruitmentGeneralRelationByUser(userID).Find(&recruitments)
	log.Println(recruitments)
	return r.Shape(recruitments)
}

// Shape set relation data
func (r Recruitment) Shape(list []RecruitmentRelation) *[]RecruitmentRelationModel {
	var recruitments []RecruitmentRelationModel
	for _, value := range list {
		if r.Mash(recruitments, value) {
			recruitments = append(recruitments, RecruitmentRelationModel{
				Recruitment:   value.Recruitment,
				User:          value.User,
				PrefectureIDs: []int64{value.Prefecture.ID}})
		}
	}
	return &recruitments
}

// Mash set relation model
func (r Recruitment) Mash(recruitments []RecruitmentRelationModel, value RecruitmentRelation) (isNew bool) {
	for i, r := range recruitments {
		if r.Recruitment.ID == value.Recruitment.ID {
			ids := &recruitments[i].PrefectureIDs
			*ids = append(*ids, value.Prefecture.ID)
			comments := &recruitments[i].Comments
			*comments = append(*comments, value.RecruitmentComment)
			return false
		}
	}
	return true
}
