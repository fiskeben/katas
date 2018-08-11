# Banko Plates

## Background

A (Danish?) variant of the game Bingo, called "Banko,"
is played with six 9x3 cards (three rows of nine fields)
where each row has five numbers from 1 to 90 (both included).
Each column represents groups of tens, so that the first column
holds numbers from 1 to 9, the next 10 to 19, and so on.
The number 90 goes into the last column.

## Goal

Create a program that can generate a random board of six banko cards so that:

* all numbers from 1 to 90 are used
* each card has three rows of five numbers
* each row has maximum 1 number per group of tens
  (so, if 12 is used, no other number from 10 to 19 is allowed)
* new cards are generated every time the program runs.

The program should be able to visualize the cards.
This means it should include blanks and not just list the numbers.

## Example Card

A card could look like this:

    ----------------------------
    |  |  |27|33|48|  |66|76|  |
    | 6|14|  |37|  |  |  |73|83|
    |  |18|22|34|  |55|63|  |  |
    ----------------------------
