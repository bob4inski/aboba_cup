#!/bin/bash

cp docker-compose.yml $1
if [ -f $1/.env ]
    then
    rm $1/.env
fi

cp .env $1/.env
cd $1

echo "FILE=$2" >> .env
echo "PATHFOLDER=$1" >> .env
echo "URL=$3" >> .env
echo "USER=$4" >> .env




