package flow

import (
	"log"

	pbFlow "github.com/micro/go-micro/flow/service/proto"
)

func stepToProto(step *Step) *pbFlow.Step {
	operations := make([]*pbFlow.Operation, 0, len(step.Operations))
	for _, op := range step.Operations {
		operations = append(operations, op.Encode())
	}

	pb := &pbFlow.Step{
		Name:       step.ID,
		Requires:   step.Requires,
		Required:   step.Required,
		Operations: operations,
	}

	return pb
}

func stepEqual(ostep *pbFlow.Step, nstep *pbFlow.Step) bool {
	if ostep.Name == nstep.Name {
		return true
	}

	return false
}

func protoToStep(pb *pbFlow.Step) *Step {
	ops := make([]Operation, 0, len(pb.Operations))
	for _, pbop := range pb.Operations {
		op, ok := Operations[pbop.Type]
		if !ok {
			log.Printf("unknown op %v\n", pbop)
			continue
		}
		nop := op.New()
		nop.Decode(pbop)
		ops = append(ops, nop)
	}

	st := &Step{
		ID:         pb.Name,
		Requires:   pb.Requires,
		Required:   pb.Required,
		Operations: ops,
	}

	return st
}
