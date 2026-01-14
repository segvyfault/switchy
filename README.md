# switchy
simple wallpaper switcher

## Usage info

so you have like 3 files

- ``$HOME/.papers`` \- (must have that) file where you have all of the wallpaper paths, you can also pass args to matugen after the path
- ``$HOME/.paperactions`` \- (optional) file where you can store additional actions after the wallpaper has been set (for example reloading kitty themes)
- ``$HOME/.paper`` \- stores last used wallpaper path

the command syntax is pretty simple

1. ``switchy`` \- just reads the files and does the thing
4. ``switchy no-write`` \- doesn't overwrite old wallpaper, useful for wm init scenarios
2. ``switchy <path>`` \- use other .papers file 
3. ``echo "~/someimg.png" | switchy -`` \- dont use .papers file and just parse whatever is in stdin

## Goals
- [x] none i thibk
- [ ] oh i actually need to move those config files and data files to ~/.config/switchy thing
