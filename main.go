package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/stobita/bike_api/adapter/gateway"
	"github.com/stobita/bike_api/adapter/presenter"
	"github.com/stobita/bike_api/controller"
	"github.com/stobita/bike_api/driver"
	"github.com/stobita/bike_api/usecase"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		recruitmentInteractor := &usecase.RecruitmentUsecase{}
		recruitmentInteractor.RecruitmentRepo = &gateway.RecruitmentRepository{
			SqlHandler: driver.NewDBConn(),
		}
		recruitmentInteractor.OutputPort = &presenter.RecruitmentPresenter{
			Writer: w,
		}
		handler := controller.RecruitmentController{
			InputFactory: func(o usecase.RecruitmentOutputPort) usecase.RecruitmentInputPort {
				return recruitmentInteractor

			},
			OutputFactory: func(w http.ResponseWriter) usecase.RecruitmentOutputPort {
				return recruitmentInteractor.OutputPort
			},
		}
		handler.GetRecruitmentByUserID(w, r)
	})
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, r)
}
