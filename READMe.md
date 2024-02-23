# Fampay Youtube Assignment (GoLang + Echo)


## Tech stack used:
- Golang 
- Echo 
- MongoDB
 
## Prerequisites 
- Docker
## Steps to run the app

Execute the following commands in shell
- Step 1: Build the docker containers
```bash 
docker-compose build
```
- Step 2: Run the containers
```bash
docker compose up 
```

The (GoLang) backend will be accessible on `localhost:8080` and MongoDB on `localhost:27017`

## Apis available

To populate the data in MongoDB, you can use a cron which will run every 10 seconds once its initial run 
and populate data for the same query.
```curl
curl --location 'localhost:8080/cron/football'
```
Here football is a example query which can be replaced.

To search the saved data in MongoDB, you have the following apis:

- To get all videos (paginated):
```curl
curl --location 'localhost:8080/search?page=1'
```

- To search for a particular string (paginated):
```curl
curl --location 'localhost:8080/search/football?page=1' 
```
Here football is a example query which can be replaced.

- To add a new api key (non-persistent), it can be done using following curl:
```curl
 curl --location --request POST 'localhost:8080/keys/apiKey'
```
Here apiKey can be replaced with actual api key.

- Health checks can be done using the following queries:
1. DB health:
```curl
curl --location 'localhost:8080/readiness'
```
2. Server health:
```curl
curl --location 'localhost:8080/liveness'
```