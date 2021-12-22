#!/bin/sh
# Exit if the variable tests fail
set -e
set +x

# Check the database variables are set
test "$VDB_COMMAND"
set +e

# docker must be run in privilaged mode for mounts to work
echo "Running the statediff service"
