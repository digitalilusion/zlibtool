# ZLib tool for inflate/deflate files

Simple tool to inflate or deflate files from shell.
The missing "gzip" equivalent for ZLib.

1. Install:

  Go get from repository
  ```
  go get github.org/digitalilusion/ziptool
  ```

2. Usage

  To deflate (compress) a file ```file.txt``` into ```file.txt.zlib```
  ```
  zlibtool -d file.txt file.txt.zlib
  ```

  To inflate (decompress) a file ```file.txt.gz``` into ```file.txt```
  ```
  zlibtool file.txt.zlib file.txt
  ```
