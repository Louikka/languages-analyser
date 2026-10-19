# languages-analyser

Similar to [GitHub Linguist](https://github.com/github-linguist/linguist), this program analyses provided directory and all subdirectories and calculates percentage of used programming languages based on amount of characters (excluding whitespaces).


## Build

```sh
$ go build .
```


## Usage

```sh
$ lan [root_directory] [options]
```

`root_directory` is optional and defaults to the current directory. Use `--help` flag for more information.


## ToDo

- Add different metrics (such as lines of code, etc.).
- Add more languages.
