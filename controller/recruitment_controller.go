package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/stobita/bike_api/usecase"
)

type Context interface {
	Keys(string) int
}

type RecruitmentController struct {
	InputFactory
	OutputFactory
}

type InputFactory func(o usecase.RecruitmentOutputPort) usecase.RecruitmentInputPort

type OutputFactory func(w http.ResponseWriter) usecase.RecruitmentOutputPort

func NewRecruitmentController(inputFactory InputFactory, outputFactory OutputFactory) *RecruitmentController {
	return &RecruitmentController{
		InputFactory:  inputFactory,
		OutputFactory: outputFactory,
	}
}

func (rc *RecruitmentController) GetRecruitmentByUserID(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	//userID, ok := ctx.Value("userID").(int)
	// if !ok {
	// 	fmt.Fprint(w, "error")
	// }
	query := r.URL.Query()
	userID, _ := strconv.Atoi(query["userID"][0])

	outputPort := rc.OutputFactory(w)
	inputPort := rc.InputFactory(outputPort)
	err := inputPort.GetRecruitmentByUserID(userID)
	if err != nil {
		fmt.Fprint(w, "error")
	}
}
