# Contributing to DevCortex.ai

First off, thank you for considering contributing to DevCortex.ai! It's people like you that make this project such a great tool.

## How Can I Contribute?

### Reporting Bugs

This section explains how to submit a bug report for DevCortex.ai.

- **Ensure the bug was not already reported** by searching on GitHub under [Issues](https://github.com/melihyilman/devcortex.ai/issues).
- If you're unable to find an open issue addressing the problem, [open a new one](https://github.com/melihyilman/devcortex.ai/issues/new). Be sure to include a **title and clear description**, as much relevant information as possible, and a **code sample** or an **executable test case** demonstrating the expected behavior that is not occurring.

### Suggesting Enhancements

This section explains how to submit an enhancement suggestion for DevCortex.ai.

- Open a new issue to start a discussion about your idea. This is the best way to get feedback before you start working on something.
- Provide a clear and detailed explanation of the feature you want to see, why it's important, and how it should work.

### Pull Requests

We love pull requests. Here's a quick guide:

1.  **Fork the repo** and create your branch from `main`.
2.  **Add tests** for any new or changed functionality.
3.  **Ensure the test suite passes** (`go test -v ./...`).
4.  **Make sure your code lints**. Use `go fmt` and `go vet`.
5.  **Issue that pull request!**

## Development Setup

1.  Fork the repository.
2.  Clone your fork: `git clone https://github.com/melihyilman/devcortex.ai.git`
3.  Create a new branch: `git checkout -b feature/your-feature-name`
4.  Run `go mod tidy` to install dependencies.
5.  Make your changes.
6.  Run the application to test your changes: `go run ./cmd/web`

## Pull Request Process

1.  Ensure any install or build dependencies are removed before the end of the layer when doing a build.
2.  Update the `README.md` with details of changes to the interface, this includes new environment variables, new tools, or changes to existing tools.
3.  You may merge the Pull Request in once you have the sign-off of at least one other developer, or if you do not have permission to do that, you may request the second reviewer to merge it for you.

Thank you for your contribution!