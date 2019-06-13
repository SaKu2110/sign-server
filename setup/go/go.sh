#!/bin/bash

###########
# Env Var #
###########

# Common
ARCH=amd64
WORK_DIR=$(mktemp -d /tmp/go)

# Go
GO_VERSION=1.12.5

#############
# Down load #
#############
# golang
curl -OL https://dl.google.com/go/go$GO_VERSION.linux-$ARCH.tar.gz
sudo tar -C /usr/local -xzf go$GO_VERSION.linux-$ARCH.tar.gz

