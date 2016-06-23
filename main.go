package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// openFile tries to open the given file.
func openFile(filename string) (*os.File, error) {
	// Open the file.
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	// Get file info.
	info, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, err
	}

	// Check if directory.
	if info.IsDir() {
		file.Close()
		return nil, ErrDirectory
	}

	// Check if file is empty.
	if info.Size() == 0 {
		file.Close()
		return nil, ErrEmpty
	}

	return file, nil
}

// getFilename prompts the user for the name of the file.
func getFilename() (string, error) {
	// Prompt for filename.
	fmt.Print("Enter filename: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return filename, nil
}

// getExample prompts the user to provide an example of
// how they want the file data formatted based on the
// first record.
func getExample(filedata []byte) (string, error) {
	// Show first 5 records of file.
	fmt.Println("First 5 records of file:")
	bfd := bytes.NewBuffer(filedata)
	scanner := bufio.NewScanner(bfd)

	var i int
	for scanner.Scan() {
		if i >= 5 {
			break
		}

		fmt.Println(scanner.Text())
		i++
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	fmt.Println("")

	// Prompt for example.
	fmt.Println("Please provide an example of how to format the first record:")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ex := scanner.Text()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return ex, nil
}

func main() {
	// Get the filename.
	var filename string
	var err error

	if len(os.Args) != 2 {
		filename, err = getFilename()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	} else {
		filename = os.Args[1]
	}

	// Open the file.
	file, err := openFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read and store all data from the file.
	fd, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Prompt for example.
	ex, err := getExample(fd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Parse the example to get bracket matches.
	matches, err := parse(ex, fd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Output the formatted data.
	fmt.Print("\nResults:\n")
	if err := output(os.Stdout, ex, fd, matches); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
