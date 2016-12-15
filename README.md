Peu
=====
[![Build Status](https://travis-ci.org/mohae/peu.png)](https://travis-ci.org/mohae/peu)

un peu: a little

## About
Peu makes things smaller, hopefully. Peu can also decompress a given stream if it's compression format is supported. This can either be done by checking the magic number or by using the appropriate method. When the stream ends, the operation ends. 

Just provide Peu with an io.Reader and io.Writer and Peu will handle the rest for you. For compression, either call `Compress` and specify the compression format or call the correct compression function. For decompression, just call `Decompress`; it will figure out the compression format used, if it is a supported one.

Peu is not designed to work with archives.  For that, there's [carchivum](https://github.com/mohae/carchivum)

## Supported compression algorithms

* lz4
* gzip
* bzip2 (decompression only)

## CLI
Peu is also a cli tool and can be found in https://github.com/mohae/peu/tree/master/cmd/peu

## License
Copyright 2015-2016 by Joel Scoble.

This is provided under the Apache 2 License. For more details, please check the included LICENSE file.
