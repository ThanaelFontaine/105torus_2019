##
## EPITECH PROJECT, 2020
## 105torus_2019
## File description:
## Makefile
##

OUT = 105torus
SRC = src/torus.go

$(OUT): re

re:
	cd src && go build -o $(OUT) && mv $(OUT) ../