# Mini_chain

A simple blockchain implementation in Go for learning purposes.

## Features

- Basic blockchain structure
- Simple Proof of Work mining
- CLI interface
- Concurrent block processing
- Basic validation


## Architecture

The project is structured into several key components:

### Blockchain
- Handles the chain of blocks
- Manages block validation
- Implements mining logic

### Types
- Define core data structures
- Implement block structure
- Defines interfaces

### Mempool
- Manages pending transactions/data
- Implements data queue system


## Usage


To build and run the project, use the following command:

```
make build
./mini_chain
```

Output:

```
 ./mini_chain
Enter a comand (add/list/validate/quit):
add
Enter data for the new block: new block
Block added successfully
Enter a comand (add/list/validate/quit):
list
Block #0 [Hash: b863f95fabd6b003017efb494f211ef723d045d77981aa62bed2c7e80df43111]
Data: Genesis Block

Block #1 [Hash: 0c6ce237f23c25061829837df5530d954def4306dd7c807b37b9c6d359fa5111]
Data: new block

Enter a comand (add/list/validate/quit):
validate
Blockchain is valid
Enter a comand (add/list/validate/quit):
quit
Exiting the program
```

