package constraint

// PredeclaredSignedInteger is a constraint that matches the
// five predeclared signed integer types.
type PredeclaredSignedInteger interface {
	int | int8 | int16 | int32 | int64
}

// SignedInteger is a constraint that matches any signed integer type.
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
