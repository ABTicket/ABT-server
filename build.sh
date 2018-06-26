#!/bin/bash

git pull origin dev

sudo docker stop ticket-server

sudo docker cp . ticket-server:/app

sudo docker start ticket-server

sudo docker exec ticket-server /bin/bash -c \
    "cd src && go build main.go && nohup ./main 1> ticket.out 2> ticket.err"
