//
// EPITECH PROJECT, 2020
// 105torus_2019
// File description:
// help
//

package main

import (
    "io/ioutil"
    "fmt"
)

func help() {
    data, err := ioutil.ReadFile("./static/h");
    check(err);

    fmt.Print(string(data));
}