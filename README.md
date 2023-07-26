### Butler

#### Easily bootstrap systems with common developer tools/workflows


## Features
 - Install common tools. Support currently for:-
   - Docker (Containerization tool)
   - Redis (In-Mem Store)
   - NVM (Node version manager)
   - Homebrew (MacOS/Linux Package Manager)



## Usage:
    butler <command> e.g butler install docker
    butler -h
    butler help
    butler install -h


## Installation
 - All binaries for specific operating systems are located in the [bin directory](/bin)
 - Run one of the 2 series of commands below according to your system.
 - These will simply install the binary to the `PATH` then make it executable.

    ```shell
   sudo curl -o /usr/bin/butler https://github.com/WilliamOtieno/butler/raw/main/bin/<specific-binary>
   sudo chmod a+rwx /usr/bin/butler
   ``` 
   ```shell
   sudo wget -O /usr/bin/butler https://github.com/WilliamOtieno/butler/raw/main/bin/<specific-binary>
   sudo chmod a+rwx /usr/bin/butler
   ```

 - To build from source, make sure you have Golang installed then run ``go build -v -o .`` in the project root.
