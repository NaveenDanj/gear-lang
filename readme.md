# Gear Lang

Gear Lang is a new programming language built using Go. It is designed to be simple and efficient, with a syntax inspired by other modern languages. This README provides an overview of the language and a sample program to help you get started.

## Features

- **Simple Syntax**: Gear Lang features a straightforward syntax that is easy to learn and use.
- **Strong Typing**: The language uses strong typing to help catch errors at compile time.
- **Functions**: Support for defining and calling functions with both integer and real number types.
- **Structures**: Define complex data types using structs.
- **Control Flow**: Support for conditional statements (`if`, `else`) and loops (`while`).
- **String Manipulation**: Concatenate and manipulate strings easily.

## Sample Program

Below is a sample program written in Gear Lang. This program demonstrates variable declarations, arithmetic operations, control flow statements, function definitions, and struct usage.

```gear
start

    import 'test2.ger';

    let int val1 = (10+11) + 23 + myVar;
    let int val2 = 20;
    let int out = (val1 + val2) + 100.45;
    let real r = 3.12;
    let string name = ('My name is naveen'+'hi') + 'Naveen Dhananjaya Hettiwaththa';
    
    if( (flag == true) || (flag == false) ) {
        print 'flag is true' + ' ' + name;
        print 'another print statement';
    }else{
        print 'another print statement';
        let int out = (val1 + val2)  + 200.9000;
    }

    print out;

    while (flag == true) {
        print 'flag is true';
    }

    let int newInt = val2;

    function int sum ( int val1, int val2) {
        let int out = val1 + val2; 
        print 'The sum is' + out;
        return out;
    }

    function real sum ( real val1, real val2) {
        let int out = val1 + val2; 
        print 'The sum is' + out;
        return out;
    }

    struct Person {
        name : string;
        age : int;
        greet1 : function void (string name){
            print 'Hello World ' + name;
            print 'This is message from greet1 function';
        }
        greet2 : function void (string name){
            print 'Hello World ' + name;
        }
        year:int;
    }

    let Person newPerson = 123;

end


## Getting Started

To get started with Gear Lang, follow these steps:

1. **Install Gear Lang**: Follow the installation instructions on our [official website](https://gear-lang.org).
2. **Write Your First Program**: Create a new file with the extension `.ger` and write your Gear Lang code.
3. **Run Your Program**: Use the Gear Lang compiler to compile and run your program.

## Documentation

For more detailed documentation, please visit our [documentation page](https://gear-lang.org/docs).

## Contributing

We welcome contributions from the community! If you'd like to contribute to Gear Lang, please read our [contributing guidelines](https://gear-lang.org/contributing).

## License

Gear Lang is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
