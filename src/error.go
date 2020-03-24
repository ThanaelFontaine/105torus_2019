//
// EPITECH PROJECT, 2020
// 105torus_2019
// File description:
// error
//

package main

import (
    "strconv"
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func is_letter(c byte) bool {
    return (c >= 48 && c <= 57) || c == 45;
}

func chk_args(args []string) int {
    if (len(args) != 8) {
        os.Exit(84);
    }

    for i := 1; i < len(args); i++ {
        if (!is_letter(args[i][0])) {
            fmt.Print("Arguments should be numbers");
            os.Exit(84);
        }
    }

    opt, err := strconv.Atoi(args[1]);
    check(err);
    if (opt < 1 || opt > 3) {
        fmt.Print("WARNING: please use only, 1, 2 or 3 for selcting the method\n");
        os.Exit(84);
    }

    return opt;
}