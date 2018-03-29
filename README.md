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


## Resources
- [Expose Pod Information to Containers Through Environment Variables]
- [Docker]
- [Kubernetes]


[Expose Pod Information to Containers Through Environment Variables]: https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/
[Docker]: https://www.docker.com/
[Kubernetes]: https://kubernetes.io/