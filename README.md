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
  ████████                              18% used
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

> **What's a session?** 

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

## Skills

## Agents

## MCP
