# Data Analysis Training
This repo provides the materials needed for the lab session of the data analysis training. 

## Setup
This project requires access to a kubernetes (k8s) cluster. For local testing, you can use [kind](https://kind.sigs.k8s.io/) to set up a local k8s cluster.

### 1. Creating a local kubernetes cluster
If you have `go1.16` or greater on your machine: 
```bash
go install sigs.k8s.io/kind@v0.25.0 # local kube cluster
```
(Alternatively, you can install `kind` via `homebrew` as well.)

```bash
kind create cluster # Use --name to specify a name; defaults to `kind`
```

### 2. Build the docker images for the demo webserver
After the image has been built, load the image into your local k8s cluster:
```bash
kind load docker-image my-custom-image-0 my-custom-image-1
```

### 3. Configure OTLP endpoint
If you are using Grafana Cloud, please follow the guide [here](https://grafana.com/docs/beyla/latest/quickstart/golang/#3-optional-get-grafana-cloud-credentials) to retrieve your OTLP endpoint and API access credentials.

Update the environment variables for the `beyla` container in `deploy.yml`:
```yaml
- name: OTEL_EXPORTER_OTLP_ENDPOINT
  value: "<your Grafana Cloud OTLP endpoint>"
- name: OTEL_EXPORTER_OTLP_HEADERS
  value: "Authorization=Basic <your Grafana Cloud access token>"
```


### 4. Deploy!
Deploy the demo webserver application to your local k8s cluster:
```bash
kubectl apply -k ./k8s
```