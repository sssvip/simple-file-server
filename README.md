# simple-file-server
Transfer file as simple as possiable via HTTP

# Usage
1. Download the binary from the [release page](https://github.com/sssvip/simple-file-server/releases), choose your os relational edition

2. Put the binary to you want transfer or browse folder
    
    windonws: open cmd in this folder then ->

         simplefileserver.exe 
         2018/01/11 16:52:53 starttd file server http://127.0.0.1:80
    
    linux. open terminal in this folder then->
    
        $ chmod 777 simplefileserver
        $ ./simplefileserver
        2018/01/11 16:52:53 starttd file server http://127.0.0.1:80
        
3. Custom file server port
    
    a. open on port 80 (Default): 
        
        > simplefileserver.exe 80
        $ ./simplefileserver 80
    b. open on port 8080:
        
        > simplefileserver.exe 8080
        $ ./simplefileserver 8080