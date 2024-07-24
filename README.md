# ASCII OUTPUT
```

             _____    _____   _____   _____             ____    _    _   _______   _____    _    _   _______  
    /\      / ____|  / ____| |_   _| |_   _|           / __ \  | |  | | |__   __| |  __ \  | |  | | |__   __| 
   /  \    | (___   | |        | |     | |    ______  | |  | | | |  | |    | |    | |__) | | |  | |    | |    
  / /\ \    \___ \  | |        | |     | |   |______| | |  | | | |  | |    | |    |  ___/  | |  | |    | |    
 / ____ \   ____) | | |____   _| |_   _| |_           | |__| | | |__| |    | |    | |      | |__| |    | |    
/_/    \_\ |_____/   \_____| |_____| |_____|           \____/   \____/     |_|    |_|       \____/     |_|    
                                                                                                              

```

## About

The ascii-output package generates ASCII art from a given string input using various banner templates. It provides functionalities to output the generated ASCII art to a file.The package supports generating ASCII art with different banner types (standard, shadow, thinkertoy) and writing the output to a specified file.



## Installation

To use the ascii-output package, ensure you have Go installed.

To clone and run this project locally, follow these steps:
```bash
git clone https://learn.zone01kisumu.ke/git/fnamayi/ascii-art-output
```
Navigate to Project Directory:
Change to the directory of the cloned repository:

```bash
cd ascii-output
```


## Usage

The package supports generating ASCII art with different banner types (`standard', shadow, thinkertoy) and writing the output to a specified file.
Command Line Interface

``` bash

go run . --output=<fileName.txt> [STRING] [BANNER]
```

* `--output=<fileName.txt>`: Specifies the output file name where the generated ASCII art will be saved.
* `[STRING]`: Input string for which ASCII art is to be generated.
* ``[BANNER]``: Optional banner type (standard, shadow, thinkertoy). Defaults to standard if not specified.

#### Example usage:

```bash
student$ go run . --output=banner.txt "hello" standard
student$ cat -e banner.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
student$ go run . --output=banner.txt 'Hello There!' shadow
student$ cat -e banner.txt
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
student$
```
This command generates ASCII art for the string "Hello" using the standard banner template and saves it to banner.txt.
Package Usage


## Functions
```bash 
GenerateASCII(input string, banner string) string{}
```
* `input (string)`: The input text that you want to convert into ASCII art.
* `banner (string)`: The style of ASCII art banner to generate. It specifies the file from which each line of ASCII art characters will be retrieved.

Generates ASCII art for the given input string using the specified banner template.
## WriteToOutputfile
```Bash
WriteToOutputFile(filename string, content string) error
```
Writes the generated ASCII art content to the specified filename.
```bash
GetLine(filename string, num int) string
```
Helper function to retrieve a specific line from the filename file based on num.
About

The ascii-output package is designed to simplify the generation and storage of ASCII art, providing flexibility with different banner styles. It is suitable for command-line tools or applications requiring text-based visualizations.
Contributors

    Arnold Adero

Authors

    Frankline Nyamayi

License

This project is licensed under the MIT License. See the LICENSE file for details.
