# Ranma Coding Interview CLI

## Table of Contents

- [Overview](#overview)
- [Commands](#commands)
  - [candidates](#candidates)
  - [pick](#pick)

## Overview

The code package consists of a Go command-line application that provides two main commands, `candidates` and `pick`. These commands are designed to help select and generate coding questions from a question bank.

## Commands

### candidates

The `candidates` command provides functionality to list available coding questions from the question bank. When executed, it lists the available question keys from the `randma.QuestionBank`. It does not take any additional arguments.

Usage:
```shell
$ ./randma candidates
reverse_linked_list
dfs_search

$ randma pick -q dfs_search
```

