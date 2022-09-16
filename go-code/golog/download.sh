#!/bin/bash
FILE() {
    file_name=("logger" "structure_zap" "unstructure" "file_writer" "stdout_writer")

    mkdir -p golog

    for name in ${file_name[@]}; do
        file_path="https://raw.githubusercontent.com/ngoctd314/go101/main/go-code/golog/$name.go"
        wget $file_path -O "golog/$name.go"
    done
}

FILE