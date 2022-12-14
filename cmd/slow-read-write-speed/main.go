package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

func main() {
	home := os.Getenv("HOME")
	mntDir, _ := ioutil.TempDir("", "")

	////mount to an existing directory, to import threat0
	//mntDir := home + "/documents"

	// Make $HOME available on a mount dir under /tmp/ . Caution:
	// write operations are lso mirrore
	root, err := fs.NewLoopbackRoot(mntDir)
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
			AllowOther:    true,
		},
	}

	// Mount the file system
	server, err := fs.Mount(home+"/Desktop", root, mountOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Serve the file system, until unmounted by calling fusermount -u
	server.Wait()
}
