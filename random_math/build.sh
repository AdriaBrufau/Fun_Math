#!/bin/bash

gcc algebra_utils.c random.c -o random `sdl2-config --cflags --libs`
./random

