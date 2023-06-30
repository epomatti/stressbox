# stressbox

Test your autoscaling configuration by running this image.

```sh
docker run -p 8080:8080 epomatti/stressbox
```

To change the default listener port, add `-e port=<PORT>` and set the publish parameter accordingly.

```sh
curl localhost:8080/cpu?x=42
```
