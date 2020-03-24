//
// EPITECH PROJECT, 2020
// 105torus_2019
// File description:
// error
//

package main;

import (
    //"fmt"
    //"io/ioutil"
)

type DbElem struct {
    key string
    data float64
}

type Db struct {
    data []*DbElem
};

/* func GetDb() *Db {
    data, err := ioutil.ReadFile("./static/db");
    check(err);
    return &Db {

    }
}

func (this *Db) add(k string, v float64) {
    elem := &DbElem{
        key: k,
        data: v}

    this.data = append(this.data, elem);

    fmt.Print(this.data);
}

func (this *Db) save() {
    fmt.Print(this.data);
    //ioutil.WriteFile("./static/db.json", bytes, 0644);
} */