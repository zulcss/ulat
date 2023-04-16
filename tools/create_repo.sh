#!/bin/bash

repo=$1

ostree --repo=$repo init --mode=archive-z2
