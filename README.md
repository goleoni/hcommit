# Happy commit

Happy commit will turn your commit messages more beautyfull, with emojis and a contexts to than.

The context should be a poject key of your company.

This should get metrics like "how many commits from that project have on your repository".

Can be a good alternative o manage the timesheet of developers.

# Snap

This project should be downloaded on snap repositories manager:

## Development

### Clean the actual snapping test

`snapcraft clean hcommit -s pull`

### Creating a testing snap to tests locally

Generating the snap:

`snapcraft`

To test the development of project, should be used:

`sudo snap install hcommit_1.0_amd64.snap --dangerous`

## To send to the market

https://www.digitalocean.com/community/tutorials/how-to-package-and-publish-a-snap-application-on-ubuntu-18-04-pt

# Usage

## Setting the global commit context

### Why use global context

Global context should be used when you will work on a specific project for a long time. Doing this will avoid the need to set a specific context on commit command every time.

### Setting the context

To set a global commit context the ./commit_setctx.sh script can be used like the example:

`. commit_setctx.sh globalctx`

After the global context is set, all commits will use this context. The only exception to it is when this prefix `CTX=anothectx` is added to a commit call. For this use only, the execution context will not be the global one, but the one added to the `CTX` environment variable.

## Commiter script usage

```
    Usage

    CTX=test ruby commiter.rb INTENTION FULL MESSAGE

    Permitted intentions
    -b --Bug fix (bugfix)
    -s --Start project
    -f --Finish project
    -d --Documentation or anyone comment on code only
    -w --Work in progress
    -r --Code review suggestion changes
    -p --Performance related changes
    -m --Maintenance changes: linter, config updates, etc.
    -rem --Code removing only
```

#### Example 1

```
# Set the next PR contexts based on the current Jira task
commit-setctx PR-123

# Start the PR execution
commit -s create the skeleton of the project

# Wip
commit -c implement cron class
commit -c iojasoijdaios

# Final commit before code review
commit -f fix final specs

# Fix some code review adjustments
commit -r code review adjustments

# Fix a bug or a problem that ocurred in the project
commit -b hahahha
```

#### Example 2

```
CTX=randomspec commit -m fix commit specs related with controller X
```

# Easy setup method

To easly setup the commiter, the execution of the `setup.sh` file will automatically include the following alias into `~/.bashrc`:

```
alias commit-ctx=". ABSOLUTE_PATH_TO_REPOSITORY/commiter/commit_setctx.sh"
alias commit="ruby ABSOLUTE_PATH_TO_REPOSITORY/commiter/commiter.rb"
```

# Greetings!

Have fun!
