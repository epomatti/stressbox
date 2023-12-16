# stressbox

![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/epomatti/stressbox/go.yml)
[![codecov](https://codecov.io/gh/epomatti/stressbox/graph/badge.svg?token=BR9Q424076)](https://codecov.io/gh/epomatti/stressbox) [![Go Report Card](https://goreportcard.com/badge/github.com/epomatti/stressbox)](https://goreportcard.com/report/github.com/epomatti/stressbox) [![codebeat badge](https://codebeat.co/badges/634c7208-edd0-4f05-8b87-a9627d9986db)](https://codebeat.co/projects/github-com-epomatti-stressbox-main)

Stressbox is a small Docker utility for Cloud Engineers and SREs that conveniently assists on testing common scenarios involving containerized infrastructure.

- **Load testing** - Simulate high CPU or memory usage to trigger autoscaling and alarms.
- **Stress testing** - Force stress test scenarios and observe the cluster resiliency.
- **Networking/firewall** - Test firewall rules and connectivity requirements within your infrastructure.

Check the [Private repositories](#private-repositories) section for examples on how to deploy the image to your cluster.

## Endpoints

Stressbox has a web interface. Deploy the image to your cluster and use the endpoints:

| Endpoint | Functionality | Example |
|----------|-------------|---------|
| /cpu?x={n} | Stresses the CPU by a factor of "x" | `curl 127.1:8080/cpu?x=42` |
| /mem?add={mb} | Increases the used memory in megabytes | `curl 127.1:8080/mem?add=100` |
| /tcp?addr={addr} | Tests a target TCP connection | `curl 127.1:8080/tcp?addr=google.com:443` |
| /envs?env={var} | Returns an environment variable | `curl 127.1:8080/envs?env=DB_NAME` |
| /json?size={n} | Returns a JSON batch of size "n" | `curl 127.1:8080/json?size=10000` |
| /exit | Exits the application | `curl 127.1:8080/exit` |
| /log?={m} | Writes to standard out | `curl 127.1:8080/log?m=Hello` |
| / | Returns a static OK | `curl 127.1:8080/` |

## Getting started

Running the image locally is an easy way to get to know the toolbox:

Pull the image:

```sh
docker pull ghcr.io/epomatti/stressbox
```

An `ARM64` release is now available on Docker Hub:

```sh
docker pull epomatti/stressbox:arm64
```

Start the container:

> 💡 Add `-e port=<PORT>` to start the listener in another port

```sh
docker run -p 8080:8080 ghcr.io/epomatti/stressbox
```

Call the `/cpu` endpoint to simulate high CPU usage:

```
curl localhost:8080/cpu?x=30
```

The `x` parameter is a simple Fibonacci length, adjust to your requirements.

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

### AWS Elastic Container Registry (ECR)

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
