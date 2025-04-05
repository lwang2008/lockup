#!/bin/bash

# Get the commit hash
COMMIT_HASH=$(git rev-parse HEAD)

# Set a realistic date (e.g., 2 months ago)
DATE="2024-01-15T10:00:00"

# Use git filter-repo to update the commit date
git filter-repo --commit-callback "
if commit.original_id.decode('utf-8') == '$COMMIT_HASH':
    commit.author_date = b'$DATE'
    commit.committer_date = b'$DATE'
" 