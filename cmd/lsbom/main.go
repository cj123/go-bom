/*
  list.go - Golang implementation of bomutils

  Copyright (C) 2013 Fabian Renn - fabian.renn (at) gmail.com

  This program is free software; you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation; either version 2, or (at your option)
  any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program; if not, write to the Free Software
  Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston MA  02110-1301 USA.

  Initial work done by Joseph Coffland. Further contributions by Julian Devlin.
  Numerous further improvements by Baron Roberts.

  Golang translation by Callum Jones cj (at) icj (dot) me
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cj123/go-bom"
)

var file string

func init() {
	flag.StringVar(&file, "file", "", "bom file to read")
	flag.Parse()
}

func main() {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	ls, err := bom.Read(f)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range ls {
		fmt.Printf("%s\t\t\t%6d\t%d\t%s\n", file.Name(), file.Mode(), file.Size(), file.ModTime())
	}

	fmt.Printf("%d files\n", len(ls))
}
