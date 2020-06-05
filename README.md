[![Build Status](https://travis-ci.org/Kalimaha/toyrobot-go.svg?branch=master)](https://travis-ci.org/Kalimaha/toyrobot-go)
[![Coverage Status](https://coveralls.io/repos/github/Kalimaha/toyrobot-go/badge.svg?branch=setup-travis)](https://coveralls.io/github/Kalimaha/toyrobot-go?branch=setup-travis)

# Toy Robot

# Test, install and run

## Test

### With Docker

Build the container with `docker build -t toy-robot-go .`, then:

```
docker run -it toy-robot-go ./simon.says test
```

### Without Docker

```
./simon.says test
```

## Install

Installation is managed through Maven with:

TODO

## Run

TODO

# Description

- The application is a simulation of a toy robot moving on a square tabletop, of dimensions 5 units x 5 units.
- There are no other obstructions on the table surface.
- The robot is free to roam around the surface of the table, but must be prevented from falling to destruction. Any movement that would result in the robot falling from the table must be prevented, however further valid movement commands must still be allowed.

Create an application that can read in commands of the following form:

```
PLACE X,Y,F
MOVE
LEFT
RIGHT
REPORT
```

- [ ] PLACE will put the toy robot on the table in position `X`, `Y` and facing `NORTH`, `SOUTH`, `EAST` or `WEST`.
- [X] The origin `(0,  0)` can be considered to be the `SOUTH WEST` most corner.
- [ ] The first valid command to the robot is a PLACE command, after that, any sequence of commands may be issued, in any order, including another PLACE command. The application should discard all commands in the sequence until a valid PLACE command has been executed.
- [X] MOVE will move the toy robot one unit forward in the direction it is currently facing.
- [X] LEFT and RIGHT will rotate the robot 90 degrees in the specified direction without changing the position of the robot.
- [X] REPORT will announce the X,Y and F of the robot. This can be in any form, but standard output is sufficient.
- [X] A robot that is not on the table can choose the ignore the MOVE, LEFT, RIGHT and REPORT commands.
- [ ] Input can be from a file, or from standard input, as the developer chooses.
- [X] Provide test data to exercise the application.

## Constraints

- [X] The toy robot must not fall off the table during movement. This also includes the initial placement of the toy robot.
- [X] Any move that would cause the robot to fall must be ignored.

## Example Input and Output

### Example a

```
PLACE 0,0,NORTH
MOVE
REPORT
```

Expected output:

```
0,1,NORTH
```

### Example b

```
PLACE 0,0,NORTH
LEFT
REPORT
```

Expected output:

```
0,0,WEST
```

### Example c

```
PLACE 1,2,EAST
MOVE
MOVE
LEFT
MOVE
REPORT
```

Expected output

```
3,3,NORTH
```

## Extensions

Now that you have a working solution to the problem above, let's extend it.

### Extension 1: Obstacles

Extend the Toy Robot with a `PLACE_OBJECT` command, to place an object on the tabletop surface. When moving the toy robot it should be prevented from bumping into previously placed objects on the tabletop surface. `PLACE_OBJECT` should place an object in front of the current location of the toy robot. For example if the toy robot is at location `(0, 0)` and facing `EAST`, an object should be placed in location `(1, 0)`.

### Extension 2: ASCII visualisation

Extend the Toy Robot with a `MAP` command, to print an ASCII map of the current tabletop. The ASCII map of the table top should indicate previously placed objects using the `X` character.

#### Example

Given a tabletop with objects at locations `(0, 1)` and `(2, 2)`, `MAP` should print:

```
00000
00000
00X00
00000
0X000
```
