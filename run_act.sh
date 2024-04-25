#!/bin/bash

# Set up a temporary working environment.
WORKDIR=$(mktemp -d)
SELF_DIR=$(basename `pwd`)
trap "rm -rf $WORKDIR" EXIT

mkdir -p "$WORKDIR/$SELF_DIR"
cp -a .github "$WORKDIR"
cp -a * "$WORKDIR/$SELF_DIR"

cd $WORKDIR && \
git init -b main . && \
git add . && \
git commit -m "Initial commit" && \
git remote add origin git@localhost:$SELF_DIR

act -C $WORKDIR -s secret1=foo -s secret2=bar
