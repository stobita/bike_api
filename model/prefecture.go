package model

// Prefecture Prefecture db model
type Prefecture struct {
	ID       int64  `xorm:"id pk autoincr" json:"id"`
	Name     string `xorm:"name" json:"name"`
	NameKana string `xorm:"name_kana" json:"nameKana"`
}

func NewPrefecture() *Prefecture {
	return new(Prefecture)
}

// GetAll get all prefectures
func (p Prefecture) GetAll() *[]Prefecture {
	var prefectures []Prefecture
	engine.Find(&prefectures)
	return &prefectures
}
