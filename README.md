## freeformatter

### A simple golang script that will do below:

    - [x] show all the MIME listings
    - [x] get a specific mime information based on the extention
    - [x] encode or decode a given string of data
    - [x] html escape a given string of data
    - [x] base64 encode or decode a given string of data

### Command


```sh

$ ./freeformatter

 Version:  0.1.0-0

  -b64decode string
        use to base64 encode/decode a data
  -b64decode-url string
        use to base64 encode/decode a data
  -b64encode string
        use to base64 encode/decode a data
  -b64encode-url string
        use to base64 encode/decode a data
  -decode string
        use to encode/decode a data
  -decode-url string
        use to encode/decode a data
  -encode string
        use to encode/decode a data
  -encode-url string
        use to encode/decode a data
  -html-esc string
        use to escape/unescape an HTML
  -html-esc-url string
        use to escape/unescape an HTML
  -mime-list
        use to show the list of MIME types
  -mime-type string
        use to show the list of MIME types
  -qr-code-gen string
        use to generate a PNG QR Code

-------------------------------------------------------


        ./freeformatter  --encode "test data here"

        ./freeformatter  --decode "test+data+here"

        ./freeformatter  --mime-type '.avi'

        ./freeformatter  --mime-list

        ./freeformatter  --qr-code-gen='{"data": "https://www.google.com.sg/","filename":"qrcode.png","size":256}'


-------------------------------------------------------

```



## Docker Binary

- [x] In order to  use it as dockerize binary


```sh

    sudo  sysctl -w net.ipv4.ip_forward=1

    sudo  docker run --rm  bayugyug/freeformatter 

    sudo  docker run --rm  bayugyug/freeformatter --mime-list

    sudo  docker run -v `pwd`:`pwd` -w `pwd` --rm  bayugyug/freeformatter         --qr-code-gen='{"data": "https://www.google.com.sg/","filename":"qrcode-hahaha.png","size":256}'

    sudo  docker run -v `pwd`:`pwd` -w `pwd` --rm  bayugyug/freeformatter:alpine  --qr-code-gen='{"data": "https://www.google.com.sg/","filename":"qrcode-hahaha.png","size":256}'

```


## As HTTP Server

- [x] In order to  use it via CURL/WGET or Browser


```sh

    sudo  sysctl -w net.ipv4.ip_forward=1

    sudo  docker run -p 7000-8000:7000-8000 -v `pwd`:`pwd` -w `pwd` -d --name freeformatter-alpine  bayugyug/freeformatter:alpine --http --port 7778

    curl -i -v 'http://192.168.2.121:7778/mime-type/?p=.mp4'

```

### License

[MIT](https://bayugyug.mit-license.org/)
