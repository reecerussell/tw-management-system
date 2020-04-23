package jwt

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Environment variables
const (
	configFileVar = "JWT_CONFIG"
)

// Standard (IANA registered) claim names.
const (
	issuer    = "iss"
	subject   = "sub"
	audience  = "aud"
	expires   = "exp"
	notBefore = "nbf"
	issued    = "iat"
	id        = "jti"
)

const (
	alg          = crypto.SHA256
	rsa256Header = "eyJhbGciOiJSUzI1NiJ9"
)

// Custom claims.
const (
	ClaimUserID = "user_id"
)

var encoding = base64.RawURLEncoding

//
// JWT
//

// JWT is a collection of data needed to create a JSON-Web-Token.
type JWT struct {
	Config *Config

	// Claims data which will be used in the signed token.
	Claims *Claims

	// Raw token data
	data []byte
}

// New creates a new instance of JWT with the given claims.
func New(c *Claims) *JWT {
	return &JWT{
		Claims: c,
	}
}

// FromToken creates a new JWT instance and simply
// sets the unexported data value. Check can be called after this.
func FromToken(token []byte) (*JWT, error) {
	t := &JWT{
		data: token,
	}

	fd, ld, sig, h, err := scan(token)
	if err != nil {
		return nil, err
	}

	err = t.parseClaims(token[fd+1:ld], sig, h)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// GetData returns the token data from the JWT.
func (jwt *JWT) GetData() []byte {
	return jwt.data
}

// String returns a sign JWT's data in string form.
func (jwt *JWT) String() string {
	return string(jwt.data)
}

// AccessToken returns an *AccessToken populated with data from jwt.
func (jwt *JWT) AccessToken() *AccessToken {
	ac := new(AccessToken)
	ac.AccessToken = jwt.String()
	ac.Expires = float64(*jwt.Claims.Expires)

	return ac
}

// LoadClaimsFromConfig loads a JSON config file, then
// sets the JWT's config and claims. This doesn't set
// the expiry, not before or issued claims.
func (jwt *JWT) LoadClaimsFromConfig() error {
	if jwt.Config == nil {
		var c Config
		err := c.LoadFromFile(os.Getenv(configFileVar))
		if err != nil {
			return err
		}

		jwt.Config = &c
	}

	jwt.Claims.Issuer = jwt.Config.Issuer
	jwt.Claims.Audiences = jwt.Config.Audiences

	return nil
}

// Sign uses a private rsa key to sign the claims of
// the JWT, creating a new token.
func (jwt *JWT) Sign(key *rsa.PrivateKey) error {
	if err := jwt.Claims.sync(); err != nil {
		return err
	}

	digest := alg.New()

	// create a new token
	l := len(rsa256Header) + 1 + encoding.EncodedLen(len(jwt.Claims.Raw))
	token := make([]byte, l, l+1+encoding.EncodedLen(key.Size()))

	i := copy(token, rsa256Header)
	token[i] = '.'
	i++
	encoding.Encode(token[i:], jwt.Claims.Raw)
	token = token[:l]

	digest.Write(token)

	// use signature space as a buffer while not set
	buf := token[len(token):]

	sig, err := rsa.SignPKCS1v15(rand.Reader, key, alg, digest.Sum(buf))
	if err != nil {
		return err
	}

	i = len(token)
	token = token[:cap(token)]
	token[i] = '.'
	encoding.Encode(token[i+1:], sig)

	jwt.data = token

	return nil
}

// Check returns whether the signature is valid and if
// the claims are "in time", i.e. have no expired.
func (jwt *JWT) Check(key *rsa.PublicKey) (bool, error) {
	fd, ld, sig, h, err := scan(jwt.data)
	if err != nil {
		return false, err
	}

	digest := alg.New()
	digest.Write(jwt.data[:ld])

	err = rsa.VerifyPKCS1v15(key, alg, digest.Sum(sig[len(sig):]), sig)
	if err != nil {
		return false, fmt.Errorf("signature mismatch")
	}

	err = jwt.parseClaims(jwt.data[fd+1:ld], sig, h)
	if err != nil {
		return false, err
	}

	exp, expOK := jwt.Claims.Number(expires)
	nbf, nfbOK := jwt.Claims.Number(notBefore)

	n := NewNumericTime(time.Now())
	// n will only be nil if time.Now() is zero.
	// Assume time.Now() won't return zero.
	// if n == nil {
	// 	return !expOK && !nfbOK, nil
	// }

	f := float64(*n)
	return (!expOK || exp > f) && (!nfbOK || nbf <= f), nil
}

// Header is a critical subset of the registered “JOSE Header Parameter Names”.
type header struct {
	Alg  string   // algorithm
	Kid  string   // key identifier
	Crit []string // extensions which must be supported and processed
}

// Buf remains in use as the Raw field.
func (jwt *JWT) parseClaims(enc, buf []byte, header *header) error {
	buf = buf[:cap(buf)]
	n, err := encoding.Decode(buf, enc)
	if err != nil {
		return fmt.Errorf("malformed payload: %v", err)
	}

	buf = buf[:n]

	c := &Claims{
		Raw:   json.RawMessage(buf),
		Set:   make(map[string]interface{}),
		KeyID: header.Kid,
	}

	if err = json.Unmarshal(buf, &c.Set); err != nil {
		return fmt.Errorf("malformed payload: %v", err)
	}

	m := c.Set
	if s, ok := m[issuer].(string); ok {
		delete(m, issuer)
		c.Issuer = s
	}

	if s, ok := m[subject].(string); ok {
		delete(m, subject)
		c.Subject = s
	}

	// “In the general case, the "aud" value is an array of case-sensitive
	// strings, each containing a StringOrURI value.  In the special case
	// when the JWT has one audience, the "aud" value MAY be a single
	// case-sensitive string containing a StringOrURI value.”
	switch a := m[audience].(type) {
	case []interface{}:
		allStrings := true
		for _, o := range a {
			if s, ok := o.(string); ok {
				c.Audiences = append(c.Audiences, s)
			} else {
				allStrings = false
			}
		}

		if allStrings {
			delete(m, audience)
		}

	case string:
		delete(m, audience)
		c.Audiences = []string{a}
	}

	if f, ok := m[expires].(float64); ok {
		delete(m, expires)
		c.Expires = (*NumericTime)(&f)
	}

	if f, ok := m[notBefore].(float64); ok {
		delete(m, notBefore)
		c.NotBefore = (*NumericTime)(&f)
	}

	if f, ok := m[issued].(float64); ok {
		delete(m, issued)
		c.Issued = (*NumericTime)(&f)
	}

	if s, ok := m[id].(string); ok {
		delete(m, id)
		c.ID = s
	}

	jwt.Claims = c

	return nil
}

func scan(token []byte) (int, int, []byte, *header, error) {
	fd := bytes.IndexByte(token, '.')
	ld := bytes.LastIndexByte(token, '.')
	if ld <= fd {
		return 0, 0, nil, nil, fmt.Errorf("missing base64 part")
	}

	buf := make([]byte, encoding.DecodedLen(len(token)))
	n, err := encoding.Decode(buf, token[:fd])
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("malformed header: %v", err)
	}
	h := new(header)
	if err := json.Unmarshal(buf[:n], h); err != nil {
		return 0, 0, nil, nil, fmt.Errorf("malformed header: %v", err)
	}

	// “If any of the listed extension Header Parameters are not understood
	// and supported by the recipient, then the JWS is invalid.”
	// — “JSON Web Signature (JWS)” RFC 7515, subsection 4.1.11
	if len(h.Crit) != 0 {
		return 0, 0, nil, nil, fmt.Errorf("unsupported critical extension in JOSE header: %q", h.Crit)
	}

	n, err = encoding.Decode(buf, token[ld+1:])
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("malformed signature: %v", err)
	}

	return fd, ld, buf[:n], h, nil
}

