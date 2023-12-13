package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/sergeyyarin/otus_go_basic/hw06_testing/jsonapp/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)

	return data, err
}
