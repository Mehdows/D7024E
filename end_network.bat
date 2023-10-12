@echo off

for /l %%x in (1, 1 , 3) do (
    if %%x LSS 10 (docker stop web0%%x) else (docker stop web%%x)
    if %%x LSS 10 (docker remove web0%%x) else (docker remove web%%x)
)

docker network rm myNetwork
