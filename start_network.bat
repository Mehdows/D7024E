@echo off

docker build . -t kadlab

for /l %%x in (1, 1 , 3) do (
    if %%x LSS 10 (docker run -td --name web0%%x -p 800%%x:80 kadlab) else (docker run -td --name web%%x -p 80%%x:8%%x kadlab)
)

docker ps
docker network create myNetwork

for /l %%x in (1, 1 , 3) do (
    if %%x LSS 10 (docker network connect myNetwork web0%%x) else (docker network connect myNetwork web%%x)
)

docker exec -it web01 /bin/bash 



