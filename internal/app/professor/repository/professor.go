package repository

import (
	"context"
	"time"

	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/devanfer02/ratemyubprof/pkg/util"
	"github.com/doug-martin/goqu/v9"
	
)

func (p *professorRepositoryImplPostgre) FetchAllProfessors(ctx context.Context, params *dto.FetchProfessorParam, pageQuery *dto.PaginationQuery) ([]entity.Professor, error) {
	var professors []entity.Professor

	qb := goqu.
		Select("id", "name", "faculty", "major", "profile_img_link").
		From(professorTableName).
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	if pageQuery.Page != 0 && pageQuery.Limit != 0{
		qb = qb.Offset((pageQuery.Page - 1) * pageQuery.Limit).Limit(pageQuery.Limit)
	}

	qb = util.AddParamsToFetchProf(qb, params)

	query, args, err := qb.ToSQL()
	if err != nil {
		return nil, err
	}

	query = p.conn.Rebind(query)
	
	rows, err := p.conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var professor entity.Professor
		if err := rows.StructScan(&professor); err != nil {
			return nil, err
		}
		professors = append(professors, professor)
	}

	return professors, nil
}

func (p *professorRepositoryImplPostgre) FetchProfessorByID(ctx context.Context, id string) (entity.Professor, error) {
	var professor entity.Professor
	qb := goqu.
		Select("id", "name", "faculty", "major", "profile_img_link").
		From(professorTableName).
		Where(goqu.I("id").Eq(id)).
		SetDialect(goqu.GetDialect("postgres")).
		Prepared(true)

	query, args, err := qb.ToSQL()
	if err != nil {
		return professor, err
	}

	query = p.conn.Rebind(query)
	
	err = p.conn.QueryRowxContext(ctx, query, args...).StructScan(&professor)
	if err != nil {
		return professor, err
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
		return 0, err
	}

	query = p.conn.Rebind(query)

	err = p.conn.QueryRowxContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
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
		return err
	}

	query = p.conn.Rebind(query)

	_, err = p.conn.ExecContext(ctx, query, args...)
	if err != nil {

		return err
	}

	return nil
}

