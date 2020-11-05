#!/bin/bash
path=$1
# shellcheck disable=SC2010
protoFiles=$(ls "$path" | grep .proto)
for fileName in $protoFiles
do
    # shellcheck disable=SC2027
    # shellcheck disable=SC2034
    tempProtoPath="./$path/$fileName"
    echo "$tempProtoPath"
    protoc --proto_path="./$path" --micro_out=. --go_out=. "$tempProtoPath"
done

