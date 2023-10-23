# Word Counter - Word, Line, and Byte Counter

gowc is a command-line utility written in Go that provides word, line, and byte counting functionality, inspired by the Unix `wc` command. It can count the number of lines, words, and bytes in a given file or standard input.

## Installation

To use gowc, you need to have Go installed on your system.

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/gowc.git

Change to the directory where the code is located:
```bash
cd gowc
```

```bash
go build -o gowc.exe
```

# Usage: gowc [options] [filename]

Options:\
-c  Count bytes\
-l  Count lines\
-w  Count words\
-m  Count Characters

If no options are provided, the program defaults to counting bytes, lines,characters and words.

If a filename is not provided, the program will read from standard input.

Examples:
- Count bytes, lines, and words in a file:
  ```
  gowc -c -l -w filename.txt
  ```

- Count lines and words from standard input (use Ctrl+D to exit):
  ```
  cat filename.txt | gowc -l -w
  ```

- Count bytes, lines, and words using the default behavior:
  ```
  gowc filename.txt
  ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments
- This project was inspired by the Unix `wc` command.

Enjoy using gowc! If you have any feedback or suggestions, please feel free to open an issue or contribute to the project.

