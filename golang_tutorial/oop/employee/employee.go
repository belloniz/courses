package employee

import (
	"fmt"
)

type employee struct {
	FirstName string
	LastName string
	TotaLeaves int
	LeavesTaken int
}

func New(FirstName string, lastName string, totaLeave int, leavesTaken int) employee {
	e := employee {FirstName, lastName, totaLeave, leavesTaken}
	return e
}

func (e employee) LeavesRemaining()  {
	fmt.Printf("%s %s has %d leaves remaining.\n", e.FirstName, e.LastName, (e.TotaLeaves - e.LeavesTaken))
}