# Brackets [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/beeker1121/brackets) [![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/beeker1121/brackets/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/beeker1121/brackets)](https://goreportcard.com/report/github.com/beeker1121/brackets)

Brackets formats data in any file with a single example of what the new format should look like.

## Usage

Build the program:

```go
$ go build
```

This example so far comes with two test files, `list.csv` and `json.txt`.

Run the program using a test file

```sh
$ ./brackets list.csv
```

To provide an example when prompted, use brackets {} to signify the start and end position of matches in the source string:

Source: `1,      "Sometimes Things Get, Whatever"        7:15`  
 Input: `Track {1} is {7}m{15}s long`  
Result: Track 1 is 7m15s long  

So far, this version cannot handle escaping brackets.

## Example

```
$ go build
$ ./brackets list.csv
First 5 records of file:
1.      "Sometimes Things Get, Whatever"        7:15
2.      "Complications" 5:31
3.      "Slip"  6:44
4.      "Some Kind of Blue"     6:19
5.      "Brazil (2nd Edit)"     5:23

Please provide an example of how to format the first record:
Track #{1} is "{Sometimes Things Get, Whatever}" and is {7}m{15}s long

Results:
Track #1 is "Sometimes Things Get, Whatever" and is 7m15s long
Track #2 is "Complications" and is 5m31s long
Track #3 is "Slip" and is 6m44s long
Track #4 is "Some Kind of Blue" and is 6m19s long
Track #5 is "Brazil (2nd Edit)" and is 5m23s long
Track #6 is "Alone With You" and is 7m30s long
Track #7 is "I Remember" and is 9m07s long
Track #8 is "Faxing Berlin (Piano Acoustic Version)" and is 1m39s long
Track #9 is "Faxing Berlin" and is 2m36s long
Track #10 is "Not Exactly" and is 8m00s long
Track #11 is "Arguru" and is 5m30s long
Track #12 is "So There I Was" and is 6m49s long
```