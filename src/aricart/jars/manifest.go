package jars

import (
	"archive/zip"
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// A Manifest here is simply just a map of Manifest headers to their values
type Manifest struct {
	Map map[string]string
}

// Parse a Manifest from a string representation
func (m *Manifest) Parse(s string) {
	if m.Map == nil {
		m.Map = make(map[string]string)
	}

	index := strings.IndexRune(s, ':')
	if index != -1 {
		key := s[0:index]
		value := strings.TrimSpace(s[index+1:])
		m.Map[key] = value
	}
}

// PrintHeaders prints the headers and values in the manifest
func (m *Manifest) PrintHeaders() {
	fmt.Printf("%s=%d\n", "Headers found in Manifest", len(m.Map))
	for k, v := range m.Map {
		fmt.Println(k)
		fmt.Println(v)
	}
}

// BundleSymbolicName returns the Bundle-SymbolicName in the Manifest or nil.
func (m *Manifest) BundleSymbolicName() string {
	v := m.Map["Bundle-SymbolicName"]
	index := strings.IndexRune(v, ';')
	if index != -1 {
		v = v[0:index]
	}
	return v
}

// ParseJar opens a jar, extracts and parses the Manifest file
func ParseJar(fileName string) (*Manifest, string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s\n", fileName)
		return nil, ""
	}

	zip, err := zip.OpenReader(fileName)
	if err != nil {
		fmt.Printf("Unable to process: %s\n\t[%q]\n", fileName, err)
		return nil, ""
	}
	defer zip.Close()

	var retVal Manifest

	for _, f := range zip.File {
		if f.Name == "META-INF/MANIFEST.MF" {
			zipEntry, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			retVal = Manifest{}
			reader := bufio.NewReader(zipEntry)
			if reader != nil {
				header := make([]string, 0)
				for {
					bytes, _, err := reader.ReadLine()
					if len(bytes) > 0 {
						if len(header) == 0 {
							header = append(header, string(bytes))
						} else if bytes[0] != ' ' {
							header = append(header, "\n")
							retVal.Parse(strings.Join(header, ""))
							header = make([]string, 0)
							header = append(header, string(bytes))
						} else {
							header = append(header, string(bytes[1:]))
						}
					}
					if err != nil {
						break
					}
				}
			}
			zipEntry.Close()
			break
		}
	}

	return &retVal, ""
}

// FindBundles examines all the jars in a file or directory and returns a
// map of all bundles it finds.
func FindBundles(path string) (*map[string]string, int) {
	var pathToJar = make(map[string]string)
	var counter = 0

	vis := func(path string, info os.FileInfo, error error) error {
		counter++
		if info.IsDir() || !strings.HasSuffix(path, ".jar") {
			return nil
		}
		mf, _ := ParseJar(path)
		if mf != nil && mf.BundleSymbolicName() != "" {
			pathToJar[path] = mf.BundleSymbolicName()
		}
		return nil
	}

	filepath.Walk(path, vis)

	return &pathToJar, counter
}
