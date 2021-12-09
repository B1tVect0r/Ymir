#!/bin/bash

for f in $(find . -name '*.pb.go'); do
    rm $f
done

for f in $(find . -name '*.proto'); do
    protoc -I=. --go_out=. $f
done

cp -r github.com/B1tVect0r/ymir/pkg/* pkg
rm -rf github.com