package main

import (
	"os"
	"context"
	"io"
	"log"

	"github.com/bytedance/sonic"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/devanfer02/ratemyubprof/internal/infra/server"
	"github.com/devanfer02/ratemyubprof/internal/app/professor/repository"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/devanfer02/ratemyubprof/internal/infra/database/postgres"
	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"github.com/devanfer02/ratemyubprof/pkg/util/formatter"
)


func main() {

	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "seed":
			seedDB()
			return 
		case "migrate":
			migrateDB()
			return 
		}
	}

	server := server.NewHttpServer()
	server.Start()
	server.GracefullyShutdown()
}

func migrateDB() {
	log.Println("[RateMyUbProf] Running migration....")

	var (
		env = env.NewEnv()
		db = database.NewDatabase(env)
	)

	driver := must(postgres.WithInstance(db.DB, &postgres.Config{}))

	m := must(migrate.NewWithDatabaseInstance(
		"file://./internal/infra/database/postgres/migrations",
		env.Database.Name, driver,
	))

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	log.Println("[RateMyUbProf] Migration completed!")
}

func seedDB() {
	log.Println("[RateMyUbProf] Running seeding database....")

	var (
		professors []dto.ProfessorStatic
		entities   []entity.Professor
		fileName   = "data/dosenub.json"

		env = env.NewEnv()
		db = database.NewDatabase(env)
		repo = repository.NewProfessorRepository(db)
	)

	file := must(os.Open(fileName))
	defer file.Close()


	data := must(io.ReadAll(file))
	mustv(sonic.Unmarshal(data, &professors))

	entities = formatter.FormatProfessorStaticToEntity(professors)
	
	client := must(repo.NewClient(false))

	mustv(client.InsertProfessorsBulk(context.Background(), entities))


	log.Println("[RateMyUbProf] Seeding completed!")
}

func must[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}
	return result 
}

func mustv(err error) {
	if err != nil {
		panic(err)
	}
}