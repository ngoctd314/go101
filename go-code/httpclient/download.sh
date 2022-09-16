#!/bin/bash
FILE() {
    file_name=("args" "client" "response" "client_option")

    mkdir -p httpclient

    for name in ${file_name[@]}; do
        file_path="https://raw.githubusercontent.com/ngoctd314/go101/main/go-code/httpclient/$name.go"
        wget $file_path -O "httpclient/$name.go"
    done
}

FILE