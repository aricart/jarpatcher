package main

import (
	"flag"
	"fmt"
	"github.com/aricart/jarpatcher/jars"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

func usage() {
	fmt.Printf("patch -s <sourceDir> -t <targetDir>\n")
}

var wg sync.WaitGroup
var replaced int

func main() {
	// the arguments we require
	sourceDirPtr := flag.String("s", "", "the source directory")
	targetDirPtr := flag.String("t", "", "the target directory to patch.")

	flag.Parse()

	if "" == *sourceDirPtr || "" == *targetDirPtr {
		usage()
		return
	}

	start := time.Now()

	// build a map of filename->bundleid name on both the source and destination
	sourceJars, sourceCount := jars.FindBundles(*sourceDirPtr)
	targetJars, targetCount := jars.FindBundles(*targetDirPtr)

	// reorganize target bundles so they are indexed by bundleid to a list of files
	work := invert(*targetJars)

	// for each source file, if the file is on the destination replace it
	toReplace := 0
	replaced = 0
	for sourcePath, bundleID := range *sourceJars {
		var matches = work[bundleID]
		if matches != nil {
			patch(bundleID, sourcePath, matches)
			toReplace += len(matches)
		}
	}

	wg.Wait()

	elapsed := time.Since(start)

	fmt.Printf("Analyzed %d source file[s] and found %d bundle[s].\n", sourceCount, len(*sourceJars))
	fmt.Printf("Analyzed %d target file[s] and found %d bundle[s].\n", targetCount, len(*targetJars))
	fmt.Printf("Replaced %d/%d file[s].\n", replaced, toReplace)
	fmt.Printf("Elapsed time %s.\n", elapsed)

}

func patch(bundleID string, src string, targets []string) {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}
	// add a wait for each of the copies
	wg.Add(len(targets))
	for _, target := range targets {
		// a closure here, that way we don't duplicate the data buffer
		// and also get access to the waitgroup
		go func() {
			err := ioutil.WriteFile(target, data, 0644)
			// update the waitgroup
			defer wg.Done()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("[R] %s: %s\n", bundleID, target)
			replaced++
		}()
	}
}

// Takes a map of k and v, and returns a map of v to list of k
func invert(src map[string]string) map[string][]string {
	retVal := make(map[string][]string)
	for k, v := range src {
		var a = retVal[v]
		if a == nil {
			a = make([]string, 0, 0)
		}
		a = append(a, k)
		retVal[v] = a
	}
	return retVal
}
