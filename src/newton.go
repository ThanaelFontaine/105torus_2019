//
// EPITECH PROJECT, 2020
// 105torus_2019
// File description:
// newton
//

package main;

import (
    "fmt"
    "os"
)

func newton(n int) {
    a := float64(0);
    b := float64(1);
    //last_res := float64(0);
    //var arr []float64;

    if (quartic(a) * quartic(b) >= 0) {
        fmt.Print("Bisection failed");
        os.Exit(0);
    }


}