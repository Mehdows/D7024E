@echo off

for /l %%x in (1, 1 , 20) do (
    if %%x LSS 10 (docker stop web0%%x) else (docker stop web%%x)
)

docker network rm myNetwork
docker network inspect myNetwork
