# RateMyUBProf

RateMyUBProf is a web application designed to help students rate and review professors at their university precisely Brawijaya University so other students could prepare the best for their upcomming class professor.

This application is using [Go v1.24](https://tip.golang.org/doc/go1.24), [PostgreSQL 16](https://www.postgresql.org/) and [RabbitMQ](https://www.rabbitmq.com/docs/download), make sure you already installed the required dependency for this application.

## Development Setup

You can use [devcontainer](https://docs.github.com/en/codespaces/setting-up-your-project-for-codespaces/adding-a-dev-container-configuration/introduction-to-dev-containers) to setup project development, just make sure you have installed docker already!.

1. Clone the repository:
```zsh
git clone https://github.com/devanfer02/ratemyubprof.git
cd ratemyubprof
```

2. Setup environment variables file `env.json` from [env.example.json](./env.example.json).

3. Setup database environment variables file `.db.env` from `.db.env.example`.

3. Run the migration with migrate cli or application main entry and makefile

```zsh
make migrate-up # use migrate cli
make migrate    # use defined method in main 
```

3. Run the application

```zsh
make run # Basic go run
make air # Live reload with air
```
   
4. Access the application at `http://localhost:{PORT}`.

5. To run unit tests, just run the command
```zsh
make test
```

## Deployment

### Docker
1. Edit database credentials in `docker-compose.yml` file
2. Spin up the containers with `docker compose`
```zsh
docker compose up -d
```

2. Access the application at `http://localhost`.

### Kubernetes
1. Edit image name, configmap in `deploy/kube` and other configurations
1. Apply the Kubernetes manifests
```zsh
kubectl apply -f deploy/kube
```

2. Access the application via the Kubernetes service.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Tech Stacks

[![My Skills](https://skillicons.dev/icons?i=golang,nginx,docker,kubernetes,postgres,rabbitmq)](https://skillicons.dev) 