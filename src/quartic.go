//
// EPITECH PROJECT, 2020
// 105torus_2019
// File description:
// help
//

package main

import (
    "math"
    "os"
    "strconv"
)

func get_nbr(n string) int {
    nbr, err := strconv.Atoi(n);
    check(err);

    return nbr;
}

func quartic(x float64) float64 {
    a4 := get_nbr(os.Args[6]);
    a3 := get_nbr(os.Args[5]);
    a2 := get_nbr(os.Args[4]);
    a1 := get_nbr(os.Args[3]);
    a0 := get_nbr(os.Args[2]);

    return (float64(a4) * math.Pow(float64(x), 4)) + (float64(a3) * math.Pow(float64(x), 3)) + (float64(a2) * math.Pow(float64(x), 2)) + (float64(a1) * float64(x)) + float64(a0);
}