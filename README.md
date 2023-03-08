# DJ
DJ stands for Distributed Judge. A tool that many TAs wish they could have. This would provide more flexibility over normal judges such as [Quera](https://quera.ir) which is used by many Iranian universities

<!-- TODO: Add an example of a non-Persian judge -->



## Badges

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](https://choosealicense.com/licenses/mit/)
![Languages count](https://img.shields.io/github/languages/count/rsharifnasab/DJ?style=for-the-badge)
![Commit Freq](https://img.shields.io/github/commit-activity/w/rsharifnasab/DJ?style=for-the-badge)

![CI status](https://github.com/rsharifnasab/DJ/actions/workflows/test.yml/badge.svg?)

## Features

- Judges student codes locally
- Flexible 
- Checks out source codes against illegal usage of language constructs.
- Doesn't need a server


## Acknowledgements

 - [Dr Mojtaba Vahidi, My Kind Supervisor](http://facultymembers.sbu.ac.ir/vahidi/)
 - [Dr Sadegh Aliakbari, My Referee](http://facultymembers.sbu.ac.ir/aliakbary/)
 - [Amir Arsalan San'ati, Who gave me the idea](https://github.com/Amirarsalan-sn)
 


## Demo

This is a demo which demonstrate how students can judge their codes locally.

[![demo](https://asciinema.org/a/335480.svg)](https://asciinema.org/a/B7EEbzwsnDVGq7pFu012wm5UM?autoplay=1)
## Installation

### Prerequisites
+ `Golang` should be installed on your system
+ Also there is heavy usage of `Bash`
+ For running tests you should have `python3` and `gcc` and `java` in your `PATH`


### Installation

Official way to use this repository is to clone the repository and run the "Distribute script".

However, you can also install with `go install` and providing a git repository:
```bash
$ go install "github.com/rsharifnasab/DJ@latest"
```

And after that, use the binary file:
  
```bash
$ DJ --help 
```

But this method is not recommended as it won't have the examples contained.
    
## Distribute for students

To distribute the project run

```bash
./scripts/distribute.sh
```

Then the distribution ZIP would be in `./bin/`.


## Running Tests

To run tests, run the following command

```bash
$ go test ./... -cover  -count 1
```


## Documentation

Exploration on various software and library choices is Documented [here](https://github.com/rsharifnasab/DJ/tree/master/docs)


## FAQ

#### Do I need this as a student?

No you don't. The course TA should clone this repository and create questions and then distribute questions alongside the judge and the binary for you.

#### Do I Have to learn Go before using this project?

No you don't. The flexible part is not programmed in Go, but instead you need a bit of bash script knowledge to tune judges and develop creative questions.

## Used By

This project is used by the following courses:

- Shahid Beheshti university, Advanced programming (WIP at the time of writing this)
- Sharif University, Compiler course for CE students

## Related

Here are some related projects

- [TA utils](https://github.com/rsharifnasab/ta_utils): another repository for TAs which contains small and useful scripts.
- [Sharif Judge](https://github.com/mjnaderi/Sharif-Judge): A free and open source online judge system for programming courses

## Roadmap

- Better Windows support
- Provide learning resources for Bash scripting
- Better API for source code checks
- Save student scores on a BlockChain


## License

[MIT](https://choosealicense.com/licenses/mit/)


## Support

For support, email rsharifnasab on gmail.com.
