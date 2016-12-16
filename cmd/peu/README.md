Peu
=====
un peu: a little

## About
Peu compresses and decompresses files.

Peu is a simple compression tool for compressing and decompressing files. Peu neither creates nor extracts archive. If you want a tool that works with archives see [car](https://github.com/mohae/car) Peu can operate on multiple files concurrently. For decompression operations, Peu will determine the format used to compress the file and extract accordingly. If the format used to compress a file is not supported, an error will occur and nothing will be done with that file.

## Usage
Peu is written in Go and should work on any supported platform and architecture combination.

To see uasage and flag information: 

    $ peu -h

### Running `peu`
Peu will accept 1 or more filenames, each file is processed individually: they will not be put into an archive.

#### Compress
By default, Peu uses `gzip` compression. To specify a different compression format, use the `-f` flag. Multiple files can be specified by providing a space separated list.

To compress:

    peu <filename> ...

##### Supported formats for compression
    * gzip  
    * lz4

#### Decompress
Peu uses the filename, minus it's extension, as the output filename, e.g.:

Input | Output  
:--|:--  
example.txt | example  
example.txt.gz | example.txt  
path/to/example.txt.gz | path/to/example.txt  


To decompress:

    peu d <filename> ...

##### Supported formats for decompression
    * bzip2
	* gzip
	* lz4

## License
Copyright 2015-2016 by Joel Scoble.

This is provided under the Apache 2 License. For more details, please check the included LICENSE file.
