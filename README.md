# Infant Lang Interpreter


# Infant Lang
Minimalistic Less Esoteric Programming Language for Infants

## How did we get here? 👑
### Inspired by BrainFuck but does things in a less esoteric way. 😜
1. More elegant keywords for infants
2. Manual simplified control over the pointer
3. Conditional implementation with the simple ‘if’
4. Memory for storing temporary data and manipulating it
5. Easiest one line for loop for repetitive tasks such as multiplication

### What can it achieve? 🤪🥳
We <3 developers but love infants more. Infant lang is developed to promote infants to start coding with the stuff they can visualize like moving their toy along a line and storing their toys in a box. With the upcoming support for threads, Infant Lang helps infants understand how multitasking works inside computers and create wonderful things.

## Pointer 👆 and Memory 📦
- One dimensional array of size 35000 is available all initialized from 0 to 35000.
- The Pointer is initialized to point the 0 box
- A single memory block is available to store an integer initialized to zero.
- Move Pointers with easy commands or with loops.

## Pointer 👆 Movement ⛵
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
- Use a loop if you wanted to move the pointer repeatedly…
    - `# Moves the Pointer right 15 times with a for loop`
        - `for 15 move pointer right`

## The print 🖨️
- Print stuff to the screen.
- Print the pointer value or value in the memory or any number.
- Want to print characters? Use `char` to print the ASCII.
- Print a tab, new line or just print a single space all with ease.

### Printing 🖨️ Methods
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

## The Memory 📦
- Store a number into memory for later use ⌚
- `# Stores a number in the memory`
    - `memory = 13`
- `# Stores the pointer value inside the memory`
    - `memory = pointer`
- Memory is always initialized to zero for null safety. It can store only integers. No decimals. 💀

## Arithmetic ➕➖✖➗
- Infant lang supports basic arithmetic of integers.
- Throws error if a division returns a decimal 💀
- Throws error if pointer or memory becomes negative 💀
- Division by zero is still not possible 😒

### Basic all possible arithmetic 🧠
- Addition, subtraction, multiplication and division of two numbers, with memory or with the pointer and store it in the memory 📦
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

## The ‘if’ ❓
- Do simple stuff with the `if` keyword.
- No `else` block to confuse the infant.
- Simple arithmetic conditions works. 😒

### Do more with the `if` 🐱‍🏍
Use the ‘if’ to do conditional stuff like checking if the memory is 79 or the pointer is 320. Give it a try!! But no negative integer thou… 😖
- `if memory == pointer `
- `if pointer == memory `
- `if memory == 13`
- `if pointer == 13`

It also supports logical checking with AND and OR

- `if memory && pointer print memory`
- `if pointer && memory print tab`
- `if memory || pointer print new`
- `if pointer || memory print char pointer`

## The ‘for’ 🏃‍♂️💨
- Do repetitive stuff with the for.
- No variable initialization. 
- Enter a positive integer or pass the pointer 👆 or memory 📦
- Heard of single line for? Bet you haven’t. It's gonna blow your mind 🤯

### Repetitive stuff is easy with the ‘for’ 🏃‍♀️💨
The easiest ‘for’ implementation that doesn’t allow a negative number. Move pointer faster with the single line loops. Pass a number or memory or pointer.

- `for 15 print char memory`
- `for pointer print 23`
- `for memory print memory`

- `# Move pointer twice`
    - `for 2 move pointer right`

- `# Pass an if in the line to do crazy stuff`
    - `for 15 if memory != pointer print char 15`


## How Infant Lang Works 🤔
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

## Installation 📦🔽
The Executable is available only for amd64 Windows Systems.

- Download the zip file from the Releases page.
- Unzip the file 
- Open Run and Type `%USERPROFILE%` and press `OK`
- Create a folder `bin` if not already created.
- Extract the executable to the `bin` folder.
- Open Command Prompt and type the following command:
    - `setx INFANT_HOME "%USERPROFILE%\bin\infant\"`
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