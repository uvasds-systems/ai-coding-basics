---
name: new-project
description: Provision a new Python project with a uv-managed virtualenv, base packages, and a .gitignore. Use when the user asks to start or scaffold a new Python project.
argument-hint: <project-name> [python-version]
---

# NEW-PROJECT

Create a fresh Python project named `$1` (Python `$2`, default 3.12) using uv.

## Steps

1. Confirm `uv` is installed (`uv --version`). If missing: `curl -LsSf https://astral.sh/uv/install.sh | sh`.
2. Stop if directory `$1` already exists. Do not overwrite.
3. Initialize the project and virtualenv:
   ```bash
   uv init $1 --python ${2:-3.12}
   cd $1
   uv venv
   ```
4. Add base packages:
   ```bash
   uv add pandas numpy matplotlib
   uv add --dev pytest ruff ipykernel
   ```
5. Create folders: `mkdir -p data/raw data/processed notebooks src tests`
6. Write `.gitignore`:
   ```
   .venv/
   __pycache__/
   *.py[cod]
   .ipynb_checkpoints/
   .pytest_cache/
   .ruff_cache/
   .env
   data/
   .DS_Store
   ```
7. Verify: `uv run python --version && uv run pytest --co -q`
8. `git init`, then `git add . && git commit -m "Initial project scaffold"`.

## Rules

- Commit `pyproject.toml` and `uv.lock`; never commit `.venv/` or data.
- Run everything through `uv run`; do not activate the venv manually.
- Report the project path, Python version, and installed packages when done.
