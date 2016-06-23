# Brackets [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/beeker1121/brackets) [![License](http://img.shields.io/badge/license-mit-blue.svg)](https://raw.githubusercontent.com/beeker1121/brackets/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/beeker1121/brackets)](https://goreportcard.com/report/github.com/beeker1121/brackets)

Brackets formats data in any file with a single example of what the new format should look like.

## Usage

Build the program:

```go
$ go build
```

This example so far comes with two test files, `list.csv` and `tracks.csv`.

Run the program using a test file

```sh
$ ./brackets list.csv
```

To provide an example when prompted, use brackets {} to signify the start and end position of matches in the source string:

Source: `23,Michael Jordan,Bulls`  
 Input: `{Jordan} played for the {Bulls}`  
Result: Jordan played for the Bulls  

## Examples

```
$ ./brackets list.csv
First 5 records of file:
23,Michael Jordan,Bulls
33,Scottie Pippen,Bulls
25,Steve Kerr,Bulls

Please provide an example of how to format the first record:
{Michael} {Jordan} played for the {Bulls} as #{23}

Results:
Michael Jordan played for the Bulls as #23
Scottie Pippen played for the Bulls as #33
Steve Kerr played for the Bulls as #25
```

```
$ ./brackets tracks.csv
First 5 records of file:
1,JOYRIDE,The Box,4:38
2,Kryoman & Pairanoid,My Squads Lit Ft SHAQ,3:56
3,SNAKEHIPS,Money On Me ft. Anderson .Paak,2:53
4,Allen Watts,Out of Reach (Thomas Hayes Remix),7:17
5,ANDRU.,ALL NIGHT,1:36

Please provide an example of how to format the first record:
\{"id": {1}, "artist": "{JOYRIDE}", "title": "{The Box}", "time": "{4:38}"\}

Results:
{"id": 1, "artist": "JOYRIDE", "title": "The Box", "time": "4:38"}
{"id": 2, "artist": "Kryoman & Pairanoid", "title": "My Squads Lit Ft SHAQ", "time": "3:56"}
{"id": 3, "artist": "SNAKEHIPS", "title": "Money On Me ft. Anderson .Paak", "time": "2:53"}
{"id": 4, "artist": "Allen Watts", "title": "Out of Reach (Thomas Hayes Remix)", "time": "7:17"}
{"id": 5, "artist": "ANDRU.", "title": "ALL NIGHT", "time": "1:36"}
{"id": 6, "artist": "Mystery Skulls", "title": "Ghost", "time": "4:17"}
{"id": 7, "artist": "Spencer & Hill ft. Lil Jon", "title": "Less Go! (Porter Robinson Remix)", "time": "7:04"}
{"id": 8, "artist": "The M Machine", "title": "When It's Gone (Mat Zo Remix)", "time": "4:56"}
```