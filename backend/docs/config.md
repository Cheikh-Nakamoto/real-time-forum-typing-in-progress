
# Config Package

The `config` package is responsible for managing project configuration settings, including reading custom environment variables from a `.env` file.

## Overview

The `config` package handles configuration for the project, allowing you to specify various settings required for your application to function correctly. It primarily deals with loading configuration variables from environment files.

## Usage

To use the `config` package in your project, you can simply import it:

```go
import "forum/config"
```

### Types
   - `ConfigValue string`

### Functions

- #### `LoadEnv() error`

    This function loads environment variables from a specified `.env` file into your application's environment.

    ##### Example

    ```go
    err := config.LoadEnv()
    if err != nil {
        // Handle the error
        fmt.Println("Error loading .env file:", err)
    }
    ```

- #### `Get(key string) ConfigValue`

    This function retrieves the value of an environment variable by its key.

    ##### Parameters

    - `key` (string): The key of the environment variable to retrieve.

    ##### Returns

    - `ConfigValue`: The value of the environment variable.

    ##### Example

    ```go
    // Assuming you have a variable named DATABASE_URL in your .env file
    databaseURL := config.Get("DATABASE_URL")
    fmt.Println("Database URL:", databaseURL)
    ```
- #### `(ConfigValue) ToString() string`

    ##### Returns
    - `string`: the string representation of the ConfigValue
    
- #### `(ConfigValue) ToInt() int`

    ##### Returns
    - `int`: the integer representation of the ConfigValue

- #### `(ConfigValue) ToBool() bool`

    ##### Returns
    - `bool`: the boolean representation of the ConfigValue

## Configuration

Before using the `config` package, you need to configure it by specifying the path to your `.env` file. Create a `.env` file in your project's root directory and define your environment variables in the following format:

```dotenv
VAR_NAME="value" #this is a comment
ANOTHER_VAR=false
ANOTHER_VAR1=4

```
