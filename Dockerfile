# Add the commands needed to put your compiled go binary in the container and
# run it when the container starts.
#
# See https://docs.docker.com/engine/reference/builder/ for a reference of all
# the commands you can use in this file.
#
# In order to use this file together with the docker-compose.yml file in the
# same directory, you need to ensure the image you build gets the name
# "kadlab", which you do by using the following command:
#
# $ docker build . -t kadlab
# docker run -p 8080:8080 -tid kadlab

FROM golang:latest

WORKDIR /app
RUN cd /app

RUN apt-get update 
RUN git clone https://github.com/Mehdows/D7024E.git
RUN cd D7024E && git pull

RUN cd D7024E && go build -o main .

EXPOSE 8080

ENTRYPOINT [ "D7024E/main" ]