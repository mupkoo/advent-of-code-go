# Advent of Code Go

Welcome to the Advent of Code Go project! This repository contains the solutions for the Advent of Code (AoC) - an annual set of Christmas-themed programming challenges.

## Getting Started

Before you can start using this project, make sure you have Go installed on your system as this project assumes a Go development environment.

### Prerequisites

- Go (1.21 or higher is recommended)
- Advent of Code session cookie (to fetch puzzles and inputs)

### Setup

1. Clone the repository to your local machine:
   ```
   git clone https://github.com/mupkoo/advent-of-code-go.git
   ```
2. Navigate to the repository directory:
   ```
   cd advent-of-code-go
   ```

## Usage

This project comes with a Makefile which simplifies several common tasks:

### Show Help

To display a list of available make commands:
```
make help
```

### Install Dependencies

Make sure all necessary Go dependencies are installed:
```
go mod tidy
```

### Check AoC Session Cookie

Ensure your AoC session cookie environment variable is set. This is required to fetch the challenges and inputs automatically:
```
export AOC_SESSION_COOKIE=your-session-cookie
make check-aoc-cookie
```

### Generate Skeleton Code

Generate skeleton Go files for a specific day and year (if not provided, defaults to the current year). Replace `$DAY` and `$YEAR` with desired values:
```
make skeleton DAY=1 YEAR=2023
```

### Fetch Challenge Input

Download the input for a specific day and year. The AoC session cookie is required:
```
make input DAY=1 YEAR=2023
```

### Fetch Challenge Prompt

Retrieve the prompt for a specific challenge using:
```
make prompt DAY=1 YEAR=2023
```

Good luck and happy coding this Advent season!
