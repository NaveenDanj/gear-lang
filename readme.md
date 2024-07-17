# Gear Lang

Gear Lang is a new programming language built using Go. It is designed to be simple and efficient, with a syntax inspired by other modern languages. This README provides an overview of the language and a sample program to help you get started.

## Features

- **Simple Syntax**: Gear Lang features a straightforward syntax that is easy to learn and use.
    ```gear
    start
        print 'Hello, World!';
    end
    ```

- **Strong Typing**: The language uses strong typing to help catch errors at compile time.
    ```gear
    start
        let int number = 10;
        let string text = 'Hello';
    end
    ```

- **Functions**: Support for defining and calling functions with both integer and real number types.
    ```gear
    start
        function int add(int a, int b) {
            return a + b;
        }

        function real multiply(real x, real y) {
            return x * y;
        }

        let int sum = add(5, 10);
        let real product = multiply(2.5, 4.0);
    end
    ```

- **Structures**: Define complex data types using structs.
    ```gear
    start
        struct Person {
            name : string;
            age : int;
        }

        let Person john = Person {
            name = 'John Doe',
            age = 30
        };
    end
    ```

- **Control Flow**: Support for conditional statements (`if`, `else`) and loops (`while`).
    ```gear
    start
        let int count = 5;

        if (count > 0) {
            print 'Count is positive';
        } else {
            print 'Count is zero or negative';
        }

        while (count > 0) {
            print count;
            count = count - 1;
        }
    end
    ```

- **String Manipulation**: Concatenate and manipulate strings easily.
    ```gear
    start
        let string greeting = 'Hello';
        let string name = 'World';
        let string message = greeting + ', ' + name + '!';
        print message;
    end
    ```

## Roadmap

### Completed

- Lexer implementation
- Tokenizer implementation
- Part of the AST implementation

### In Progress

- AST building phase

### Upcoming

- Implement access to struct properties and related features
- Complete the standard library implementation
- Improve error handling and reporting
- Enhance the language documentation
- Add more examples and tutorials
- Implement package management system
- Optimize the compiler for better performance
- Develop an integrated development environment (IDE) or plugin support for popular IDEs
- Increase test coverage and add more unit tests
- Gather user feedback and iterate on language design
- Add support for additional platforms and operating systems




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
