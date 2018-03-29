# GO ECHO

A Simple go TCP echo server. Written to learn and test [Kubernetes] networking.

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