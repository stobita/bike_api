package presenter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stobita/bike_api/entity"
)

type RecruitmentPresenter struct {
	Writer http.ResponseWriter
}

func (rp *RecruitmentPresenter) Render(recruitments []entity.Recruitment) {
	log.Printf("recruitments:%v", recruitments)

	for _, value := range recruitments {
		fmt.Fprint(rp.Writer, value.Title)
	}
}

func (rp *RecruitmentPresenter) RenderError(err error) {
	fmt.Fprint(rp.Writer, err)
}
