# 2. Back tracing encrypted strings
If you see a function call that accepts a string, and there's a pointer to an address in the `.data` section, it may be a string that gets decrypted at runtime.

![](img/ghidra-decompilation-GetEnvironmentVariableW.png)
![](img/ghidra-disassembly-GetEnvironmentVariableW.png)

Here we have a call to `GetEnvironmentVariableW`, with a `.data` section address as the first argument.  Double-click the `DAT_0040ead8` in the disassembly to navigate to its address in the `.data` section.

![](img/ghidra-data-0040ead8.png)

Here we can see cross-references to this address. We're interested in the write `(W)`, which is where this address gets written to. Double-click `FUN_00404036:0040504f(W)` to navigate to the instruction that writes to this address.

![](img/Pasted%20image%2020230427100620.png)

We land on the `MOV [DAT_0040ead8],EAX` instruction. `EAX` holds the result of the previous `CALL FUN_0040a59a` instruction, which is the function that converts to a wide character string. As this function uses the `__fastcall` calling convention, the first argument is passed via the `ECX` register, not on the stack. So look up to the previous `MOV ECX,dword ptr [DAT_0040eab4]` before that `CALL` instruction. `DAT_0040eab4` will hold the address of the decrypted string that's being converted, so double-click it to navigate to its address in the `.data` section.

![](img/ghidra-data-0040eab4.png)

Then, follow the write again.

![](img/Pasted%20image%2020230427101014.png)

We land on the `MOV [DAT_0040eab4],EAX` instruction. We can look up to see a `CALL FUN_00408746` instruction, which is the decrypt function, so `EAX` contains a pointer to the decrypted string. And since the strings are first decoded from base64, the call before that should be the base64 decode function. Just above that call we can see the base64 encoded string stored in the `.rdata` section being passed via the `ECX` register. Now we can decrypt it with the `rc4` tool.