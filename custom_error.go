package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notfoundError struct {
	Message string
}

func (n *notfoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"ID cannot be empty"}
	}

	if id != "123" {
		return &notfoundError{"Data not found for ID: " + id}
	}

	return nil
}

func main() {
	err := SaveData("123", nil)

	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation Error:", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notfoundError); ok {
		// 	fmt.Println("Not Found Error:", notFoundErr.Message)
		// } else {
		// 	fmt.Println("Unknown Error:", err)
		// }

		switch finalErr := err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", finalErr.Error())
		case *notfoundError:
			fmt.Println("Not Found Error:", finalErr.Message)
		default:
			fmt.Println("Unknown Error:", finalErr)
		}
	} else {
		fmt.Println("Data saved successfully")
	}

}
