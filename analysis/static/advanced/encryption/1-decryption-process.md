# 1. The decryption process
The second function call in the entry point (`FUN_00404036`) decrypts strings in the `.rdata` section and stores them as global variables in the `.data` section.

It uses 3 functions to decode, decrypt and convert strings which are then used throughout the malware.

![](img/decryption-asm-1.png)

```c
iVar1 = FUN_00401806((int)"VXso/wY72A==",&local_8);
DAT_0040eb1c = FUN_00408746(&DAT_0040e228,iVar1,&local_8,(int)"edinayarossiya");
```

![](img/Pasted%20image%2020230427094228.png)

```c
DAT_0040e208 = FUN_0040a59a(DAT_0040eb1c);
```
First, a base64[^1] encoded string is decoded into raw data:

![](img/Pasted%20image%2020230427095022.png)
```c
iVar1 = FUN_00401806((int)"VXso/wY72A==",&local_8);
```
This takes the base64 encoded string `"VXso/wY72A=="`, decodes it, and stores it temporarily.

`iVar1`/`EAX` will now contain a pointer to the raw data, which is encrypted:
```
00000000: 55 7b 28 ff 06 3b d8                             U{(..;.
```

Next, the the data is decrypted using the Rivest Cipher 4 (RC4) algorithm with the cipher key `"edinayarossiya"` and stores it in a global variable:

![](img/Pasted%20image%2020230427094840.png)
```c
DAT_0040eb1c = FUN_00408746(&DAT_0040e228,iVar1,&local_8,(int)"edinayarossiya");
```

`DAT_0040eb1c` will now contain a pointer to the decrypted string:
```
00000000: 5c 43 43 2e 74 78 74                             \CC.txt
```

At the end of the function, the decrypted strings are converted into wide character strings and pointers to those are stored as global variables.

![](img/Pasted%20image%2020230427094245.png)
```c
DAT_0040ea88 = FUN_0040a59a(DAT_0040ebc4);
```

`DAT_0040ea88` will now contain a pointer to the wide character string.

Some Windows APIs accept wide strings (2 bytes / 16 bits per character) so it depends which of the decrypted strings get used.

[^1]: Base64 is used to encode binary data and represent it using printable characters. It's recognizable by its use of the characters `a-z`, `A-Z`, `0-9`, `+`, and `/` *(base64 URL encoding uses the `-` and `_` characters instead of `+` and `/` which have a special meaning in URLs)*. With 64 characters you can represent 6 bits of information (`2^6=64`). You'll often see the `=` character as padding at the end of the string which pads the string length to a multiple of 4, since bytes are 8 bits, and 6x4=24 bits is divisable by 8.