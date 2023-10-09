@echo off

docker build . -t kadlab

for /l %%x in (1, 1 , 3) do (
    if %%x LSS 10 (docker run -d --name web0%%x -p 800%%x:80 kadlab) else (docker run -d --name web%%x -p 80%%x:80 kadlab)
)

docker ps
docker network create myNetwork

for /l %%x in (1, 1 , 3) do (
    if %%x LSS 10 (docker network connect myNetwork web0%%x) else (docker network connect myNetwork web%%x)
)

docker network inspect myNetwork
docker exec -it web01 /bin/bash 



