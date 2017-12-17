# slow-loris
Slow Loris attack written in Go

# Usage
Install: `go get github.com/awoldes/slow-loris`
Launch: `slow-loris -u https://test.com -p 443 -c 200`

## Defaults
Running `slow-loris -u http://test.com` will use port 80 as a default, with 100 concurrent connections.

## Help
Runing `slow-loris -h` will show the arguments that can be passed to the program.

# Disclaimer
A slow loris attack is **malicious**. This program was not created for malicious use or with malicious intent, but just for fun. The author is not responsible for any damage cauded by this package. 
