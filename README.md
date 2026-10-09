# AI Coding Basics

**LLMs + Code: Getting Started**

For purposes of this primer, we assume the user is working with a coding agent such as Claude
Code or Codex, etc. Claude Code (CC) will be the primary reference point.

1. Install the CC command-line tool and authenticate. Students should sign into their Anthropic accounts provided through UVA.

2. Authenticate CC using `claude /login` and using option #1 - subscription model. This will bounce your authentication off of your signed-in browser.

3. Invoke CC with `claude`. This can be from within a only a terminal, or within the terminal pane of your favorite IDE.

## Setup

The easiest way to initially set up any project using CC is to run `claude /setup`. This will interview you about the work you will be doing in the current directory.

To configure global settings, you have a couple of options:

    - `/settings` - shows you global settings including model, auto-complete, feedback, etc.
    - Global `CLAUDE.md` file - within your home directory you will find a `~/.claude/` diretory. Inside of that you can create a `CLAUDE.md` file in which you can declare, in plain language, how you want to work, your preferences, your habits, and even snippets of code.

> **Read a sample `CLAUDE.md` file** - from Andrej Karpathy [here](https://github.com/multica-ai/andrej-karpathy-skills?source=post_page-----5b515c9e4cca-----------------------------------------). Note that an individual project can have its own `CLAUDE.md` file that will supplement any global `CLAUDE.md` you might have installed.

## Usage / Stats

At any time you can check your token usage with the `/usage` command. **Be aware that different models consume tokens at different rates**.

```
  Session
  
  Total cost:            $0.0000
  Total duration (API):  0s
  Total duration (wall): 15s
  Total code changes:    0 lines added, 0 lines removed
  Usage:                 0 input, 0 output, 0 cache read, 0 cache write
  
  Current session
  ████████                                          18% used
  Resets 1:39pm (America/New_York)

  What's contributing to your limits usage?
  Approximate, based on local sessions on this machine — does not include other devices or 
  claude.ai

  Last 24h · these are independent characteristics of your usage, not a breakdown

  100% of your usage came from /setup
   Heavy skills can be scoped down or run with a cheaper model via skill 
   frontmatter.
  
  Skills                  % of usage
  /setup                        100%
```

> **What's a session?** "Session" means two different things here. A CC *session* is one conversation, from launching `claude` until you `/exit` or `/clear`. It has its own context history, and the top "Session" block reports the cost, duration, and tokens for that conversation only. The "Current session" bar means something else: it measures your subscription's usage limit over a rolling window of about 5 hours. That window starts with your first message and resets at the time shown, no matter how many CC conversations you open or close in between.

To get more precise stats on your recent usage, issue the `/stats` command:

```
      Oct Nov Dec Jan Feb Mar Apr May Jun Jul Aug Sep Oct
      ················································▓·░·
  Mon ···············································▒█▒▒·
      ·············▒·································▒▓▒▓░
  Wed ···············································▓██░█
      ···············································▓·▓··
  Fri ···············································███░░
      ···················································

      Less ░ ▒ ▓ █ More

  All time · Last 7 days · Last 30 days

  Favorite model: Sonnet 5        Total tokens: 276.4m

  Sessions: 85                    Longest session: 26d 2h 16m
  Active days: 25/387             Longest streak: 5 days
  Most active day: Sep 11         Current streak: 1 day
  Input 7.9k · Output 1.6m · Cache read 268.5m · Cache write 6.3m

  Your input and output are ~19x the tokens in Brave New World
```

## Exit a Session

To get out of CC and reclaim a normal terminal command, simply enter `/exit`.

## Resume a Session

CC saves each session's transcript locally (under `~/.claude/projects/`), organized by the directory
you launched `claude` from. Any session can be resumed later, restoring its full conversation history.
This is separate from CC's *memory* features (such as `CLAUDE.md`), which carry across all sessions.

- `claude --resume` (or `-r`) - presents a list of previous sessions for the current directory. Use the up/down arrow keys to select, then press Enter.
- `claude --continue` (or `-c`) - skips the list and resumes the most recent session.
- `/resume` - opens the same list from within a running session.

> NOTE: It's tempting to always resume the previous session as you work on a project. While this works, the entire conversation history is sent to the model again with every message, so a long session costs more tokens on each turn. Prompt caching reduces this, but the cache expires after a few minutes, so the first message after resuming pays close to full price for the whole history. Very long contexts can also lower the quality of responses. Therefore, only resume a session when you are continuing the same line of thought or specific process. A fresh session in the same project is better for new work. Within a session, `/clear` starts over without exiting, and `/compact` replaces the history with a summary.

## Try It Out

A simple way to try out CC is to write your own application schema (definition of what it should do)
and then ask CC to build the code for it. While this is NOT recommended for building realworld
data science applications, it helps demonstrate the power of descriptive language to drive code
generation.

1. Create a new file named `SCHEMA.md` and describe an application.
2. As an example, let's write some code that will generate Pi to 20 decimals, as accurately as possible.
3. Do not use language-specific mentions of functions, classes, or packages (unless you need to), leaving your SCHEMA as code-agnostic as possible.
4. A sample prompt:

    Generate the code to perform the mathematical calculations necessary to print
    out the value of Pi to twenty decimal places. This calculation should be AS 
    ACCURATE AS POSSIBLE, using the relevant method, package, library, or approach
    as you determine.
    
    As you generate the code, also generate a markdown file that (1) explains how to
    use the code (i.e. how to run, compile, etc.); and (2) explains the calculation
    method and how this is verifiable using standard tests.
    
    If you need to install software, packages, create a virtual environment, etc. you
    will ask before doing anything, and get my consent.

5. Generate the code and specify a language (Python, C++, Go, etc.)

    ```
    ▗ ▗   ▖ ▖  Claude Code v2.1.295
               Opus 5.5 · Claude Enterprise
     ▘▘   ▝▝   ~/Development/ai-coding-basics
    
    
    ────────────────────────────────────────────────────────────────────────────────────────────────────
    ❯ generate the application in @SCHEMA.md using the Go language
    ────────────────────────────────────────────────────────────────────────────────────────────────────
      ⏵⏵ auto mode on (shift+tab to cycle)   
    ```

Notice that within a prompt you can use the `@` character to scope your prompt to a specific
file or folder.

## Skills

## Agents

## MCP
