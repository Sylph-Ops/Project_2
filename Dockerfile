FROM        golang:1.15.7-buster
WORKDIR     /app
COPY        ./  /app
ENV         DATABASE_URL="127.0.0.1:5432"
ENV         Limit = 10
EXPOSE      8080
CMD         [ "go", "run", "main.go" ]