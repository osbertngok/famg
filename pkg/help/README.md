# Fullstack Architecture Makefile Generator


## Background

Currently the need to scaffolding a full stack project using python-typescript is surging. This is a util project creating a scaffolding framework using golang. It has two modes - interactive cli or config based.

It needs to know:

* The path (absolute or relative) of the new git repo;
* name of the git repo;
* full name of the git repo;

It would check:
* If target path already exists;
* If python3 is present;
* If node is present; 

It would produce:

* .gitignore based on python
* .ve3 folder, with pyvenv.cfg
* Setup black, isort, mypy
* pyproject.toml
* README.md

## Usage

To use the tool, you can either provide a config file or use the command line arguments.

```bash
famg --config-file="config.yaml"
```

```bash
famg --parent-path="." --name="todo-app" --fullname="Todo App"
```

To display the help message, you can use the following command:

```bash
famg --help
```