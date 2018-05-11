# webApp #
webApp is container image to simulating CPU/Memory intensive or high latency workload.

# Introduction #
This web app is built with apache httpd and php. It can simulate  CPU intensive, Memory intensive, or High latency workload.
It can accept Get or Post requests to simulate different level of workloads.

It also enables the **mod_status** of the httpd server, so that user can have a easier way to understand the workloads.


# Build docker image #
Build a docker image. (`beekman9527/webapp` is the image in dockerhub)
```console
$ export imageName="webapp"
$ docker build -t $imageName .
```

# Run container #
First, run the container from the image.
```console
# 28080 is host port.
$ docker run -d -p 28080:8080 $imageName
```

Second, access the web Page via:
```console
http://hostIP:28080/index.html
or
http://localhost:28080/index.html
```

# Test workload #

All the pages can be accessed via HTTP **GET** or **POST**, and the parameters are the same.

### Latency simulation ###
It has one parameter **value**, the duration to delay(or sleep in server side), in milliseconds.

For example, delay 30 ms. Access it via web browser
```console
http://localhost:28080/workload.php/?value=30
```
Access it via curl Post
```console
curl -H 'Content-Type: application/x-www-form-urlencoded' -X POST -d 'value=30' http://localhost:28080/workload.php
```

### Memory intensive simulation ###
It has two parameters, **memory** is amount of memory (in MB) to consume; the other is **value**, the duration to hold the 
memory, in milliseconds.

For example, consume 10 MB memory, and hold the memory for 110 ms.
Access it via web browser
```console
http://localhost:28080/memwork.php/?value=110&memory=10
```
Access it via curl Post
```console
curl -H 'Content-Type: application/x-www-form-urlencoded' -X POST -d 'value=110&memory=10' http://localhost:28080/memwork.php
```

### CPU intensive simulation ###
It will use as much CPU as possible to compute the MD5 for huge amout of integers. 
It has one parameter, **cpu**, indicating the amount of computation.

Access it via web browser
```console
http://localhost:28080/cpuwork.php/?cpu=30
```
Access it via curl Post
```console
curl -H 'Content-Type: application/x-www-form-urlencoded' -X POST -d 'cpu=30' http://localhost:28080/cpuwork.php
```

# Stable workload generator #
[webclient](https://github.com/songbinliu/webclient) can be used to generate stable workload for this webApp.
This `webclient` can also deployed in the same Kubernetes cluster to send requests to this webApp.
```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: httpclient
  namespace: default
  labels:
    purpose: generate-http-load
spec:
  replicas: 2
  selector:
    matchLabels:
      app: httpclient
  template:
    metadata:
      labels:
        app: httpclient
    spec:
      serviceAccount: default
      containers:
      - name: httpclient
        image: beekman9527/webclient:v1
        imagePullPolicy: IfNotPresent
        args:
        - --v=3
        - --threadNum=6
        - --logtostderr
        - --target=http://music.default:8080/cpuwork.php/?cpu=20
        - --rps=2
```
