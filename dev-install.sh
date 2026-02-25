#!/usr/bin/env bash
set -e

go build -o quill .
sudo mv quill /usr/local/bin/quill
echo "Done. $(quill --version)"
