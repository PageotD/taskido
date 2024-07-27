# Taskido
Taskido is a lightweight, command-line task management tool written in Go. Designed for simplicity and efficiency, Taskido helps you keep track of your tasks and projects with ease.

## Commit Message Guidelines

To maintain a clear and structured history in the project, I use a system based on the Angular commit message conventions. This ensures consistency and helps automate versioning and changelog generation.

### Commit Message Format

Each commit message should be structured as follows:

```bash
<type>(<scope>): <subject>
```
- `<type>`: The type of change being made (e.g., feat, fix, docs, etc.)
- `<scope>`: (Optional) The scope of the change (e.g., ui, api, build). This can be omitted if not relevant.
- `<subject>`: A brief description of the change.

### Commit Types

Must be one of the following:

* **build**: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
* **ci**: Changes to our CI configuration files and scripts (examples: CircleCi, SauceLabs)
* **docs**: Documentation only changes
* **feat**: A new feature
* **fix**: A bug fix
* **perf**: A code change that improves performance
* **refactor**: A code change that neither fixes a bug nor adds a feature
* **test**: Adding missing tests or correcting existing tests

### Summary

Use the summary field to provide a succinct description of the change:

* use the imperative, present tense: "change" not "changed" nor "changes"
* don't capitalize the first letter
* no dot (.) at the end

By adhering to these conventions, I ensure my commit history is clear, consistent, and useful for tracking changes and generating automated releases.

## Automatic Versioning with Semantic Versioning (SemVer)

Automatic versioning in the project follows the Semantic Versioning (SemVer) guidelines. Version numbers are updated based on the type and content of commit messages. Here’s how versioning is handled:
Version Number Format

### SEMVER based versioning
Version numbers are in the format MAJOR.MINOR.PATCH:

* **MAJOR**: Incremented for incompatible changes, significant internal changes (major refactoring, architectural improvements, substantial updates to dependencies, a significant evolution of the project), or strategic decisions (decisions driven by project goals, compatibility considerations, or policy changes that require a new major version to reflect a substantial shift in the project).
* **MINOR**: Incremented for adding functionality in a backward-compatible manner.
* **PATCH**: Incremented for backward-compatible bug fixes.

### Adapted automatic versioning Rules

* **MAJOR Version**: Incremented for breaking changes with "VERSION CHANGE" in the commit message.
* **MINOR Version**: Incremented for new features (feat or feature).
* **PATCH Version**: Incremented for bug fixes (fix).
* **No Version Increment**: For other commit types such as docs, style, chore, refactor, and perf.

This approach ensures that version numbers accurately reflect the impact of changes made to the codebase, following the Semantic Versioning principles.