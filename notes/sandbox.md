# sandbox options

-   Multiple secure, isolated Linux containers [+](https://openvz.org/)

-   Jail Chroot Project is an attempt of write a tool that builds a chrooted environment. The main goal of Jail is to be as simple as possible, and highly portable [+](http://www.jmcresearch.com/projects/jail/)

-   LXC and LXD: image based sandbox [+](https://linuxcontainers.org/)

-   chroot

-  isolate (used in MOE and CMS) [+](https://github.com/ioi/isolate) and [+](http://www.ucw.cz/moe/isolate.1.html)

-  compilebox (docker based but with js!) [+](https://github.com/remoteinterview/compilebox)
[introduced here](https://www.linkedin.com/pulse/how-does-online-judge-works-ahmad-faiyaz/)


-  vagrant:available for windows too, work with virtualbox backend
[+](https://www.vagrantup.com/downloads)


-  firecracker: VM but really fast to boot [+](https://jvns.ca/blog/2021/01/23/firecracker--start-a-vm-in-less-than-a-second/)

worth read: [+](http://coldattic.info/post/40/) 
and this perl script for timeout and mem usage [+](https://github.com/pshved/timeout)

---

# lxd

## install

```bash
pacman -S lxd
sudo systemctl start lxd.service
sudo lxd init
lxd init
```

## config

-   cluster: no
-   storage : yes (default, backend=dir)
-   network: yes (default)

## commands

```bash
# download new emage and start conatiner
lxc launch ubuntu:20.04 ubuntuone

# start
lxc start ubuntuone

# run command
lxc exec ubuntuone -- bash


```

## to run without sudo

```bash
sudo usermod -a -G lxd <username>
sudo newgrp lxd
```

## note

-   storage pool is mandatory
-   use `dir` back-end for storage pool
-   add to `/etc/default/lxc`

```
lxc.idmap = u 0 100000 65536
lxc.idmap = g 0 100000 65536
```

-   create subuid and subgid files

```
# in /etc/subuid
root:1000000:65536
roozbeh:1000000:65536
# in /etc/subgid
root:1000000:65536
roozbeh:1000000:65536
```
+ storage pool is here `/var/lib/lxd/storage-pools`

## cheetsheat
+ [github gist](https://gist.github.com/berndbausch/a6835150c7a26c88048763c0bd739be6) 

+ [publish image](https://ubuntu.com/blog/publishing-lxd-images)

+ [limiting resources](https://www.maketecheasier.com/limit-lxd-containers-resources/)

