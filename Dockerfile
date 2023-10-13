FROM golang:latest


RUN apt-get update && apt-get install inetutils-ping -y

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
#RUN go get github.com/Mehdows/D7024E/test/GOWEBAPI/main
RUN cd /build && git clone https://github.com/Mehdows/D7024E.git
RUN cd /build/D7024E/ && git pull

RUN cd /build/D7024E/ && go build -o main .


WORKDIR /build/D7024E/
EXPOSE 80

#ENTRYPOINT [ "/build/D7024E/main" ]

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
