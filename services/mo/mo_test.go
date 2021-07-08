package mo

import (
	"fmt"
	"mo-for-desktop/model/resp_info"
	"mo-for-desktop/services/space"
	"testing"
)

func TestMo_ListSpaces(t *testing.T) {
	success := resp_info.Success(space.ListAll())

	fmt.Println(success)
}
