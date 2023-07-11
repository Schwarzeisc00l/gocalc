## gocalc
Terminal based calculator I wrote with GO as a begginer project.

# Usage
`Number 1 (1, 300, 43.432432)`
`Operator (+,-,*,/,:..)`
`Number 2 (192, 3.14, 4096)`




# Install
Required packages are:
`wget` `sudo`

`cd /tmp && sudo wget https://raw.githubusercontent.com/Schwarzeisc00l/gocalc/main/install.sh && sudo bash install.sh`


# Build

Required packages:
`go`

Firstly clone the repo locally and change the directory inside.
`$ git clone https://github.com/Schwarzeisc00l/gocalc cd gocalc` 
Build.
`go build gocalc.go`

Give it permission to run as a binary.
`chmod +x gocalc`


Move the binary to /usr/bin to easily execute it.
`sudo mv gocalc /usr/bin/`

Profit.


