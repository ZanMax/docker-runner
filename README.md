# Docker Runner

Run different prepared docker containers from UI / console

## Recipes
You can configure your own recipes in `config.json` file.
Preconfigured recipes are in `config.json` file:

- MySQL      	      
- Postgres   	      
- Mongo
- Redis
- RabbitMQ
- Memcached

## Additional recipes
You can add additional recipes to the `recipes` folder. 
docker-compose.yml file is required for each recipe.
README.md file is optional and will be displayed in the console after docker-compose up.
Recipe examples is in `recipes` folder:

- Gitlab
- Jenkins
- Pi-hole
- Wireguard

## Compile

### Linux
GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64 docker-runner.go

### MacOS
GOOS=darwin GOARCH=amd64 go build -o bin/macos-amd64 docker-runner.go