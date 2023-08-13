# stressbox

Simulate stress/load behavior such as high CPU and memory by running this image on your infrastructure with your favorite load testing tool.

Available endpoints:

| Endpoint | Functionality | Example |
|----------|-------------|---------|
| /cpu?x={n} | Stresses the CPU by a factor of "x" | `curl 127.1:8080/cpu?x=42` |
| /mem?add={mb} | Increases the used memory in megabytes | `curl 127.1:8080/mem?add=100` |
| /tcp?addr={addr} | Tests a target TCP connection | `curl 127.1:8080/tcp?addr=google.com:443` |
| /envs?env={var} | Returns an environment variable | `curl 127.1:8080/envs?env=DB_NAME` |
| /json?size={n} | Returns a JSON batch | `curl 127.1:8080/json?size=10000` |
| /exit | Exits the application | `curl 127.1:8080/exit` |
| /log?={m} | Writes to standard out | `curl 127.1:8080/log?m=Hello` |

Additionally, hitting `/` will return a static "OK".

## Running the image

Get the image:

```sh
docker pull ghcr.io/epomatti/stressbox
```

Run it with the port binding:

```sh
docker run -p 8080:8080 ghcr.io/epomatti/stressbox
```

To change the default listener port, add `-e port=<PORT>` and set the publish parameter accordingly.

Call the `/cpu` endpoint to simulate high CPU usage. The `x` parameter is a simple Fibonacci length. Adjust to your requirements.

```
curl localhost:8080/cpu?x=30
```

## Load testing

Choose your favorite load testing tool. Here's an example with K6:

```js
// script.js
import http from 'k6/http';
import { sleep } from 'k6';

const URL = __ENV.URL
console.log(`Calling ${URL}`);

export default function () {
  const res = http.get(URL);
  sleep(0);
}
```

Running the load testing:

```sh
k6 run -e URL="https://yourserver/cpu?x=30" --vus 10 --duration 300s script.js
```

## Private repositories

Command samples to upload this image to your docker repository:

### AWS ECR

```sh
docker pull ghcr.io/epomatti/stressbox
docker tag ghcr.io/epomatti/stressbox "$account.dkr.ecr.$region.amazonaws.com/stressbox:latest"
aws ecr get-login-password --region $region | docker login --username AWS --password-stdin "$account.dkr.ecr.$region.amazonaws.com"
docker push "$account.dkr.ecr.$region.amazonaws.com/stressbox:latest"
```

### Azure Container Registry (ACR)

```sh
az acr login --name "$acr"
docker pull ghcr.io/epomatti/stressbox
docker tag ghcr.io/epomatti/stressbox "$acr.azurecr.io/$repository:latest"
docker push "$acr.azurecr.io/$repository:latest"
```
