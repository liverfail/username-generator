
# username-generator

## Overview
username-generator is a tool for anonymous users to generate random usernames for various websites, with the purpose of operational security (OPSEC) by avoiding the reuse of usernames across different platforms. You can customize how the usernameis generated to your liking.

## Installation
Clone the repository and build the project using the following commands:

```
git clone https://github.com/liverfail/username-generator.git
cd username-generator
go build -o username-generator
```

## Usage
The following command-line flags are available to customize the username generation process:

- `-pattern`: Specifies the pattern for generating usernames. Default pattern is "noun|verb|number".
- `-numbermax`: Specifies the maximum random for numbers in the generated username. Default max is 100.
- `-cleansymbols`: Removes symbols from the generated username if set to true. Default is true.
- `-capitalize`: Capitalizes each word in the generated username if set to true. Default is true.

Example usage:
```
./username-generator -pattern="adjective|noun|number" -numbermax=4 -clean=false -capitalize=true
```

## Command Flags
- `pattern`: Specifies the pattern for generating username.
- `numbermax`: Specifies the maximum random for numbers.
- `cleansymbols`: Removes symbols from the generated username if set to true.
- `capitalize`: Capitalizes each word in the generated username if set to true.

## Contributing
If you have any suggestions, bug reports, or want to contribute to the project, feel free to open an issue or create a pull request.

## License
This project is licensed under the [GPLv3 License](LICENSE).
