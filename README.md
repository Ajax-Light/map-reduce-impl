# Distributed Systems Assignment

Simple Map-Reduce implementation in Go

## Dependencies

* Docker

## Usage

Bring up docker containers and network:

``` docker compose up -d ```

Teardown:

``` docker compose down ```

Rebuild images after code changes:

``` docker compose up --build ```

## Working

* Go routines for Mapper and Reducer Tasks
* RPC for Job Submissions to master node
