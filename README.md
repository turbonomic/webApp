# webApp #
webApp in container image for simulating CPU intensive or Memory intensive workload.

# introduction #
This web app is built with apache httpd and php. It can simulate  CPU intensive, Memory intensive, or High latency workload.


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

Thrid, access the web Page
```console
http://hostIP:28080/index.html
or
http://localhost:28080/index.html
```
