package reader

import	"encoding/json"
import	"fmt"
import	"io"
import	"os"

import	"github.com/sergeyyarin/otus_go_basic/hw02_fix_app/types"


func ReadJSON(filePath string, limit int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)

	res := data

	return res, nil
}
