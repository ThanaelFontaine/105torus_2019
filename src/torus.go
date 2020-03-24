//
// EPITECH PROJECT, 2020
// 105torus_2019
// File description:
// help
//

package main

import (
    "os"
    "strconv"
)

func main() {
    args := os.Args;
    for v := range args {
        if (args[v] == "-h") {
            help();
            os.Exit(0);
        }
    }

    operation := chk_args(args);
    precision, err := strconv.Atoi(args[7]);
    check(err);

    switch operation {
        case 1:
            bisection(precision);
            break;
        case 2:
            break;
        case 3:
            break;
    }

    /* db := GetDb();
    db.add("test", float64(1));
    db.save(); */
}