package api

type Credential interface {
	GetKey() string
	GetCertificate() []byte
	GetEncryptedPrivateKey() []byte
	GetAttributes() []CredentialAttribute
}

type CredentialAttribute interface {
	GetKey() string
	GetValue() interface{}
}

type CredentialStore interface {
	Save(credential Credential) (bool, error)
}
