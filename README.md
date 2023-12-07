# LDAP Authentication Module

This is an LDAP authentication module designed to authenticate users using LDAP. It returns an API key for successful logins and LDAP errors for unsuccessful logins.

## Pre-required Services

- MySQL Server
- goose: [Installation Guide](https://github.com/pressly/goose#install)

## Getting Started

Before you start using the LDAP authentication module, follow these steps:

### 1. Configuration File

Create a `config.yml` file in the `config` folder based on `config.example.yml` to set up the necessary configurations.

### 2. Environment Variables

Create a `.env` file and include the `CONFIG_PATH` variable, which should point to the static path for the `config` folder. Alternatively, you can define `CONFIG_PATH` as an environment variable.

Example `.env` file:

```dotenv
CONFIG_PATH=/path/to/config/folder
```

### 3. Update Script
Create an `update.sh` file in the shell folder based on `update.example.sh` to handle database updates.

## Usage

To use the LDAP authentication module, follow these steps:

1. Ensure the required services (MySQL Server) are running.
2. Set up the configuration files as described in the "Getting Started" section.
3. Execute the update script (`update.sh`) to update the database. The updater is now being called every time you start the app; feel free to comment out this part if desired.

## Contributing

We welcome contributions to enhance the LDAP authentication module.

### How to Contribute

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and ensure the tests pass.
4. Submit a pull request with a clear description of your changes.

Thank you for contributing!

## License

This project is licensed under the [MIT License](LICENSE). See the [LICENSE](LICENSE) file for details.
