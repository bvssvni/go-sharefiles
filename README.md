go-sharefiles
=============

A program for sharing files with other machines through the browser.  
BSD license.  
For version log and license, see the individual files.  

ShareFiles

ShareFiles lists links to all files in the "shared" directory at the address:  

    http://<your ip>:8080

When you click on a link, it opens a new tab with the content.  
You can use it to:

1. Display images, text, HTML files on the local network.  
2. Share files.  

##Download

[OSX 64 bit](go-sharefiles/sharefiles-osx-64bit)

[Linux 64 bit](go-sharefiles/sharefiles-linux-64bit)

##Build

To build from source, you need ![Go](http://golang.org/) installed.  

1. Download this repository.

2. In the Terminal window, type:

    go build sharefiles.go
    
3. Run by typing

    ./sharefiles    (Mac, Linux)
    
    sharefiles      (Windows)
    
