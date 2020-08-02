#!/bin/bash

# load and run scraping scripts and parse output to CSV
# TODO - implement Docker stuff
# Set up initial database - should only run on initialization

psql recipes < start.sql

./scrape.sh



