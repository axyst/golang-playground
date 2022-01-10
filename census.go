// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	ret := Resident{Name: name, Age: age, Address: address}
	return &ret
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	for _, v := range r.Address {
		if v == "" {
			return false
		}
	}
	return r.Name != "" && len(r.Address) != 0
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	ret := 0
	for _, v := range residents {
		if v.HasRequiredInfo() {
			ret++
		}
	}
	return ret
}
