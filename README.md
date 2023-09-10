# ascii-art-color

### Usage

Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII

```console
Usage: 
go run . [STRING] [BANNER]
go run . [STRING]
```

*** !!Carefully read about using the options!! ***

### Flags

#### reverse flag is used to convert ascii-art from file to normal text

```console
Example of usage reverse option:

Usage: go run . [OPTION]
Ex: go run . --reverse=<filename>
```

All cases with reverse flag:
```console
go run . --reverse=<filename> (banner)?
go run . --reverse=<filename> --output=<filename> (banner)?
go run . --reverse=<filename> (justify flag)? --color=<color> (banner)?
go run . --reverse=<filename> (justify flag)? --color=<color> <letters to be colored> (banner)?
```

#### color flag colorize text in terminal

To change the color of the output it must be possible to use a flag `--color=<type>`, in which type can be :
 - "red", "green", "blue", "yellow", "orange", "purple", "cyan", "pink", "gray", "black", "white"
 - rgb (Ex.: "rgb(255,0,0)")
 - hex (Ex.: "#ff01a12")

color flag can be provided with optional string with letters, that programm should color

```console
Example color flag: 
go run . [OPTIONS] --color="<color>" <letters to be colored> <text> <banner>
go run . [OPTIONS] --color="<color>" <letters to be colored> <text>
go run . [OPTIONS] --color="<color>" <text>

You cannot use!!!
go run . [OPTIONS] --color="<color>"  <text> <banner>

Example:
go run . --color="rgb(255,0,0)" "name" "Hi, my name is Rick"
```

#### align flag change the alinment of the output

To change the alignment of the output it must be possible to use a flag `--align=<type>`, in which type can be :
 - center
 - left
 - right
 - justify

#### output flag use to write ascii art to txt file

To write ascii art to file it must be possible to use a flag `--output=<filename.txt>`

 - ***you cannot use output option with align and color flags***

## Our team

 - [eismagulo](https://01.alem.school/git/eismagulo)

 - [mkabyken](https://01.alem.school/git/mkabyken)