//
// CLAIMS
//

// RegisteredClaims “JSON Web Token Claims” is a subset of the IANA registration.
// See <https://www.iana.org/assignments/jwt/claims.csv> for the full listing.
//
// Each field is optional—there are no required claims. The string values are
// case sensitive.
type RegisteredClaims struct {
	// Issuer identifies the principal that issued the JWT.
	Issuer string `json:"iss,omitempty"`

	// Subject identifies the principal that is the subject of the JWT.
	Subject string `json:"sub,omitempty"`

	// Audiences identifies the recipients that the JWT is intended for.
	Audiences []string `json:"aud,omitempty"`

	// Expires identifies the expiration time on or after which the JWT
	// must not be accepted for processing.
	Expires *NumericTime `json:"exp,omitempty"`

	// NotBefore identifies the time before which the JWT must not be
	// accepted for processing.
	NotBefore *NumericTime `json:"nbf,omitempty"`

	// Issued identifies the time at which the JWT was issued.
	Issued *NumericTime `json:"iat,omitempty"`

	// ID provides a unique identifier for the JWT.
	ID string `json:"jti,omitempty"`
}

// Claims are the (signed) statements of a JWT.
type Claims struct {
	// RegisteredClaims field values take precedence.
	RegisteredClaims

	// Set maps claims by name, for usecases beyond the Registered fields.
	// The Sign methods copy each non-zero Registered value into Set when
	// the map is not nil. The Check methods map claims in Set if the name
	// doesn't match any of the Registered, or if the data type won't fit.
	// Entries are treated conform the encoding/json package.
	//
	//	bool, for JSON booleans
	//	float64, for JSON numbers
	//	string, for JSON strings
	//	[]interface{}, for JSON arrays
	//	map[string]interface{}, for JSON objects
	//	nil for JSON null
	//
	Set map[string]interface{}

	// Raw encoding as is within the token. This field is read-only.
	Raw json.RawMessage

	// “The "kid" (key ID) Header Parameter is a hint indicating which key
	// was used to secure the JWS.  This parameter allows originators to
	// explicitly signal a change of key to recipients.  The structure of
	// the "kid" value is unspecified.  Its value MUST be a case-sensitive
	// string.  Use of this Header Parameter is OPTIONAL.”
	// — “JSON Web Signature (JWS)” RFC 7515, subsection 4.1.4
	KeyID string
}

