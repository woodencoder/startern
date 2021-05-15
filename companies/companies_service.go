package companies

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	"startern/generated"
	"os"
)
import "context"

type CompaniesService struct {
}

func (c *CompaniesService) GetCompanies(ctx context.Context, request *generated.GetCompaniesRequest) (*generated.GetCompaniesResponse, error) {
	var results = []*generated.Company{}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close(context.Background())

	//res := conn.QueryRow(context.Background(), "select id, title, sourceUrl from companies")
	//res.Scan().
	res, err := conn.Query(context.Background(), "select id, title from companies")

	for res.Next() {
		var id string
		var title string
		//var sourceUrl string

		res.Scan(id, title)

		company := generated.Company{Id: id, Name: title}
		results = append(results, &company)
	}

	response := generated.GetCompaniesResponse{}
	response.Companies = results
	return &response, nil
}

func (s *CompaniesService) AddCompany(ctx context.Context, company generated.Company) error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}
	defer conn.Close(context.Background())

	//res := conn.QueryRow(context.Background(), "select id, title, sourceUrl from companies")
	//res.Scan().
	_, err = conn.Query(context.Background(), "INSERT INTO companies (title) VALUES ($1)", company.Name)

	if err != nil {
		println("SUCCESSFULLY")
	} else {
		println("SUCCESSFULLY ADDED COMPANY: %s", company.Name)
	}

	return err
}

func (s *CompaniesService) ConnectToDB(ctx context.Context) {
	//conn.ExecContext(ctx, "SELECT ")
}

func (c CompaniesService) mustEmbedUnimplementedCompaniesServer() {
	panic("implement me")
}