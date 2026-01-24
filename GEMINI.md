# Gemini Project Overview

This is a monorepo for various learning projects, managed with Bazel. It contains several sub-projects in different languages, including Go and JavaScript.

## Project Breakdown

### Go Projects

*   **`hello`**: A collection of small, self-contained Go programs demonstrating various language features and concepts.
    *   **`hello/v1`**: A simple "hello world" HTTP server.
    *   **Building and Running**: Use `bazel run` to execute the individual programs. For example, to run the `v1` server:
        ```bash
        bazel run //hello/v1:v1_http_s
        ```

*   **`try-go-clean-template`**: A more complex Go application that serves as a template for building services using the "Clean Architecture" pattern.
    *   **Project Structure**: The project is organized into layers, with business logic kept separate from infrastructure and framework code.
    *   **Building and Running**: The project uses a `Makefile` for common tasks.
        ```bash
        # Start the application and its dependencies
        make compose-up
        make run
        ```
    *   **Testing**:
        ```bash
        # Run integration tests
        make compose-up-integration-test
        ```

### JavaScript Projects

*   **`ical4notion`**: A Node.js application that creates an iCal feed from a Notion database.
    *   **Running the project**:
        ```bash
        # Install dependencies
        npm install
        # Run the server
        PORT=[PORT] node index.js [NOTION_API_KEY] [CALENDAR_NAME]
        ```
    *   **Docker**: The project can also be run using Docker Compose:
        ```bash
        docker-compose up -d
        ```

## Development Conventions

*   **Bazel**: The primary tool for managing the monorepo.
*   **Git Submodules**: Some of the sub-projects are included as Git submodules.
*   **Clean Architecture**: The `try-go-clean-template` project follows the "Clean Architecture" principles, which emphasizes separation of concerns and dependency inversion. This is a good reference for how to structure complex applications.

## Submodule Management

This project uses git submodules to include other repositories. Here are some common commands for managing submodules:

*   **Update submodules to the latest commit on their respective branches:**
    ```bash
    git submodule update --remote --merge
    ```
*   **Pull changes in all submodules:**
    ```bash
    git submodule foreach git pull
    ```
*   **Add changes in all submodules to be committed:**
    ```bash
    git submodule foreach git add .
    ```
*   **Commit changes in all submodules:**
    ```bash
    git submodule foreach git commit -m "Your commit message"
    ```
*   **Push changes in all submodules:**
    ```bash
    git submodule foreach git push
    ```
*   **Push changes in the main project and its submodules:**
    ```bash
    git push --recurse-submodules=on-demand
    ```
