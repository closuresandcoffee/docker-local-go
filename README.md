# Local Golang Docker Development Environment

## Prerequisites
You will need to have [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed.
You will need to copy `.env.example` to `.env`. Please see the `Dockerfile` and `docker-compose.yml` files for more information on how those settings are used.

## To bring up the container
While in the directory for this repository, run the following command:
```
docker compose up -d
```

## To bring down the container
While in the directory for this repository, run the following command:
```
docker compose down
```

## To log into container
```
docker ps # Will list the docker containers; grab the name or id
docker exec -it <id_or_name_of_container> bash
```