package command

import (
	"fmt"

	"github.com/schweller/rumor/service"
)

func Usage() {
	count, limit := service.GetUsage()
	fmt.Println(fmt.Sprintf("You already used %v characters out of your limit %v", count, limit))
}
