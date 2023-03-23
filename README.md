# curl-to-ffmpeg

ffmpeg headers format is a pain.
This tool is for building ffmpeg headers from curl command line.

## Usage

First, prepare a curl command line.
For example, using a browser, click on "Copy as cURL".
And

```sh
$ curl-to-ffmpeg << EOS
curl 'https://example.com/' \
  -H 'accept: */*' \
  -H 'cache-control: no-cache'
EOS
```

Then, got

```
ffmpeg -headers 'accept: */*'$'\r\n''cache-control: no-cache' -i https://example.com/
```

Give additional an output file name, etc, and run.
