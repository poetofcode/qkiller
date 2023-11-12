### What's the purpose?

This is a program designed for Windows. It kills the qbittorrent process if the outline process (VPN client) is running at the same time.

### Make it as service

```
nssm.exe install qkiller "c:\work\qkiller\qkiller.exe"
nssm.exe start qkiller
```