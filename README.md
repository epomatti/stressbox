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

Call the `/cpu` endpoint to simulate high CPU usage. The `x` parameter is a simple Fibonacci length. Adjust to your requirements.

```
curl localhost:8080/cpu?x=30
```

### Load Testing

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

```
k6 run -e URL="https://yourserver/cpu?x=30" --vus 10 --duration 300s script.js
```
