# GitHub User Activity CLI

This Go-based command-line tool retrieves and displays recent activity for a specified GitHub user.  It fetches events from the GitHub API and prints the repository names associated with those events.

## Getting Started

### Prerequisites

*   Go installed ([https://golang.org/doc/install](https://golang.org/doc/install))
*   Git (optional, for cloning the repository)

### Installation

1.  **Clone the repository (optional):**

    ```bash
    git clone [https://github.com/luisferllub230/github_user_activity.git](https://www.google.com/search?q=https://github.com/luisferllub230/github_user_activity.git)  # Replace with your actual repo URL if different
    cd github_user_activity
    ```

2.  **Get the dependencies:**

    ```bash
    go mod tidy
    ```

### Usage

The CLI tool accepts the following flags:

*   `-gh` or `--github`:  Specifies the GitHub username for which to retrieve activity.  This is a *required* flag.
*   `-pq` or `--quantity`: Specifies the number of GitHub event pages to retrieve.  Each page contains a maximum of 30 events.  The default is 30 (one page).
*   `-h` or `--help`: Displays the help message with usage instructions.

**Examples:**

*   **Get activity for a user (default 30 events):**

    ```bash
    go run main.go -gh your_github_username
    ```

*   **Get activity for a user (100 events - approximately 3 pages):**

    ```bash
    go run main.go -gh your_github_username -pq 3
    ```

*   **Display help:**

    ```bash
    go run main.go -h
    ```

### Building the executable

You can build the executable using the following command:

```bash
go build
```

### Roadmap project
https://roadmap.sh/projects/github-user-activity
