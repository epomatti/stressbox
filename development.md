# Local development

Running the application:

```sh
go run main.go
```

Building and running the image locally:

```sh
docker build -t stressbox .
docker run -p 8080:8080 --rm stressbox
```

Testing locally:

```
curl localhost:8080/cpu?x=30
```

Push a tag to trigger a Docker build and publish:

```
git tag v0.0.1
git push --tags
```
