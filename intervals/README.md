# Sequences

## Goal

Create a program that accepts two inputs:

* a set of include intervals
* a set of exclude intervals

The sets of intervals can be given in any order
and they may be empty or overlapping.

The program should output the result of taking all of
the includes and remove/subtract the excludes.

The output should be given as non-overlapping intervals
in a sorted order.

Intervals will only contain integers.

### Examples

#### Example 1

Includes: 10-100

Excludes: 20-30

Output should be: 10-19, 31-100

#### Example 2

Includes: 50-5000, 10-100

Excludes: none

Output should be: 10-5000

#### Example 3

Includes: 10-100, 200-300

Excludes: 95-205

Output should be: 10-94, 206-300

#### Example 4

Includes: 10-100, 200-300, 400-500

Excludes: 95-205, 410-420

Output should be: 10-94, 206-300, 400-409, 421-500
