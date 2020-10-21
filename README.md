Sample Microservice
======

# Usage
## Golang

### execution type: docker
```shell script
cd ~/{devicename}/Runtime
git clone {url: sample-microsevice.git}
cd sample-microservice/golang
make
cp ./config/docker.yml /var/lib/aion/config/project.yml
cd ~/{devicename}/aion-core/deployment/k8s
bash ./kubectl-apply.sh
```

### execution type: directory
```shell script
cd ~/{devicename}/Runtime
git clone {url: sample-microsevice.git}
cd sample-microservice
cp ./config/directory.yml /var/lib/aion/config/
./kanban-server
./sevice-broker -c /var/lib/aion/config/directory.yml
```
