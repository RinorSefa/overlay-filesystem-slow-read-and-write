# Instructions for installation

Tested environment requirement
- Raspberry PI 4 B Rev 1.4
- Debian GNU/Linux 11
- go version go1.18.4 linux/arm64
- $Home/Desktop directory should exist (Mount Point)

Navigate to the filesystem application for the file type and start the application.

```
$ cd cmd/slow-read-write-speed/main.go
$ go run main.go
```

To unmount the filesystem
```
$ fusermount -u mountPath
```



## Author of code in files

- cmd/slow-read-write-speed/main.go (author)
- fs/loopback.go (modified create and open, method)
- fuse/opcode.go (modified doRead function)
- fuse/read.go (modified Bytes method)
- fs/files.go (modified Write method)

# Additional notes

To avoid slow file copying, it is possible to change the mount directory to an existing directory.