// sync updates the Raw field. When the Set field is not nil,
// then all non-zero Registered values are copied into the map.
func (c *Claims) sync() error {
	var payload interface{}

	if c.Set == nil {
		payload = &c.RegisteredClaims
	} else {
		payload = c.Set

		if c.Issuer != "" {
			c.Set[issuer] = c.Issuer
		}
		if c.Subject != "" {
			c.Set[subject] = c.Subject
		}
		switch len(c.Audiences) {
		case 0:
			break
		case 1: // single string
			c.Set[audience] = c.Audiences[0]
		default:
			array := make([]interface{}, len(c.Audiences))
			for i, s := range c.Audiences {
				array[i] = s
			}
			c.Set[audience] = array
		}
		if c.Expires != nil {
			c.Set[expires] = float64(*c.Expires)
		}
		if c.NotBefore != nil {
			c.Set[notBefore] = float64(*c.NotBefore)
		}
		if c.Issued != nil {
			c.Set[issued] = float64(*c.Issued)
		}
		if c.ID != "" {
			c.Set[id] = c.ID
		}
	}

	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	c.Raw = json.RawMessage(bytes)
	return nil
}

// Add inserts a definition to the claims set. If the claim
// is a registered claim, it will have to be the right data type.
// If a definition already exists, it will be overriden, otherwise
// will be inserted.
func (c *Claims) Add(name string, value interface{}) error {
	// c.Set could be nil, so create an empty map
	if c.Set == nil {
		c.Set = make(map[string]interface{})
	}

	if name == issuer {
		if s, ok := value.(string); ok {
			c.Issuer = s
			return nil
		}

		return fmt.Errorf("claim '%s' has to be a string", issuer)
	}

	if name == subject {
		if s, ok := value.(string); ok {
			c.Subject = s
			return nil
		}

		return fmt.Errorf("claim '%s' has to be a string", subject)
	}

	if name == audience {
		if a, ok := value.([]interface{}); ok {
			auds := make([]string, len(a))
			allStrings := true
			for i, o := range a {
				if s, is := o.(string); is {
					auds[i] = s
				} else {
					allStrings = false
				}
			}

			if allStrings {
				c.Audiences = auds
				return nil
			}
		}

		return fmt.Errorf("claim '%s' has to be a string array", audience)
	}

	if name == expires {
		if f, ok := value.(float64); ok {
			c.Issued = (*NumericTime)(&f)
			return nil
		}

		return fmt.Errorf("claim '%s' has to be a float64", issued)
	}

	if name == notBefore {
		if f, ok := value.(float64); ok {
			c.NotBefore = (*NumericTime)(&f)
			return nil
		}

		return fmt.Errorf("claim '%s' has to be a float64", notBefore)
	}

	if name == issued {
		if f, ok := value.(float64); ok {
			c.Issued = (*NumericTime)(&f)
			return nil
		}

		return fmt.Errorf("claim '%s' has to be a float64", issued)
	}

	if name == id {
		if s, ok := value.(string); ok {
			c.ID = s
			return nil
		}

		return fmt.Errorf("claim '%s' has to be a string", id)
	}

	v, ok := c.Set[name]
	if ok {
		if v == value {
			return nil
		}
	}

	c.Set[name] = value

	return nil
}

