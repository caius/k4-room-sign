# Kindle 4 Room Sign

An idea for a door.

## Notes

Build shit for it with

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build ./...
```

`phd` is the Phone Home Daemon, and listens on port 40317 which is open in the firewall by default.

`powerd` appears to control the power functions.

`framework` is the kindle "app" that powers the UI/UX you see. Also listens to key presses - if you turn that off it won't respond to keys.

`eips` lets you draw stuff on screen

`powerd_test -s` tells you stuff about the power/battery state of the kindle, needs `powerd` running to use it.

Needs golang 1.16 or older because glibc is 2.12.1 on the device, which is ancient.
