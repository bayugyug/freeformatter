## freeformatter

    - [x] show all the MIME listings
    - [x] get a specific mime information based on the extention
    - [x] encode or decode a given string of data
    - [x] html escape a given string of data
    - [x] base64 encode or decode a given string of data

### Command List:


```sh
$ ./freeformatter
2016/09/20 11:11:05 Version:  0.1.0-0

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



        


