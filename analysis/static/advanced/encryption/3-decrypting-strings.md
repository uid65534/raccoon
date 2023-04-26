# 3. Decrypting strings
Use the `rc4` binary to decrypt base64 encoded strings present in the malware.

### Linux installation
Unzip the archive and copy the `rc4` binary into your `/bin` folder.
```bash
unzip rc4.zip
chmod +x rc4
sudo cp rc4 /bin
```

### Usage
Use `-k` to specify the key, and `-b64` to specify base64 encoded data.
```
rc4 -k edinayarossiya -b64 VXso/wY72A==
```

You can also export the key to the `RC4_KEY` environment variable so you don't need to specify it in the command.
```
export RC4_KEY=edinayarossiya
rc4 -b64 VXso/wY72A==
```
