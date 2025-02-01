# GNU Core Utils

The GNU Core Utilites (coreutils) are a set of command line tools included by default in any Linux distribution, designed to perform basic tasks within the terminal, such as moving, viewing or editing files.

There are a number of other programs which fall into roughly the same category as the coreutils. For example diffutils, which includes `diff` and `cmp`, findutils which includes `find` and `xargs`, GNU grep, GNU tar, GNU less and util-linux. All of these programs will be covered here. 

## mv - coreutils

mv is used to move files/directories

`mv src dest` moves `src` to `dest`  
`mv -t dest src1 src2 src3 ... srcn` moves `src1`-`n` to directory `dest`. `-t` means "target directory"  
`mv -i` tells mv to ask for confirmation before moving files/folders

## cp - coreutils

cp is used to copy files/directories

`cp src dest` copies `src` to `dest`  
`cp -r src dest` copies `src` recursively to `dest`, meaning `src` will be copied along with it's contents/subdirectories  
`cp -t dest src1 src2 src3 ... srcn` moves `src1`-`n` to directory `dest`  
`cp -i` tell cp to ask for confirmation

## ls - coreutils

ls is used to list the files in the current directory, optinally including file attributes

`ls` - List files in current directory  
`ls dir` - List files in directory "dir"  
`ls -a` - List files including hidden files  
`ls -A` - List files including hidden files (not including . and ..)  
`ls -N` - Do not put quotes around files/directories with spaces in name  
`ls --group-directories-first` - List directories first  
`ls --file-type` - Add / to end of directory names  
`ls --color=auto` - Display different colours for different types of file (one for directories, one for standard files, one for symbolic links, one for executable files etc)  
`ls -l` - List files and display attributes  
`ls -lh` - List files and display attributes (human readable formats, i.e. 1.0K instead of 1024)  

## cat - coreutils

cat is used to output the contents of a file to stdout. It stands for "concatenate`, as it's intended purpose is to combine files together.

`cat file` will output the contents of `file` to stdout (the terminal)  
`cat file1 file2` will output the contents of `file1` followed by those of `file2` to stdout  

## wc - coreutils

wc is used to count the number of lines in a file/stdin

`wc -l file` - Show number of lines in file  
`wc -l *` - Show number of lines in all files in directory  
`cat file | wc -l` - Show number of lines in file (doesn't show filename)  
`cat file1 file2 | wc -l` - Show number of lines in file1 and file2  

## ln - coreutils

ln is used to link files

## mkdir - coreutils

mkdir is used to create directories

`mkdir dirname` - Create directory with name "dirname"  
`mkdir -p dirname/subdir` - Create directory "dirname" and subdirectory "subdir"

## rm - coreutils

rm is used to delete files/folders. **By default there is no recycle bin when using rm, files are gone forever**

`rm file` will delete `file`  
`rm -t` tells rm to ask for confirmation  
`rm -r folder` will delete `folder` recursively, meaning `folder` will be deleted along with all of it's contents  
`rm -f file` will delete without asking for confirmation. This is the default behaviour, however `-f` can be specified to overwrite `-i` if for example `rm` was aliased to `rm -i`

## rmdir - coreutils

rmdir is used to remove empty directories. Useful if you don't want to accidentally remove a directory with contents.

`rmdir dir` will delete `dir`

## chown - coreutils

## chmod - coreutils

## dd - coreutils

dd is used to copy data to a disk

`sudo dd input.iso /dev/sda` - Write input.iso to the drive /dev/sda **will overwrite everything on /dev/sda**  
`sudo dd input.iso /dev/sda --info=progress` - Show write progress, requires `progress` to be installed  
`sudo dd input.iso bs=1M /dev/sda` - Set block size to 1M  

## df - coreutils

df is used to view mounted filesystems and their total, used and avalible sizes, as well as mountpoints

## du - coreutils

du is used to list the sizes of files/directories. It defaults to displaying block sizes, and directories only

`du -b` - Display bytes instead of block sizes  
`du -h` - Use human readable values (i.e. 1K) instead of block sizes  
`du -a` - Display files and directories  

## tee - coreutils

## cur - coreutils

## tr - coreutils

## sort - coreutils

## uniq - coreutils

## head - coreutils

head is used to output the first n lines of a file

`head -n x file` - Display first x lines of file  

## tail - coreutils

tail is used to output the last n lines of a file, the opposite of head

`tail -n x file` - Display last x lines of file  

## sha1/244/256/384/512sum/md5sum - coreutils

The "sum" programs are used to get the hash of a file. **SHA1 is not recommended as it is compromised**

`sha256sum file` - Get the SHA256 sum of a file  
`md5sum file` - Get the MD5 sum of a file  

## diff - diffutils

## cmp - diffutils

## find - findutils

## xargs - findutils

## grep - GNU grep

## less - GNU less

less is a terminal pager, designed to output a file starting at the top, scrolling on pressing `Return`, and scrolling a page on pressing `Space`.

`less file` - Run less on file  

## tar - GNU tar

tar is used to make and extract tar archives.

`tar -xf file.tar.gz` - Extract the files from "file.tar.gz"

## cal - util-linux

cal is used to display a calendar in the terminal, including displaying the current date.

`cal` - Display the current month  
`cal -y` - Display the current year  
`cal -n 3` - Display the current month, and 2 months into the future  
By default, cal can include blank lines if a month only has 5 weeks. To fix this, cal can be aliased to `unbuffer cal | sed '/^ *$/d'`. This will remove the extra line while keeping the highlighting of the current day thanks to unbuffer. 

## lsblk - util-linux

lsblk is used to list drivers and partitions. 

`lsblk` - Show currently inserted drives, partitions and mountpoints  
`lsblk -l` - Output as a list, with no branches  
`lsblk -J` - Output as JSON  

## kill - util-linux

## pkill - procps-ng
