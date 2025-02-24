# NameSniper
A lightweight, command-line OSINT (Open-Source Intelligence) tool written in Go, designed for the reconnaissance phase of ethical hacking and security research.

---
### MVP Key Features

    **Flexible Input:** Search using a first name and optional last name via CLI arguments.
    **API Integration:** Currently supports Google Custom Search JSON API with a daily query limit tracker (100 free queries/day).
    **Query Limit Management:** Tracks usage and enforces the free tier limit, resetting daily.
    **Extensible Design:** Built with modularity in mind for easy addition of new data sources (e.g., Twitter, GitHub).
    **Timeout Handling:** Configurable HTTP timeout to prevent hangs during API requests.
