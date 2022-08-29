# DJ
DJ stands for Distributed Judge. A tool that many TAs wish they would have. This would provide more flexibility over normal judges such as [Quera](https://quera.ir)



## Badges

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](https://choosealicense.com/licenses/mit/)
![Languages count](https://img.shields.io/github/languages/count/rsharifnasab/DJ?style=for-the-badge)
![Lines count](https://img.shields.io/tokei/lines/github/rsharifnasab/DJ?style=for-the-badge)
![Commit Freq](https://img.shields.io/github/commit-activity/w/rsharifnasab/DJ?style=for-the-badge)

![CI status](https://github.com/rsharifnasab/DJ/actions/workflows/test.yml/badge.svg?)

## Features

- Judge student codes locally
- Most flexible judge out there
- Check out source codes against illegal usage of language constructs.
- Doesn't need server


## Acknowledgements

 - [Dr Mojtaba Vahidi, My Kind Supervisor](http://facultymembers.sbu.ac.ir/vahidi/)
 - [Dr Sadegh Aliakbari, My Referee](http://facultymembers.sbu.ac.ir/aliakbary/)
 - [Amir Arsalan San'ati, Who gave me the idea](https://github.com/Amirarsalan-sn)
 


## Demo

This is a demo which demonstrate how students can judge their code locally.

[![demo](https://asciinema.org/a/335480.svg)](https://asciinema.org/a/B7EEbzwsnDVGq7pFu012wm5UM?autoplay=1)
## Installation

Official way to use this repository is clone the repository and run distribute script

You can also install the binary file like this (which is not recommended because it does not contain examples)
```bash
  go install "github.com/rsharifnasab/DJ@latest"
```

And then use binary file like this:
  
```bash
  DJ --help 
```
    
## Distribute for students

To Distribute the project run

```bash
  ./scripts/dsitribute.sh
```

Then the distribution zip would be in `./bin/` directory.


## Running Tests

To run tests, run the following command

```bash
  go test ./... -cover  -count 1
```


## Documentation

Exploration on various software and library choices is Documented [here](https://github.com/rsharifnasab/DJ/tree/master/docs)


## FAQ

#### Do I need this as a student?

No you don't. The course TA should clone this repository and create questions and then distribute questions alongside the judge and the binary for you.

#### Should I learn Go before using this project?

No you shouldn't. The flexible part is not written in Go, but instead you need a bit of bash script to tune judges and develop creative questions.




## Used By

This project is used by the following courses:

- Shahid Beheshti university, Advanced programming (WIP at that time)
- Sharif University, Compiler course for CE students



## Related

Here are some related projects

- [TA utils](https://github.com/rsharifnasab/ta_utils): another repository for TAs which contains small and useful scripts.
- [Sharif Judge](https://github.com/mjnaderi/Sharif-Judge): A free and open source online judge system for programming courses

## Roadmap

- Better support for Windows
- Provide learning resources for Bash script
- Better API for source code checks
- Save student scores on a Block Chain


## License

[MIT](https://choosealicense.com/licenses/mit/)


## Support

For support, email rsharifnasab@gmail.com.

