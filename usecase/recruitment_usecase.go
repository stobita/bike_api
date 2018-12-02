package usecase

import "github.com/stobita/bike_api/entity"

type RecruitmentInputPort interface {
	GetRecruitmentByUserID(userID int) error
}

type RecruitmentOutputPort interface {
	Render([]entity.Recruitment)
	RenderError(err error)
}

type RecruitmentRepository interface {
	GetRecruitmentByUserID(userID int) ([]entity.Recruitment, error)
}

type RecruitmentUsecase struct {
	OutputPort RecruitmentOutputPort
	Repository RecruitmentRepository
}

func (r *RecruitmentUsecase) GetRecruitmentByUserID(userID int) error {
	recruitment, err := r.Repository.GetRecruitmentByUserID(userID)
	if err != nil {
		r.OutputPort.RenderError(err)
	}
	r.OutputPort.Render(recruitment)
	return nil
}
