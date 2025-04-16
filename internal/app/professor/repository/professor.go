package repository

import (
	"context"
	"log"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	apperr "github.com/devanfer02/ratemyubprof/pkg/http/errors"
	"github.com/devanfer02/ratemyubprof/pkg/util"
	"github.com/doug-martin/goqu/v9"
)

func (p *professorRepositoryImplPostgre) FetchAllProfessors(ctx context.Context, params *dto.FetchProfessorParam, pageQuery *dto.PaginationQuery) ([]entity.Professor, error) {
	var professors []entity.Professor

	goqu.Func("/", goqu.I("sum_diff_rate"), goqu.I("review_count")).As("diff_rating")
	qb := goqu.
		Select(
			goqu.I("p.id"), 
			goqu.I("p.name"),
			goqu.I("p.faculty"),
			goqu.I("p.major"),
			goqu.I("p.profile_img_link"),
			goqu.COUNT(goqu.I("r.id")).As("reviews_count"),
			goqu.Func("AVG", goqu.Func("COALESCE", goqu.I("r.difficulty_rating"), goqu.V(0))).As("avg_diff_rate"),
			goqu.Func("AVG", goqu.Func("COALESCE", goqu.I("r.friendly_rating"), goqu.V(0))).As("avg_friendly_rate"),
		).
		GroupBy(goqu.I("p.id")).
		From(goqu.T(professorTableName).As("p")).
		LeftJoin(
			goqu.T(reviewTableName).As("r"),
			goqu.On(goqu.I("r.prof_id").Eq(goqu.I("p.id"))),
		).
		Order(goqu.I("name").Asc()).
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	if pageQuery.Page != 0 && pageQuery.Limit != 0{
		qb = qb.Offset((pageQuery.Page - 1) * pageQuery.Limit).Limit(pageQuery.Limit)
	}

	qb = util.AddParamsToFetchProf(qb, params)

	
	query, args, err := qb.ToSQL()
	if err != nil {
		return nil, apperr.NewFromError(err, "Failed to fetch all professors").SetLocation()
	}
	
	log.Println("qb", query)
	query = p.conn.Rebind(query)
	
	rows, err := p.conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, apperr.NewFromError(err, "Failed to fetch all professors").SetLocation()
	}

	for rows.Next() {
		var professor entity.Professor
		if err := rows.StructScan(&professor); err != nil {
			return nil, apperr.NewFromError(err, "Failed to fetch all professors").SetLocation()
		}
		professors = append(professors, professor)
	}

	return professors, nil
}

func (p *professorRepositoryImplPostgre) FetchProfessorByID(ctx context.Context, id string) (entity.Professor, error) {
	var professor entity.Professor
	qb := goqu.
		Select(
			goqu.I("p.id"), 
			goqu.I("name"),
			goqu.I("faculty"),
			goqu.I("major"),
			goqu.I("profile_img_link"),
			goqu.COUNT(goqu.I("r.id")).As("reviews_count"),
			goqu.Func("AVG", goqu.Func("COALESCE", goqu.I("r.difficulty_rating"), goqu.V(0))).As("avg_diff_rate"),
			goqu.Func("AVG", goqu.Func("COALESCE", goqu.I("r.friendly_rating"), goqu.V(0))).As("avg_friendly_rate"),
		).
		GroupBy(goqu.I("p.id")).
		From(goqu.T(professorTableName).As("p")).
		LeftJoin(
			goqu.T(reviewTableName).As("r"),
			goqu.On(goqu.I("r.prof_id").Eq(goqu.I("p.id"))),
		).
		Where(goqu.I("p.id").Eq(id)).
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	query, args, err := qb.ToSQL()
	if err != nil {
		return professor, apperr.NewFromError(err, "Failed to fetch professor by id").SetLocation()
	}

	query = p.conn.Rebind(query)
	
	err = p.conn.QueryRowxContext(ctx, query, args...).StructScan(&professor)
	if err != nil {
		return professor, apperr.NewFromError(err, "Failed to fetch professor by id").SetLocation()
	}

	return professor, nil
}

func (p *professorRepositoryImplPostgre) GetProfessorItems(ctx context.Context, params *dto.FetchProfessorParam) (uint64, error) {
	var count uint64

	qb := goqu.
		From(professorTableName).
		Select(goqu.COUNT(goqu.Star())).
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	qb = util.AddParamsToFetchProf(qb, params)

	query, args, err := qb.ToSQL()
	if err != nil {
		return 0, apperr.NewFromError(err, "Failed to get professor counter").SetLocation()
	}

	query = p.conn.Rebind(query)

	err = p.conn.QueryRowxContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, apperr.NewFromError(err, "Failed to get professor counter").SetLocation()
	}

	return count, nil
}

func (p *professorRepositoryImplPostgre) InsertProfessorsBulk(ctx context.Context, professors []entity.Professor) error {
	records := make([]goqu.Record, len(professors))
	for i, d := range professors {
		records[i] = goqu.Record{
			"id":               d.ID,
			"name":             d.Name,
			"faculty":          d.Faculty,
			"major":            d.Major,
			"profile_img_link": d.ProfileImgLink,
			"created_at":       time.Now(),
			"updated_at":       time.Now(),
		}
	}

	qb := goqu.
		Insert(professorTableName).
		Rows(records).
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	query, args, err := qb.ToSQL()
	if err != nil {
		return apperr.NewFromError(err, "Failed to insert professor bulk").SetLocation()
	}

	query = p.conn.Rebind(query)

	_, err = p.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return apperr.NewFromError(err, "Failed to insert professor bulk").SetLocation()
	}

	return nil
}

