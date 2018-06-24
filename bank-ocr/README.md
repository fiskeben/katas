# Bank OCR

From [JoeJag](https://code.joejag.com/coding-dojo/bank-ocr/])

This assignment consists of four user stories
but I will only implement the first three
because the fourth is just guesswork in my opinion.

Follow the link above to see more about what it is.

## Introduction

Your manager has recently purchased a machine that assists in reading letters and faxes 
sent in by branch offices. 
The machine scans the paper documents, and produces a file with a number of entries. 
You will write a program to parse this file.

## Specification

### User Story 1

The following format is created by the machine:

````
    _  _     _  _  _  _  _
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|

````

Each entry is 4 lines long, and each line has 27 characters. 
The first 3 lines of each entry contain an account number written using pipes and 
underscores, and the fourth line is blank.

Each account number should have 9 digits, all of which should be in the range 0-9. 
A normal file contains around 500 entries.

Write a program that can take this file and parse it into actual account numbers.

### User Story 2

You find the machine sometimes goes wrong while scanning. 
You will need to validate that the numbers are valid account numbers using a checksum. 
This can be calculated as follows:

````
account number:  3  4  5  8  8  2  8  6  5
position names:  d9 d8 d7 d6 d5 d4 d3 d2 d1

checksum calculation:
((1*d1) + (2*d2) + (3*d3) + ... + (9*d9)) mod 11 == 0
````

### User Story 3

Your boss is keen to see your results. 
He asks you to write out a file of your findings, one for each input file, 
in this format:

````
457508000
664371495 ERR
86110??36 ILL
````

The output file has one account number per row. 
If some characters are illegible, they are replaced by a `?`. 
In the case of a wrong checksum, or illegible number, 
this is noted in a second column indicating status.

