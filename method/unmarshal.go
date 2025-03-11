package method

import "fmt"

/// Need to define custom error message for unmarshallers to work
// It is defined as an Error() method against a struct
// Idea from Claude.ai
type mismatch struct{
    this string
}

func (m *mismatch) Error() string{
    return fmt.Sprintf("Can't work with '%s'", m.this)
}

// Custom unmarshal methods
// These are called when the corresponding field is found in the JSON byte stream
func (sex *Sex) UnmarshalText(text []byte) error{
    which, ok := lookupSex[string(text)]
    if ok{
        *sex = which
        return nil
    }
    return &mismatch{this: string(text)}
}

func (hand *Handoff) UnmarshalText(text []byte) error{
    which, ok := lookupHand[string(text)]
    if ok{
        *hand = which
        return nil
    }
    return &mismatch{this: string(text)}
}

