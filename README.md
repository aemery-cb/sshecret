# SSHecret 

unencrypted secret store using badgerdb and ssh for auth

# config

requires a config.toml file be located in the same directory as the binary.

```
<username> = "ssh-rsa blah"
```
# roadmap
- encryption
- configure config and db location
- HOTP and TOTP codes
