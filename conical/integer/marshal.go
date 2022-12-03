package integer

type conicalIntegerMarshal struct {
	Namespace *string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Name      *string `json:"name,omitempty" yaml:"name,omitempty"`
	Version   *string `json:"version,omitempty" yaml:"version,omitempty"`
	Min       *int    `json:"min,omitempty" yaml:"min,omitempty"`
	Max       *int    `json:"max,omitempty" yaml:"max,omitempty"`
}

// func (t conical.ConicalInteger) MarshalJSON() ([]byte, error) {
// 	m := conicalIntegerMarshal{
// 		Namespace: s2p(t.Name),
// 	}

// 	return json.Marshal(m)
// }

// func s2p[T int | string](v T) *T {
// 	var z T

// 	if v == z {
// 		return nil
// 	}

// 	return &v
// }

// func (t *ConicalInteger) UnmarshalJSON(data []byte) error {
// 	var v string

// 	if err := json.Unmarshal(data, &v); err != nil {
// 		return err
// 	}

// 	e, err := Parse(v)

// 	if err != nil {
// 		return err
// 	}

// 	*t = e

// 	return nil
// }
