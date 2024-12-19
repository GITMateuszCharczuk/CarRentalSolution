#!/bin/bash

command_exists() {
    command -v "$1" &> /dev/null
}

if ! command_exists go; then
    echo "Golang is not installed. Installing..."
    wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz -O go.tar.gz
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go.tar.gz
    rm go.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
    echo "Golang installed successfully."
fi

export PATH=$PATH:/usr/local/go/bin

if ! command_exists wire; then
    echo "Wire is not installed. Installing..."
    export PATH=$PATH:$HOME/go/bin
    go install github.com/google/wire/cmd/wire@latest
    echo "export PATH=\$PATH:\$HOME/go/bin" >> ~/.bashrc
    echo "Wire installed successfully."
fi

export PATH=$PATH:$HOME/go/bin

echo "Running wire generation..."
wire ../cmd/service

if [ $? -eq 0 ]; then
    echo "Command executed succesfully"
else
    echo "Error occured"
fi
