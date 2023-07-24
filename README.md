# stressbox

Simulate stress/load behavior such as high CPU and memory by running this image on your infrastructure with your favorite load testing tool.

Available endpoints:

| Endpoint | Functionality | Example |
|----------|-------------|---------|
| /        | Returns a static "OK" | `curl 127.1:8080` |
| /cpu?x={n} | Calculates a Fibonacci sequence | `curl 127.1:8080/cpu?x=42` |
| /envs?env={var} | Returns an environment variable | `curl 127.1:8080/envs?env=DB_NAME` |
| /tcp?addr={addr} | Tests a TCP connection | `curl 127.1:8080/tcp?addr=google.com:443` |
| /json?size={n} | Returns a JSON batch | `curl 127.1:8080/json?size=10000` |
| /exit | Exits the application | `curl 127.1:8080/exit` |

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

## Load Testing

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

## Private image repository

If you prefer or required running in your private registry, simply tag and push the image to your repository.

Example with AWS ECR:

```sh
docker pull ghcr.io/epomatti/stressbox
docker tag ghcr.io/epomatti/stressbox "$account.dkr.ecr.$region.amazonaws.com/stressbox:latest"
aws ecr get-login-password --region $region | docker login --username AWS --password-stdin "$account.dkr.ecr.$region.amazonaws.com"
docker push "$account.dkr.ecr.$region.amazonaws.com/stressbox:latest"
```