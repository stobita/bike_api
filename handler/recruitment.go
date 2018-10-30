package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stobita/bike_api/lib"
	"github.com/stobita/bike_api/model"
)

// RecruitmentJSON create params
type RecruitmentJSON struct {
	Title             string  `json:"title" binding:"required"`
	Content           string  `json:"content" binding:"required"`
	Image             string  `json:"image"`
	MinEngineCapacity int     `json:"minEngineCapacity"`
	MaxEngineCapacity int     `json:"maxEngineCapacity"`
	MinAge            int     `json:"minAge"`
	MaxAge            int     `json:"maxAge"`
	PrefectureID      []int64 `json:"prefectureId"`
}

// RecruitmentCommentJSON create params
type RecruitmentCommentJSON struct {
	RecruitmentID int64  `json:"recruitmentID" binding:"required"`
	Content       string `json:"content" binding:"required"`
}

// RecruitmentResponseJSON response params
type RecruitmentResponseJSON struct {
	ID                int64   `json:"id"`
	Title             string  `json:"title"`
	Content           string  `json:"content"`
	Image             string  `json:"image"`
	MinEngineCapacity int     `json:"minEngineCapacity"`
	MaxEngineCapacity int     `json:"maxEngineCapacity"`
	MinAge            int     `json:"minAge"`
	MaxAge            int     `json:"maxAge"`
	PrefectureID      []int64 `json:"prefectureId"`
	UserName          string  `json:"userName"`
	UserImage         string  `json:"userImage"`
}

// RecruitmentCommentResponseJSON response params
type RecruitmentCommentResponseJSON struct {
	ID            int64  `json:"id"`
	RecruitmentID int64  `json:"recruitmentID"`
	Content       string `json:"content"`
	IsOwn         bool   `json:"isOwn"`
}

// PostRecruitment create recruitment
func PostRecruitment(c *gin.Context) {
	var json RecruitmentJSON
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, err.Error())
		return
	}

	var filePath string
	if json.Image != "" {
		filePath, _ = lib.ImageUpload(json.Image)
	}

	err := model.NewRecruitment().Create(
		c.Keys["userId"].(int64),
		json.Title,
		json.Content,
		filePath,
		json.MinEngineCapacity,
		json.MaxEngineCapacity,
		json.MinAge,
		json.MaxAge,
		json.PrefectureID,
	)

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.AbortWithStatus(200)
}

// GetRecruitment get recruitment
func GetRecruitment(c *gin.Context) {
	recruitments := *model.Recruitment{}.GetAllWithRelation()
	var result []RecruitmentResponseJSON
	for _, value := range recruitments {
		recruitment := RecruitmentResponseJSON{
			ID:                value.Recruitment.ID,
			Title:             value.Recruitment.Title,
			Content:           value.Recruitment.Content,
			Image:             value.Recruitment.Image,
			MinEngineCapacity: value.Recruitment.MinEngineCapacity,
			MaxEngineCapacity: value.Recruitment.MaxEngineCapacity,
			MinAge:            value.Recruitment.MinAge,
			MaxAge:            value.Recruitment.MaxAge,
			UserName:          value.User.Name,
			UserImage:         value.User.Image,
			PrefectureID:      value.PrefectureIDs,
		}
		result = append(result, recruitment)
	}
	c.JSON(200, result)
}

// PostRecruitmentComment post comment to recruitment
func PostRecruitmentComment(c *gin.Context) {
	var json RecruitmentCommentJSON
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, lib.ErrorResponse(err.Error()))
		return
	}
	err := model.NewRecruitmentComment().Create(
		json.RecruitmentID,
		c.Keys["userId"].(int64),
		json.Content,
	)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.AbortWithStatus(200)
}

// GetRecruitmentComments get comment to recruitment
func GetRecruitmentComments(c *gin.Context) {
	var result []RecruitmentCommentResponseJSON
	recruitmentID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	recruitmentComments := *model.NewRecruitmentComment().GetByRecruitmentID(recruitmentID)
	for _, value := range recruitmentComments {
		recruitmentComment := RecruitmentCommentResponseJSON{
			ID:            value.ID,
			RecruitmentID: value.RecruitmentID,
			Content:       value.Content,
			IsOwn:         value.UserID == c.Keys["userId"].(int64),
		}
		result = append(result, recruitmentComment)
	}
	c.JSON(200, result)
}

func GetUserRecruitmentPost(c *gin.Context) {
	recruitments := *model.NewRecruitment().GetUserRecruitment(c.Keys["userId"].(int64))
	c.JSON(200, recruitments)
}
