#!/bin/bash

scrape /Path/to/data2.csv

sort /Path/to/data2.csv > /Path/to/testdata2.csv
sort /Path/to/data.csv > /Path/to/testdata.csv
comm -1 -3 /Path/to/testdata.csv /Path/to/testdata2.csv | linkedInbot

mv /Path/to/data2.csv /Path/to/data.csv
