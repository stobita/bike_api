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

	recruitmentController := controller.RecruitmentController{
		InputFactory: func(o usecase.RecruitmentOutputPort) usecase.RecruitmentInputPort {
			interactor := &usecase.RecruitmentUsecase{
				Repository: &gateway.RecruitmentRepository{SqlHandler: driver.NewDBConn()},
				OutputPort: o,
			}
			return interactor
		},
		OutputFactory: func(w http.ResponseWriter) usecase.RecruitmentOutputPort {
			return &presenter.RecruitmentPresenter{Writer: w}
		},
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		recruitmentController.GetRecruitmentByUserID(w, r)
	})

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, r)
}
