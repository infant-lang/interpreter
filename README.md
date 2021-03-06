# Infant Lang Interpreter

## Change Log

- ### `1.0.3` (2022-01-15)
    - Added new function which checks for parsing error
    - Now prints error message if parsing error occurs which was not available before thereby making debugging difficult
    - Optimized error handling of Grammar matching
    - Feature request of `Sandbox` API

- ### `1.0.2` (2022-01-14)
  - Utility function will save it to global message variables.
  - Utility function can no longer print to the `stdout`.
  - Added support for the `sandbox` API
  - Lexer now supports `#` comments
  - Optimized handling of Tokenization Error by the Lexer

- ### `1.0.1` - (2021-12-31)
  - Initial release
  - Interpreter uses `stdin` for getting the `file` to interpret
  - Individual functions can print to the `stdout`

# Infant Lang
Minimalistic Less Esoteric Programming Language for Infants

## How did we get here? ð
### Inspired by BrainFuck but does things in a less esoteric way. ð
1. More elegant keywords for infants
2. Manual simplified control over the pointer
3. Conditional implementation with the simple âifâ
4. Memory for storing temporary data and manipulating it
5. Easiest one line for loop for repetitive tasks such as multiplication

### What can it achieve? ð¤ªð¥³
We <3 developers but love infants more. Infant lang is developed to promote infants to start coding with the stuff they can visualize like moving their toy along a line and storing their toys in a box. With the upcoming support for threads, Infant Lang helps infants understand how multitasking works inside computers and create wonderful things.

## Pointer ð and Memory ð¦
- One dimensional array of size 35000 is available all initialized from 0 to 35000.
- The Pointer is initialized to point the 0 box
- A single memory block is available to store an integer initialized to zero.
- Move Pointers with easy commands or with loops.

## Pointer ð Movement âµ
### Move the pointer with commands so easy that an infant can do.
- `# Moves the Pointer left by 1 box`
    - `move pointer left`
- `# Moves the Pointer right by 1 box`
    - `move pointer right`
- `# Moves the Pointer Left 5 boxes`
    - `move pointer left 5`
- `# Moves the Pointer Right 10 boxes`
    - `move pointer right 10`

### Info About the Pointer
- Default takes 1 but can enter number to move pointers to more than 1
- Use a loop if you wanted to move the pointer repeatedlyâ¦
    - `# Moves the Pointer right 15 times with a for loop`
        - `for 15 move pointer right`

## The print ð¨ï¸
- Print stuff to the screen.
- Print the pointer value or value in the memory or any number.
- Want to print characters? Use `char` to print the ASCII.
- Print a tab, new line or just print a single space all with ease.

### Printing ð¨ï¸ Methods
- Print stuff to the output screen with the 'print' keyword. Defaults to a single space.
- `# Prints the pointer value`
    - `print pointer`
- `# Prints the ascii value of the number in the memory`
    - `print char memory`
- `# Prints a number`
    - `print 15`

- Print a space, tab or a new line with the below keywords `space`, `tab` or `new`
- `# Prints a space`
    - `print space`
- `# Prints a tab`
    - `print tab`
- `# Prints a new line`
    - `print new`

## The Memory ð¦
- Store a number into memory for later use â
- `# Stores a number in the memory`
    - `memory = 13`
- `# Stores the pointer value inside the memory`
    - `memory = pointer`
- Memory is always initialized to zero for null safety. It can store only integers. No decimals. ð

## Arithmetic ââââ
- Infant lang supports basic arithmetic of integers.
- Throws error if a division returns a decimal ð
- Throws error if pointer or memory becomes negative ð
- Division by zero is still not possible ð

### Basic all possible arithmetic ð§ 
- Addition, subtraction, multiplication and division of two numbers, with memory or with the pointer and store it in the memory ð¦
- All Possible Arithmetic Operations
    - `# Addition`
        - `memory = memory + pointer`
        - `memory = memory + memory`
        - `memory = memory + 15`
    - `# Subtraction`
        - `memory = memory - pointer`
        - `memory = memory - memory`
        - `memory = memory - 12`
    - `# Multiplication`
        - `memory = memory * pointer`
        - `memory = memory * memory`
        - `memory = memory * 8`
    - `# Division`
        - `memory = memory / pointer`
        - `memory = memory / memory`
        - `memory = memory / 2`

## The âifâ â
- Do simple stuff with the `if` keyword.
- No `else` block to confuse the infant.
- Simple arithmetic conditions works. ð

### Do more with the `if` ð±âð
Use the âifâ to do conditional stuff like checking if the memory is 79 or the pointer is 320. Give it a try!! But no negative integer thouâ¦ ð
- `if memory == pointer `
- `if pointer == memory `
- `if memory == 13`
- `if pointer == 13`

It also supports logical checking with AND and OR

- `if memory && pointer print memory`
- `if pointer && memory print tab`
- `if memory || pointer print new`
- `if pointer || memory print char pointer`

## The âforâ ðââï¸ð¨
- Do repetitive stuff with the for.
- No variable initialization. 
- Enter a positive integer or pass the pointer ð or memory ð¦
- Heard of single line for? Bet you havenât. It's gonna blow your mind ð¤¯

### Repetitive stuff is easy with the âforâ ðââï¸ð¨
The easiest âforâ implementation that doesnât allow a negative number. Move pointer faster with the single line loops. Pass a number or memory or pointer.

- `for 15 print char memory`
- `for pointer print 23`
- `for memory print memory`

- `# Move pointer twice`
    - `for 2 move pointer right`

- `# Pass an if in the line to do crazy stuff`
    - `for 15 if memory != pointer print char 15`


## How Infant Lang Works ð¤
1. Infant Lang is an interpreter inspired by BrainFuck.
2. Written in GoLang.
3. It reads the code line by line.
4. It instantiates two variables:
    - `pointer` to keep track of the pointer position
    - `memory` is initialized to zero.
5. A Lexer splits the code into tokens.
6. The Lexer then checks if the token is a keyword or a number.
7. The Lexer sends the token to the Parser.
8. Parser parses the tokens and executes the code.

## Installation ð¦ð½
The Executable is available only for amd64 Windows Systems.

- Download the zip file from the Releases page.
- Unzip the file 
- Open Run and Type `%USERPROFILE%` and press `OK`
- Create a folder `bin` if not already created.
- Extract the executable to the `bin` folder.
- Open Command Prompt and type the following command:
    - `setx INFANT_HOME "%USERPROFILE%\bin\infant"`
    - `setx PATH "%INFANT_HOME%;%PATH%"`
- Close the Command Prompt.
- Open a new Command Prompt and type `infant -v` or `infant --version` to check if the installation was successful.
- Now pass any file ending with `.infant` to the executable.
- Eg: In the Command Prompt type `infant hello.infant`

## For Future Use 
- Start a Go Project `go mod init github.com/USERNAME/project`

## Build
- `go build .`

## Arguments
- `filename`: The file to be interpreted
- Eg: `infant.exe test.txt`