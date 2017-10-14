#!/bin/sh
go test -cover

cd ./model; go test -cover;cd ..
cd ./controller; go test -cover; cd ..
cd ./global; go test -cover; cd ..
cd ./middleware; go test -cover; cd ..
cd ./ui; go test -cover; cd ..
