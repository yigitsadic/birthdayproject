package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	validAccessToken  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiaXNzIjoiYmlydGhkYXktYXBwIiwiY29tcGFueV9pZCI6IjIiLCJleHAiOjE5MDE2MDY3NjksImlhdCI6MTY4MDY4MTk2OSwidG9rZW5fdHlwZSI6IkFjY2Vzc1Rva2VuIn0.Exz73fFaHKb22i40MzBNS05ANY6TQWOL5zj6lCeAC1E"
	validRefreshToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiaXNzIjoiYmlydGhkYXktYXBwIiwiZXhwIjoxOTAxNTk1OTY5LCJpYXQiOjE2ODA2NzExNjksInRva2VuX3R5cGUiOiJSZWZyZXNoVG9rZW4ifQ.qzonS856Aemc4n7on_NFyxBeZiPxO5cBeBdWoZ1F8qc"
)

func TestGenerateAccessToken(t *testing.T) {
	t.Parallel()

	j := NewJWT("ABCD")

	got, err := j.GenerateAccessToken("342", "23")

	require.Nil(t, err)
	assert.Equal(t, got[0:2], "ey") // JWT tokens start with ey

	parsed, err := j.ParseToken(got)

	assert.Nil(t, err)
	assert.Equal(t, AccessTokenKey, parsed.TokenType)
}

func TestGenerateRefreshToken(t *testing.T) {
	t.Parallel()

	j := NewJWT("ABCD")

	got, err := j.GenerateRefreshToken("342")

	require.Nil(t, err)
	assert.Equal(t, got[0:2], "ey") // JWT tokens start with ey
	parsed, err := j.ParseToken(got)

	assert.Nil(t, err)
	assert.Equal(t, RefreshTokenKey, parsed.TokenType)
}

func TestParseToken(t *testing.T) {
	j := NewJWT("ABCD")

	t.Run("it should parse valid token without any problem", func(t *testing.T) {
		t.Parallel()

		got, err := j.ParseToken(validAccessToken)

		expected := &JWTClaims{
			CompanyId: "2",
			TokenType: AccessTokenKey,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "birthday-app",
				Subject:   "1",
				IssuedAt:  jwt.NewNumericDate(time.Date(2023, 4, 5, 8, 6, 9, 0, time.UTC)),
				ExpiresAt: jwt.NewNumericDate(time.Date(2030, 4, 5, 8, 6, 9, 0, time.UTC)),
			},
		}

		require.Nil(t, err)

		a := assert.New(t)

		a.Equal(expected.TokenType, got.TokenType)
		a.Equal(expected.RegisteredClaims.Issuer, got.RegisteredClaims.Issuer)
		a.Equal(expected.RegisteredClaims.Subject, got.RegisteredClaims.Subject)
		a.Equal(expected.RegisteredClaims.IssuedAt.String(), got.RegisteredClaims.IssuedAt.String())
		a.Equal(expected.RegisteredClaims.ExpiresAt.String(), got.RegisteredClaims.ExpiresAt.String())
	})

	t.Run("it should give an error a token signed with another secret", func(t *testing.T) {
		t.Parallel()

		tok := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiaXNzIjoiYmlydGhkYXktYXBwIiwiY29tcGFueV9pZCI6IjIiLCJmaXJzdF9uYW1lIjoieWlnaXQiLCJsYXN0X25hbWUiOiJzYWRpYyIsImVtYWlsIjoieWlnaXRAZ29vZ2xlLmNvbSIsImV4cCI6MTkwMTU5NTk2OSwiaWF0IjoxNjgwNjcxMTY5fQ.ISQS2hEn_ieRqKoubfRRTmY6RaGExDWrcgsXGWKz6LA"
		got, err := j.ParseToken(tok)

		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, jwt.ErrTokenSignatureInvalid.Error())
	})

	t.Run("it should give an error if its expired", func(t *testing.T) {
		t.Parallel()

		tok := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiaXNzIjoiYmlydGhkYXktYXBwIiwiY29tcGFueV9pZCI6IjIiLCJmaXJzdF9uYW1lIjoieWlnaXQiLCJsYXN0X25hbWUiOiJzYWRpYyIsImVtYWlsIjoieWlnaXRAZ29vZ2xlLmNvbSIsImV4cCI6MTY4MDY5MTE2OSwiaWF0IjoxNjgwNjcxMTY5fQ.McSWm8o_KRC18J5xDzvrI164wQ0qifA1fzzZ-uyhNuo"
		got, err := j.ParseToken(tok)

		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "token is expired")
	})
}

func TestRefreshToken(t *testing.T) {
	t.Run("refresh refresh token", func(t *testing.T) {
		t.Parallel()

		a := assert.New(t)
		j := NewJWT("ABCD")

		original, err := j.ParseToken(validRefreshToken)
		a.Nil(err)

		newTok, err := j.RefreshToken(validRefreshToken)
		a.Nil(err)

		parsed, err := j.ParseToken(newTok)
		a.Nil(err)

		a.Equal(original.Issuer, parsed.Issuer)
		a.Equal(original.Subject, parsed.Subject)
		a.Equal(original.TokenType, parsed.TokenType)
		a.NotEqual(original.IssuedAt.String(), parsed.IssuedAt.String())
		a.NotEqual(original.ExpiresAt.String(), parsed.ExpiresAt.String())
	})

	t.Run("unknown token type should not refreshed", func(t *testing.T) {
		t.Parallel()

		j := NewJWT("ABCD")
		tok := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiaXNzIjoiYmlydGhkYXktYXBwIiwiZXhwIjoxOTAxNTk1OTY5LCJpYXQiOjE2ODA2NzExNjksInRva2VuX3R5cGUiOiJXZWlyZFRva2VuIn0.T8HNCaAEorYpYK79Ysp6bf-AbYFlQBwJPbHUH9GMWk0"

		res, err := j.RefreshToken(tok)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "unknown token type")
		assert.Equal(t, "", res)

	})

	t.Run("refresh access token", func(t *testing.T) {
		t.Parallel()

		a := assert.New(t)

		j := NewJWT("ABCD")

		oldParsed, err := j.ParseToken(validAccessToken)
		a.Nil(err)
		a.NotNil(oldParsed)

		newTok, err := j.RefreshToken(validAccessToken)
		a.Nil(err)
		a.NotEmpty(newTok)

		parsed, err := j.ParseToken(newTok)
		a.Nil(err)
		a.NotNil(parsed)

		a.Equal(oldParsed.Issuer, parsed.Issuer)
		a.Equal(oldParsed.Subject, parsed.Subject)
		a.Equal(oldParsed.CompanyId, parsed.CompanyId)
		a.Equal(oldParsed.TokenType, parsed.TokenType)
		a.NotEqual(oldParsed.IssuedAt.String(), parsed.IssuedAt.String())
		a.NotEqual(oldParsed.ExpiresAt.String(), parsed.IssuedAt.String())
	})
}
