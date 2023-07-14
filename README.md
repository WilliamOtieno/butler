### Butler

#### Easily bootstrap systems with common developer tools/workflows


## Features
 - Install common tools e.g docker


## Usage:
    butler <command> e.g butler install docker redis k8s

## Installation
 - All binaries for specific operating systems are located in the [bin directory](/bin)
 - Run one of the 2 commands below according to your system.


    curl -O https://github.com/WilliamOtieno/butler/raw/main/bin/<specific-binary>
    wget https://github.com/WilliamOtieno/butler/raw/main/bin/<specific-binary>

 - To add butler to PATH however, run
    ```sudo cp path/to/specific-bin /bin/```
 - To build from source, make sure you have Golang installed then run ``go build -v -o .`` in the project root.