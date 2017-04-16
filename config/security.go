package config

import (
	jwt "github.com/dgrijalva/jwt-go"
)

/*
 * configuration related to JWT security settings
 */
type SecurityConfigDefn struct {
	Enabled       bool
	SigningMethod jwt.SigningMethod
}
