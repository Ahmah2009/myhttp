## MyHttp
```myhttp``` tool which makes http requests and prints the address of the request along with the MD5 hash of the response.

## Examples
``` bash
$> ./myhttp http://www.adjust.com http://google.com

http://google.com e8a48fc244c07743ad85c88f1db22de6
http://www.adjust.com 1456e2ead9a2279f6174bfa4524e3a94
```

```bash
$> ./myhttp adjust.com
http://adjust.com 1456e2ead9a2279f6174bfa4524e3a94

```
```bash
$> ./myhttp -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com

http://google.com da848060fd247320c5c3e9f77c02f6ff
http://adjust.com 1456e2ead9a2279f6174bfa4524e3a94
http://facebook.com a3440d56c22b7ecf9f2f2ff2c2f60a03
http://yandex.com 70e7a027b62ca47d512d0cb736b6762c
http://twitter.com c9a81eca2353ca9cda7f69d04a0d6a85
http://reddit.com/r/notfunny ff3b2b7dcd9e716ca0adcbd208061c9a
http://reddit.com/r/funny ff3b2b7dcd9e716ca0adcbd208061c9a
http://baroquemusiclibrary.com 90ae3b48812f3a109715ac1224e7125d
http://yahoo.com 4d28f5fae9f23af5a6cb4af02074eaef
```

## Testing

to run the unit test for a specific module execute

```bash
go test
```