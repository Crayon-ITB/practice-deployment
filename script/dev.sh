docker build -t golang-pipeline:dev .

docker run --rm --env-file .env -p 8080:8080 golang-pipeline:dev