# GoPurpleAmerica
Solution for a nifty assignment, Purple America

The [link](http://nifty.stanford.edu/2014/wayne-purple-america/) to the assignment

The assignment provides with USA states points data and election results
and asks to generate an image showing results of an election by coloring
the states or counties to the according colors

# Usage
```
purple [options] -r STATE -rd STATES_DATA.zip -o output.svg
Options:
	-r/--region REGION_NAME		Selects region to draw (USA by default)
	-rd/--regions-data REGIONS.zip	Selects archive containing region points data
	-d/--data PATH.zip		Selects archive containing statisics data
	-y/--year INT			Suffix of the state statistics file name
	-n/--colors FILE_PATH		Draws with colors presented in the file
	-N OUTPUT			Saves an example of a color file
	-h/--help			Shows this message
	-o/--output OUTPUT.svg		Output path
	-s/--scale FLOAT		Scales the result by the given factor (10 by default)
	-sw/--stroke-width FLOAT	Sets stroke width
	-sc/--stroke-color R,G,B,A	Sets stroke color
```

# Examples
> purple -rd regions.zip -r CA -o output.svg --scale 10

![Result](https://github.com/MrPythoneer/GoPurpleAmerica/raw/main/data/example.png)
___
> purple -rd regions.zip -r CA -o output.svg -d elections.zip -y 2004 --scale 10

![Result](https://github.com/MrPythoneer/GoPurpleAmerica/raw/main/data/example_colored.png)
