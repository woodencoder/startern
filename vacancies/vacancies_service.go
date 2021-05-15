package vacancies

import (
	"context"
	"fmt"
	"startern/generated"
)

func (v VacanciesService) GetVacancies(ctx context.Context, request *generated.GetVacanciesRequest) (*generated.GetVacanciesResponse, error) {
	var results = []*generated.Vacancy{}
	for i := 0; i < 1000; i++{
		vacancy := generated.Vacancy{}
		vacancy.Id = fmt.Sprintf("%d", i)

		vacancyInfo := generated.VacancyInfo{}
		vacancyInfo.Title = "Job Hunter"
		vacancy.Info = &vacancyInfo
		results = append(results, &vacancy)
	}
	res := generated.GetVacanciesResponse{}
	res.Vacancies = results
	return &res, nil
}

func (v VacanciesService) mustEmbedUnimplementedVacanciesServer() {
	panic("implement me")
}

type VacanciesService struct {

}