# CLAUDE.md

Python data science project. Analysis, modeling, and reporting on tabular data.

## Environment

- Use pipenv: `pipenv install` to set up, `pipenv shell` to activate.
- Add packages with `pipenv install <pkg>` (dev tools: `pipenv install --dev <pkg>`).
- Default Python version; do not pin unless asked.

## Layout

- `data/raw/` — original data, never modified
- `data/processed/` — cleaned outputs from scripts
- `notebooks/` — exploration only; move reusable code into `src/`
- `src/` — importable modules (loading, cleaning, features, models)
- `tests/` — pytest tests for `src/`
- `reports/` — figures and summaries

## Commands

- Run tests: `pipenv run pytest`
- Lint/format: `pipenv run ruff check . && pipenv run ruff format .`
- Launch notebooks: `pipenv run jupyter lab`

## Conventions

- pandas for data frames, scikit-learn for models, matplotlib/seaborn for plots.
- Functions over classes unless state is needed. Type hints on public functions.
- Set random seeds explicitly for reproducibility.
- Paths via `pathlib`, relative to the project root. No hardcoded absolute paths.
- Keep notebooks clean: restart and run all before committing.

## Rules

- Never commit data files, credentials, or `.env`.
- Do not modify anything in `data/raw/`.
- Ask before adding heavy dependencies (e.g. torch, spark).
