# switchy
simple wallpaper switcher

## Usage info

so you have like 2 files  
one of them is swichy's internal usage one

- ``$HOME/.config/switchy/papers`` \- (must have that) file where you have all of the wallpaper paths, you can also pass args to matugen after the path
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
4. ``go build . && strip switchy`` (strip command to reduce the executable size)
