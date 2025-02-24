# NameSnipe - OSINT Reconnaissance Tool

[![Go](https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white)](https://golang.org/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue?style=flat-square)](https://www.apache.org/licenses/LICENSE-2.0)

**NameSnipe** is a lightweight, command-line OSINT (Open-Source Intelligence) tool written in Go, designed for the reconnaissance phase of ethical hacking and security research.

---
### MVP Key Features

- **Flexible Input:** Search using a first name and optional last name via CLI arguments.

- **API Integration:** Currently supports Google Custom Search JSON API with a daily query limit tracker (100 free queries/day).

- **Query Limit Management:** Tracks usage and enforces the free tier limit, resetting daily.

- **Extensible Design:** Built with modularity in mind for easy addition of new data sources (e.g., Twitter, GitHub).

- **Timeout Handling:** Configurable HTTP timeout to prevent hangs during API requests.
