//
// EPITECH PROJECT, 2020
// 105torus_2019
// File description:
// error
//

package main;

import (
    "fmt"
    "os"
    "math"
)

func bisection(n int) {
    a := float64(0);
    b := float64(1);
    last_res := float64(0);
    var arr []float64;

    if (quartic(a) * quartic(b) >= 0) {
        fmt.Print("Bisection failed");
        os.Exit(0);
    }

    for i := 0; true; i++ {
        med := float64((float64(a) + float64(b)) / 2);
        f_med := quartic(med);
        arr = append(arr, med);

        if ((last_res - f_med) > 0) {
            if ((last_res - f_med) < math.Pow(10, float64(-n))) {
                break;
            }
        }

        if ((last_res - f_med) > 0) {
            last_res = f_med;
        } else {
            last_res = 0;
        }

        if (quartic(a) * f_med < 0) {
            b = med;
        } else if (quartic(b) * f_med < 0) {
            a = med;
        /* } else if (f_med == 0) {
            fmt.Print("Found exact: ");
            fmt.Print(med);
            os.Exit(0); */
        } else {
            fmt.Print("Bisection failed");
            os.Exit(0);
        }
    }

    for i := 0; i < len(arr) - 1; i++ {
        fmt.Printf("x = %.6f\n", arr[i]);
    }
}