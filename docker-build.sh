env GOOS=linux GOARCH=amd64 go build
docker build -t quay.io/bszeti/battlefield-ui .
docker push quay.io/bszeti/battlefield-ui
rm -f battlefield-ui
