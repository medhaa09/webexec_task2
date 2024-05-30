
# Star Wars API Integration with Go Backend

## Overview

This project demonstrates how to create a backend server in Go that interacts with the Star Wars API to fetch planet data. The fetched data is cached locally to improve performance and reduce API load. The server exposes an HTTP endpoint to retrieve the cached or freshly fetched data.

## Prerequisites

- Go 1.15+
- `github.com/patrickmn/go-cache` package for caching

## Installation

### Clone the repository

```bash
git clone https://github.com/medhaa09/webexec_task2.git
```

### Install dependencies
```bash
go mod tidy
```
### run
```bash
go run main.go
```
### Access the API endpoint
To retrieve planet data, use the following endpoint:

```bash
http://localhost:8080/planet
```
### Configuration
You can adjust caching parameters in the main() function of main.go:
```bash
c = cache.New(10*time.Minute, 1*time.Minute)
```
10*time.Minute: Default expiration time for cached items.
1*time.Minute: Interval to purge expired items from the cache.
