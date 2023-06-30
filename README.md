# stressbox

Simulate CPU and memory autoscaling behavior by running this image on your cloud infrastructure with your favorite load testing tool.

```sh
docker pull ghcr.io/epomatti/stressbox
```

### Running the image

Simply run it with the port binding:

```sh
docker run -p 8080:8080 ghcr.io/epomatti/stressbox
```

To change the default listener port, add `-e port=<PORT>` and set the publish parameter accordingly.

### CPU

Call the `/cpu` endpoint to simulate high CPU usage.

The `x` parameter is a simple Fibonacci length. Adjust to your requirements.

```sh
curl localhost:8080/cpu?x=30
```
