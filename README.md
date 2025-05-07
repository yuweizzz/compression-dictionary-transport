# compression-dictionary-transport

A simple OpenResty implementation of HTTP compression dictionary transport and related tools.

## Tools

brotli_dictionary_generator: A dictionary generator provided by Brotli, used to train the dictionary.

zstd_dictionary_converter: Converts the Zstandard-formatted dictionary to a raw dictionary.

## Openresty implementation

Depends on Zstandard; include [zstd shared library](https://github.com/facebook/zstd) and [zstd-ffi](https://github.com/yuweizzz/zstd-ffi).

### Dynamic compression

Dynamic compression requires training the dictionary first, and all subsequent content will be compressed using the same dictionary.

To train the dictionary:

```bash
# use zstd to train dictionary
zstd --train content/* -o zstd_dict.dat
zstd_dictionary_converter -i zstd_dict.dat -o dict.dat

# use brotli_dictionary_generator
brotli_dictionary_generator dict.dat content/*
```

Change the `dict_location` in the OpenResty configuration.

### Static compression

Static compression will use the response as a dictionary. It is suitable for resources that change little between different versions.

There is no need to train the dictionary first.

## Ref

- [RFC draft](https://datatracker.ietf.org/doc/draft-ietf-httpbis-compression-dictionary/)
- [Mozilla Guides](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Compression_dictionary_transport)
- [Chrome Docs](https://chromium.googlesource.com/chromium/src.git/+/main/docs/experiments/compression-dictionary-transport.md)
- [Cloudflare worker for implementing HTTP compression dictionary transport](https://github.com/pmeenan/dictionary-worker)
- [Getting Real (small) With Compression Dictionaries](https://calendar.perfplanet.com/2024/getting-real-small-with-compression-dictionaries/)
