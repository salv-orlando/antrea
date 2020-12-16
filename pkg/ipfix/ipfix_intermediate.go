// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipfix

import (
	"fmt"

	ipfixentities "github.com/vmware/go-ipfix/pkg/entities"
	ipfixintermediate "github.com/vmware/go-ipfix/pkg/intermediate"
)

var _ IPFIXAggregationProcess = new(ipfixAggregationProcess)

// IPFIXAggregationProcess interface is added to facilitate unit testing without involving the code from go-ipfix library.
type IPFIXAggregationProcess interface {
	Start()
	Stop()
	ForAllRecordsDo(callback ipfixintermediate.FlowKeyRecordMapCallBack) error
	DeleteFlowKeyFromMapWithoutLock(flowKey ipfixintermediate.FlowKey)
}

type ipfixAggregationProcess struct {
	AggregationProcess *ipfixintermediate.AggregationProcess
}

func NewIPFIXAggregationProcess(messageChan chan *ipfixentities.Message, workerNum int, correlateFields []string) (*ipfixAggregationProcess, error) {
	ap, err := ipfixintermediate.InitAggregationProcess(messageChan, workerNum, correlateFields)
	if err != nil {
		return nil, fmt.Errorf("error while initializing IPFIX intermediate process: %v", err)
	}

	return &ipfixAggregationProcess{
		AggregationProcess: ap,
	}, nil
}

func (ap *ipfixAggregationProcess) Start() {
	ap.AggregationProcess.Start()
}

func (ap *ipfixAggregationProcess) Stop() {
	ap.AggregationProcess.Stop()
}

func (ap *ipfixAggregationProcess) ForAllRecordsDo(callback ipfixintermediate.FlowKeyRecordMapCallBack) error {
	return ap.AggregationProcess.ForAllRecordsDo(callback)
}

func (ap *ipfixAggregationProcess) DeleteFlowKeyFromMapWithoutLock(flowKey ipfixintermediate.FlowKey) {
	ap.AggregationProcess.DeleteFlowKeyFromMapWithoutLock(flowKey)
}
