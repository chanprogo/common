  
```
docker image build -t test:latest .
docker image ls
docker images
```

`docker image build` == Old: `docker build`  




`docker image rm [OPTIONS] IMAGE [IMAGE...]`  
`docker image rm $(docker image ls -q) -f`  
`docker image prune`  
