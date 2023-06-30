# Local development

Running the application:

```sh
go run main.go
```

Testing locally:

```
curl localhost:8080/cpu?x=30
```

Building and running the Docker image locally:

```sh
docker build -t stressbox .
docker run -p 8080:8080 --rm stressbox
```

Push a tag to trigger a Docker build and publish:

```
git tag v*.*.*
git push --tags
```
