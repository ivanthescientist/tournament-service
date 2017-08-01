package validators

type ValidatedDTO interface {
	IsValid() bool
}
