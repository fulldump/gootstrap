package gootstrap

import "fmt"

func ExampleRunAll() {
	worker := func(name string) Runner {
		return func() (func() error, func() error) {
			return func() error {
					fmt.Println("start", name)
					return nil
				}, func() error {
					fmt.Println("stop", name)
					return nil
				}
		}
	}

	runner := RunAll(worker("http"), worker("metrics"))
	start, stop := runner()
	_ = start()
	_ = stop()

}
