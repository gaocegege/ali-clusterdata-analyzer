package data

import (
	"fmt"
)

type InstanceStatus string

const (
	WAITING        = "Waiting"
	RUNNING        = "Running"
	TERMINATED     = "Terminated"
	formatTemplate = `
	StartTimestamp      %d
	EndTimestamp        %d
	JobID               %d
	taskID              %d
	machineID           %d
	Status              %s
	SequenceNumber      %d
	TotalSequenceNumber %d
	MaxCPU              %f
	AverageCPU          %f
	MaxNormMemory       %f
	AverageNormMemory   %f
`
)

type BatchInstance struct {
	StartTimestamp      int
	EndTimestamp        int
	JobID               int
	taskID              int
	machineID           int
	Status              InstanceStatus
	SequenceNumber      int
	TotalSequenceNumber int
	MaxCPU              float64
	AverageCPU          float64
	MaxNormMemory       float64
	AverageNormMemory   float64
}

func NewBatchInstanceFrom(record []string) *BatchInstance {
	return &BatchInstance{}
}

func (bi BatchInstance) String() string {
	return fmt.Sprintf(formatTemplate, bi.StartTimestamp, bi.EndTimestamp, bi.JobID, bi.taskID, bi.machineID, bi.Status, bi.SequenceNumber, bi.TotalSequenceNumber, bi.MaxCPU, bi.AverageCPU, bi.MaxNormMemory, bi.AverageNormMemory)
}
