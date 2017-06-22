# webApp #
webApp is container image to simulating CPU/Memory intensive or high latency workload.

# Introduction #
This web app is built with apache httpd and php. It can simulate  CPU intensive, Memory intensive, or High latency workload.
It can accept Get or Post requests to simulate different level of workloads.

# Run it #
First, build an docker image.
```console
$ export imageName="webAppImage"
$ docker build -t $imageName .
```

Second, run the image.
```console
# 28080 is host port.
$ docker run -d -p 28080:8080 $imageName
```

Thrid, access the web Page via:
```console
http://hostIP:28080/index.html
or
http://localhost:28080/index.html
```

# Test workload #

all the page can be accessed via HTTP.PUT or HTTP.POST, and the parameters are the same.

## Latency simulation ##
access it via web browser
```console
# value is time to delay(or sleep in server side), in millisecond.
http://localhost:28080/workload.php/?value=30
```
access it via curl Post
```console
curl -H 'Content-Type: application/x-www-form-urlencoded' -X PUT -d 'value=30' http://localhost:28080/workload.php
```
