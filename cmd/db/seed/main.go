package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/bytedance/sonic"
	"github.com/devanfer02/ratemyubprof/internal/app/professor/repository"
	"github.com/devanfer02/ratemyubprof/internal/dto"
	"github.com/devanfer02/ratemyubprof/internal/entity"
	"github.com/devanfer02/ratemyubprof/internal/infra/database"
	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"github.com/devanfer02/ratemyubprof/pkg/util/formatter"
)

func main() {
	var (
		err        error
		professors []dto.ProfessorStatic
		entities []entity.Professor
		fileName   = "data/dosenub.json"
	)

	log.Println("Fetching data from file: ", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if err := sonic.Unmarshal(data, &professors); err != nil {
		panic(err)
	}

	log.Println("Formatting data")
	entities = formatter.FormatProfessorStaticToEntity(professors)

	env := env.NewEnv()
	db := database.NewDatabase(env)
	repo := repository.NewProfessorRepository(db)
	client, err := repo.NewClient(false)

	if err != nil {
		panic(err)
	}

	log.Println("Inserting data")
	if err := client.InsertProfessorsBulk(context.Background(), entities); err != nil {
		panic(err)
	}
}
