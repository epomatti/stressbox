## Local development

Running the application:

```sh
go run main.go
```

Building and running the image locally:

```sh
docker build -t epomatti/stressbox .
docker run -p 8080:8080 --rm epomatti/stressbox
```

Testing

```
curl localhost:8080/cpu?x=42
```
