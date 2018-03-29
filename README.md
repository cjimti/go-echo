[![Docker Container Image Size](https://shields.beevelop.com/docker/image/image-size/cjimti/go-echo/latest.svg)](https://hub.docker.com/r/cjimti/go-echo/)
[![Docker Container Layers](https://shields.beevelop.com/docker/image/layers/cjimti/go-echo/latest.svg)](https://hub.docker.com/r/cjimti/go-echo/)
[![Docker Container Pulls](https://img.shields.io/docker/pulls/cjimti/go-echo.svg)](https://hub.docker.com/r/cjimti/go-echo/)

# GO ECHO

A Simple go TCP echo server. Written to learn and test [Kubernetes] TCP networking.

## Test with [Docker]

Run the container from a terminal:
```bash
docker run --rm -it -e TCP_PORT=2701 -e NODE_NAME="EchoNode" -p 2701:2701 cjimti/go-echo
```

In another terminal run:
```bash
telnet localhost 2701
```

## Test with [Kubernetes]

```bash
cd k8s
kubectl create -f . 
```

You should now have two TCP echo containers running:

```bash
kubectl get pods --selector=app=tcp-echo
```

```bash
NAME                                   READY     STATUS    RESTARTS   AGE
tcp-echo-deployment-777d856787-5fhb4   1/1       Running   0          27s
tcp-echo-deployment-777d856787-rh9tp   1/1       Running   0          27s
```

You should also have a service that connection port 32701 to the pods:

```bash
kubectl get service --selector=app=tcp-echo-service
```

```bash
NAME               TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
tcp-echo-service   NodePort   10.102.207.113   <none>        2701:32701/TCP   35m

```

Echo some data, replace ANY_NODE_IP with a location of a node:
```
telnet ANY_NODE_IP 32701
```

After connecting, type the word hello and hit return:
```bash
Trying x.x.x.x...
Connected to node1.example.com.
Escape character is '^]'.
Welcome, you are connected to node node1.example.com.
Running on Pod tcp-echo-deployment-777d856787-rh9tp.
In namespace default.
With IP address 192.168.33.39.
Service default.
hello
hello
```

## Resources
- [Expose Pod Information to Containers Through Environment Variables]
- [Docker]
- [Kubernetes]


[Expose Pod Information to Containers Through Environment Variables]: https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/
[Docker]: https://www.docker.com/
[Kubernetes]: https://kubernetes.io/