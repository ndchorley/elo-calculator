#!/bin/bash

go clean -cache && \
    go test ./calculator && \
    go build
