# switchy
simple wallpaper switcher

## Usage info

so you have like 3 files

- ``$HOME/.config/switchy/papers`` \- (must have that) file where you have all of the wallpaper paths, you can also pass args to matugen after the path
- ``$HOME/.config/switchy/paperactions`` \- (optional) file where you can store additional actions after the wallpaper has been set (for example reloading kitty themes)
- ``$HOME/.config/switchy/paper`` \- stores last used wallpaper path

the command syntax is pretty simple

1. ``switchy`` \- just reads the files and does the thing
4. ``switchy no-write`` \- doesn't overwrite old wallpaper, useful for wm init scenarios
2. ``switchy <path>`` \- use other .papers file 
3. ``echo "~/someimg.png" | switchy -`` \- dont use .papers file and just parse whatever is in stdin

## Goals
- [x] none i thibk
- [x] move config garbage to ~/.config/switchy 

## How to compile

1. ``git clone https://github.com/segvyfault/switchy``
2. ``cd switchy``
3. ``go mod tidy``
4. ``go build -o switchy -ldflags "-s -w" main.go`` (the ldflag thing is to reduce file size, you can use strip command instead if you want)
