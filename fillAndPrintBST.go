package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	numbers, err := getNumbersFromConsole()
	if err == nil {
		BuildAndPrintBST(os.Stdout, numbers)
	} else {
		fmt.Println(err.Error())
	}
}

func getNumbersFromConsole() ([]int, error) {
	numberOfArguments := len(os.Args)
	if numberOfArguments > 1 {
		numbers := make([]int, 0, numberOfArguments-1)
		for i := 1; i < numberOfArguments; i++ {
			var number, err = strconv.Atoi(os.Args[i])
			if err != nil {
				return nil, errors.New("Solo se admiten números")
			}
			numbers = append(numbers, number)
		}
		return numbers, nil
	} else {
		return nil, errors.New("Mínimo se necesita 1 número")
	}
}

func BuildAndPrintBST(writer io.Writer, numbers []int) {
	var treeRoot = buildBST(numbers)
	printInOrder(writer, treeRoot)
}

func printInOrder(writer io.Writer, t *tree) {
	if t == nil {
		return
	}

	printInOrder(writer, t.left)
	fmt.Fprintf(writer, "%v ", t.value)
	printInOrder(writer, t.right)
}

func buildBST(numbers []int) (treeRoot *tree) {
	treeRoot = node(numbers[0])

	for _, number := range numbers[1:] {
		add(treeRoot, number)
	}
	return
}

func node(value int) (t *tree) {
	t = new(tree)
	t.value = value
	return
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		return node(value)
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