// String returns the claim when present and if the representation is a JSON string.
func (c *Claims) String(name string) (value string, ok bool) {
	// try Registered first
	switch name {
	case issuer:
		value = c.Issuer
	case subject:
		value = c.Subject
	case audience:
		if len(c.Audiences) == 1 {
			return c.Audiences[0], true
		}
		if len(c.Audiences) != 0 {
			return "", false
		}
	case id:
		value = c.ID
	}
	if value != "" {
		return value, true
	}

	// fallback
	value, ok = c.Set[name].(string)
	return
}

// Number returns the claim when present and if the representation is a JSON number.
func (c *Claims) Number(name string) (value float64, ok bool) {
	// try Registered first
	switch name {
	case expires:
		if c.Expires != nil {
			return float64(*c.Expires), true
		}
	case notBefore:
		if c.NotBefore != nil {
			return float64(*c.NotBefore), true
		}
	case issued:
		if c.Issued != nil {
			return float64(*c.Issued), true
		}
	}

	// fallback
	value, ok = c.Set[name].(float64)
	return
}

// Copy takes all non-empty/non-nil values from n and copies them to c.
func (c *Claims) Copy(n *Claims) {
	if n == nil {
		return
	}

	if n.Audiences != nil {
		c.Audiences = n.Audiences
	}

	if n.Expires != nil {
		c.Expires = n.Expires
	}

	if n.Issued != nil {
		c.Issued = n.Issued
	}

	if n.Issuer != "" {
		c.Issuer = n.Issuer
	}

	if n.KeyID != "" {
		c.KeyID = n.KeyID
	}

	if n.NotBefore != nil {
		c.NotBefore = n.NotBefore
	}

	if n.Subject != "" {
		c.Subject = n.Subject
	}

	if n.Set != nil {
		for k, v := range n.Set {
			// If this returns an error, ignore and move on.
			_ = c.Add(k, v)
		}
	}
}

//
// NUMERIC TIME
//

// NumericTime implements NumericDate: “A JSON numeric value representing
// the number of seconds from 1970-01-01T00:00:00Z UTC until the specified
// UTC date/time, ignoring leap seconds.”
type NumericTime float64

// NewNumericTime returns the the corresponding representation with nil for the
// zero value. Do t.Round(time.Second) for slighly smaller token production and
// compatibility. See the bugs section for details.
func NewNumericTime(t time.Time) *NumericTime {
	if t.IsZero() {
		return nil
	}

	n := NumericTime(float64(t.UnixNano()) / 1e9)

	return &n
}

// Time returns the Go mapping with the zero value for nil.
func (n *NumericTime) Time() time.Time {
	if n == nil {
		return time.Time{}
	}
	return time.Unix(0, int64(float64(*n)*float64(time.Second))).UTC()
}

// String returns the ISO representation or the empty string for nil.
func (n *NumericTime) String() string {
	if n == nil {
		return ""
	}
	return n.Time().Format(time.RFC3339Nano)
}

//
// KEYS
//

// KeyRegister holds a public and private RSA key, used
// to sign and check JWTs.
type KeyRegister struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// NewKeyRegisterFromFile reads data from a given file
// then builds a new KeyRegister,
func NewKeyRegisterFromFile(privateKeyFile string, keyPassword []byte) (*KeyRegister, error) {
	bytes, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}

	return NewKeyRegister(bytes, keyPassword)
}

// NewKeyRegister returns a KeyRegister, built by reading
// a private RSA key data, in a PEM format.
func NewKeyRegister(privateKeyData, keyPassword []byte) (*KeyRegister, error) {
	var keys KeyRegister

	err := keys.loadPEM(privateKeyData, keyPassword)
	if err != nil {
		return nil, err
	}

	return &keys, err
}

// ReadPEM returns a key from PEM-encoded data. PEM encryption
// is enforced for non-empty password values.
func (k *KeyRegister) loadPEM(data, password []byte) error {
	block, _ := pem.Decode(data)
	if block == nil {
		return fmt.Errorf("invalid key format")
	}

	if x509.IsEncryptedPEMBlock(block) {
		bytes, err := x509.DecryptPEMBlock(block, password)
		if err != nil {
			return err
		}
		block.Bytes = bytes
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	k.PrivateKey = key
	k.PublicKey = &key.PublicKey

	return nil
}

//
// CONFIG
//

// Config is a type which maps to a JSON config file.
type Config struct {
	Audiences   []string `json:"audiences"`
	Issuer      string   `json:"issuer"`
	ExpiryHours float64  `json:"expiryHours"`
}

// LoadFromFile reads JSON data from file at the give path,
// then sets c's data.
func (c *Config) LoadFromFile(filePath string) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("couldn't read config file: %v", err)
	}

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return fmt.Errorf("config file is an invalid format: %v", err)
	}

	return nil
}
