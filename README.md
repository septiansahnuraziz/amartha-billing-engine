# Amartha - Billing Engine

## Getting Started

To start using the amartha-billing-engine console, follow the steps below:

1. Clone the repository:

    ```bash
    git clone https://github.com/septiansahnuraziz/amartha-billing-engine.git
    ```

2. Jump into the project directory:

    ```bash
    cd amartha-billing-engine
    ```

3. Build the application:

    ```bash
    go build
    ```

4. Run the server from console:

    ```bash
    ./amartha-billing-engine server
    ```

## Available Commands

The following commands are available in the amartha-billing-engine console:

### server

- Description: Starts the server.
- Usage:
    ```bash
    go run . server
    ```

### migrate

- Description: Migrates the database.
- Usage:
    ```bash
    go run . migrate
    ```
- Optional flags:
    - --step: Sets the maximum migration steps.
    - --direction: Sets the migration direction. up is default value


### create-migration [filename]

- Description: Creates a new database migration file.
- Usage:
    ```bash
    go run . create-migration [migration name]
    ```
- Example:
    ```bash
    go run . create-migration create_customers_table 
    ```


## Requirement
- Go version 1.22.5 as minimum
- Postgres 12 as minimum

## Local Development
- **Please follow step by step:**
- `$ go mod tidy`
- `$ cp config-example.yml config.yml`
- Please change in file `.env`, configuration `database` and `redis`
- run hot reloading:
  #### Server
    ```
    PORT=4000 make run
    ```

- See http://localhost:4000/v1/ping

## Local Development using Docker
- ```$ docker build -t amartha-billing-engine .```
- ```$ docker run -p 4000:4000 amartha-billing-engine```

## Linter
- `$ make lint-prepare`
- `$ make lint`
- Please do resolve

## Docs
- [Documentation and Collection API](https://erafone.atlassian.net/wiki/spaces/EM/pages/1320452130/Design+System "Click")



