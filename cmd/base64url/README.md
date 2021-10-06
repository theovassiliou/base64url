# base64url

base64url encodes and decodes Base64URL data, as specified in RFC 4648.  With no options, base64url reads raw data from stdin and writes encoded data as a continuous block to stdout.

Note: To decode hard wrapped encoded data, line breaks are removed before decoding.

## Usage

`base64url -i fileA.txt -o fileB.txt` encodes the content of file `fileA.txt` into a file `fileB.txt`

    Usage: base64url [options]

    Options:
    --decode, -d        decodes input
    --input, -i         input file ("-" for stdin, default -)
    --output, -o        output file ("-" for stdout, default -)
    --break, -b         Insert line breaks every count characters.  Default is 0, which generates an
                        unbroken stream
    --version, -v       display version
    --help, -h          display help

    Version:
        base64url 1.0-src
    Read more:
        github.com/theovassiliou/base64url

## Examples

`base64url -i testdocs/text/Text.txt -o testdocs/text/encodedText.txt`
base64url encodes the file `testdocs/text/Text.txt` into file `testdocs/text/encodedText.txt`.
This would be equivalent to `base64url -i testdocs/text/Text.txt > testdocs/text/encodedText.txt`

`base64url -d -i testdocs/text/encodedText.txt -o testdocs/text/decodedText.txt`
base64url decode the file `testdocs/text/encodedText.txt` in file `testdocs/text/decodedText.txt`

The same could be achieved with
`cat testdocs/text/encodedText.txt | base64url -d > testdocs/text/decodedText.txt`

To verify that original text and decoded text are exactly the same we could use

`diff testdocs/text/decodedText.txt testdocs/text/Text.txt`

Another proof that the encoding is reversible can be generated with

    $ echo "This is a Test" | base64url | base64url -d
    This is Test

This takes the string "This is a Test" as input via stdin and pipes the resulting encoded string as input
to the decoding.

Using the `-b count` parameter a line break every `count` character is added. `base64url` respects this possibility
and removes line breaks prior to decoding a string

    $ echo "This is a Test" | base64url -b 5
    VGhpc
    yBpcy
    BhIFR
    lc3QK
    $ echo "This is a Test" | base64url -b 5 | base64url -d
    This is a Test

The directory `testdocs/binary` contains a binary file `RandomNumbers.dat` you can play around with.
