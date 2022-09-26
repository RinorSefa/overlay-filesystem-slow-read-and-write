package main

import (
	"log"
	"os"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

func main() {
	home := os.Getenv("HOME")

	// Make $HOME available on a mount dir under /tmp/ . Caution:
	// write operations are lso mirrore
	root, err := fs.NewLoopbackRoot(home)
	if err != nil {
		log.Fatal(err)
	}

	mountOpts := &fs.Options{
		MountOptions: fuse.MountOptions{
			Debug:         true,
			MaxWrite:      10,
			MaxReadAhead:  10,
			SyncRead:      true,
			MaxBackground: 1,
		},
	}

	// Mount the file system
	server, err := fs.Mount(".", root, mountOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Serve the file system, until unmounted by calling fusermount -u
	server.Wait()
}
