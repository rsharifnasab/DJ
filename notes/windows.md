windows need to run bash scripts
or what?!

-   batsh: compile some new langauge to both bash and batch [+](https://github.com/batsh-dev-team/Batsh)

    -   discontinued
    -   need to compile from ocaml source

-   mingw

    -   only gcc, g++, not bash
    -   previously installed in ITP course
    -   a bit hard to install

-   cygwin

    -   good environment
    -   enough coreutils commands
    -   no gcc, g++: can be installed with cyg-get
    -   gcc can be install with installer exe (maybe stressing) [+](https://gist.github.com/miguelmota/e7d54775e5ded24dfe3a9b35327977be)
    -   "bash" is not already in path, some other executable is.
    -   add whole bin folder to path? polute the PATH



-   msys2
    - arch based 
    - they need to install packages (gcc, g++ and more) with pacman -Syu 
    - I think needs really much internet and time.


-   tools
    -   choco: really good installer install with one command (need powershell as administrator)
    -   cygwin installer script
        [+](https://github.com/miguelgrinberg/cygwin-installer/blob/master/install-cygwin.bat) and
        [+](https://github.com/rtwolf/cygwin-auto-install/blob/master/cygwin-install.bat) and
        [+](https://github.com/vegardit/cygwin-portable-installer)
    - need dos2unix for running scripts :) could be installed with cygwin
    
    - [+](add sth to path)




## draft HOW-TO install

1. install chocolaty
  + you can install from [here](https://chocolatey.org/install)
  + tldr: open powershell as administrator and run this
```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

2. install needed software with chocolaty
  + open another powershell as administrator 
  + choco install cygwin cyg-get # chocolateygui? vscode? mingw?
  + press `A` to accept all


3. add cygwin bin folder to PATH
  + find cygwin installation foler (the hard part, maybe I distribute cygwin myself like mingw)
  + copy path for example `C:\cygwin64\bin\`
  + search `edit system environment variables` and open it
  + look into the second half (System variables part)
  + open `PATH` for edit
  + add copied path at the end:
    windows 10: create new entry
    other windows: enter a ";" and paste after that


4. install needed software with cyg-get
  + open normal cmd 
  + run this command:
```batch
cyg-get install dos2unix g++ 
```


5. test compiler and bash in cmd
  + run `g++ --version` and see the output
  + run `bash -c "g++ --version"`

6. test all together

+ create a cpp hello-world file named `a.cpp`
+ put this in a file named `a.bash`

```bash
g++ "$1" -o "a.exe"
./a.exe
```

+ open a cmd in same directory as `a.bash` and `a.cpp`
+ run this

```batch
dos2unix a.bash 
bash a.bash a.cpp
```
