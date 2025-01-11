package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// What does the program do?
// Analizes project from a target path
// Gets: Line Count
// Tree: Output of Tree

// dumps -> output to file or terminal
// This means it needs cobra for setting up the command structure in a simple way
// Just use Flags: this should be better for this use case.

// Start: Just clone the Python Tool 1 -> 1

func dumpFileContents(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fileContents := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		fileContents = append(fileContents, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(filePath)
		return []string{}, err
	}

	return fileContents, nil
}

func main() {
	tree := true
	cloc := false
	fileData := true

	outFile, err := os.Create("output.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer outFile.Close()
	w := bufio.NewWriter(outFile)

	root := "/Users/cameronolson/Developer/dl_work/Prototypes/demo-app-testing/echelon-beta/choir-converter"
	fmt.Fprintf(w, "Project Root: %s\n\n", root)
	w.Flush()

	if cloc {
		clocCmd := exec.Command("cloc", ".")
		clocCmd.Dir = root
		clocCmd.Stdout = outFile
		err = clocCmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	if tree {
		treeCmd := exec.Command("tree", "-I", "*__pycache__")
		treeCmd.Dir = root
		treeCmd.Stdout = outFile
		err = treeCmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Need to check if it ends with something by adding *.pth
	toSkip := []string{
		".DS_Store", "__init__.py", ".pth", ".npy", ".index", ".wav", ".bin", ".pt", ".gitkeep", ".json",
	}

	if fileData {
		err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && info.Name() == "__pycache__" {
				return filepath.SkipDir
			}

			for _, val := range toSkip {
				if val == info.Name() || val == filepath.Ext(path) {
					return nil
				}
			}

			if info.Name() == ".DS_Store" || info.Name() == "__init__.py" {
				return nil
			}

			if !info.IsDir() {
				fileContents, err := dumpFileContents(path)
				if err != nil {
					log.Println(err)
				}
				fmt.Fprintln(w, "----------------------")
				fmt.Fprintf(w, "File Name: %s\n", info.Name())
				fmt.Fprintln(w, "----------------------")
				fmt.Fprintln(w)
				for _, line := range fileContents {
					fmt.Fprintln(w, line)
				}
				fmt.Fprintln(w)
				w.Flush()
			}
			return nil
		})

		if err != nil {
			log.Println(err)
			return
		}
	}
}
