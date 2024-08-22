// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package stateless

import (
	"encoding/json"
	"errors"

	"github.com/Kevionte/Go-Sovereign/common/hexutil"
	"github.com/Kevionte/Go-Sovereign/core/types"
)

var _ = (*extWitnessMarshalling)(nil)

// MarshalJSON marshals as JSON.
func (e extWitness) MarshalJSON() ([]byte, error) {
	type extWitness struct {
		Block   *types.Block    `json:"block"       gencodec:"required"`
		Headers []*types.Header `json:"headers"       gencodec:"required"`
		Codes   []hexutil.Bytes `json:"codes"`
		State   []hexutil.Bytes `json:"state"`
	}
	var enc extWitness
	enc.Block = e.Block
	enc.Headers = e.Headers
	if e.Codes != nil {
		enc.Codes = make([]hexutil.Bytes, len(e.Codes))
		for k, v := range e.Codes {
			enc.Codes[k] = v
		}
	}
	if e.State != nil {
		enc.State = make([]hexutil.Bytes, len(e.State))
		for k, v := range e.State {
			enc.State[k] = v
		}
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (e *extWitness) UnmarshalJSON(input []byte) error {
	type extWitness struct {
		Block   *types.Block    `json:"block"       gencodec:"required"`
		Headers []*types.Header `json:"headers"       gencodec:"required"`
		Codes   []hexutil.Bytes `json:"codes"`
		State   []hexutil.Bytes `json:"state"`
	}
	var dec extWitness
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Block == nil {
		return errors.New("missing required field 'block' for extWitness")
	}
	e.Block = dec.Block
	if dec.Headers == nil {
		return errors.New("missing required field 'headers' for extWitness")
	}
	e.Headers = dec.Headers
	if dec.Codes != nil {
		e.Codes = make([][]byte, len(dec.Codes))
		for k, v := range dec.Codes {
			e.Codes[k] = v
		}
	}
	if dec.State != nil {
		e.State = make([][]byte, len(dec.State))
		for k, v := range dec.State {
			e.State[k] = v
		}
	}
	return nil
}
