package analyzer

import "github.com/gaocegege/ali-clusterdata-analyzer/data"
import "fmt"

type Analyzer struct {
	BatchInstances []*data.BatchInstance
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		BatchInstances: []*data.BatchInstance{},
	}
}

func (a *Analyzer) ImportBatchInstance(batchInstances []*data.BatchInstance) {
	a.BatchInstances = append(a.BatchInstances, batchInstances...)
}

func (a *Analyzer) String() string {
	return fmt.Sprintf(`
		BatchInstances [%v]`, a.BatchInstances)
}
