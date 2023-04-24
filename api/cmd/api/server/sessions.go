package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yigitsadic/birthday-app-api/internal/auth"
	"github.com/yigitsadic/birthday-app-api/internal/responses"
	"github.com/yigitsadic/birthday-app-api/internal/sessions"
)

const RefreshTokenCookieName = "birthday-refresh-token"

// HandleCreateSession creates tokens with given credentials.
func (s *Server) HandleCreateSession(w http.ResponseWriter, r *http.Request) {
	var dto sessions.SessionDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		responses.RenderUnprocessableEntity(w, "invalid form body")
		return
	}

	if err := dto.Validate(); err != nil {
		responses.RenderUnprocessableEntity(w, "Form attributes are not valid")
		return
	}

	result, err := s.SessionRepository.FindUser(r.Context(), dto.Email)
	if err != nil {
		responses.RenderUnprocessableEntity(
			w,
			"Invalid credentials. Please check email/password.",
		)
		return
	}

	compare_result := auth.ComparePasswordHash(dto.Password, result.PasswordHash)

	if !compare_result {
		responses.RenderUnprocessableEntity(
			w,
			"Invalid credentials. Please check email/password.",
		)
		return
	}

	jwtStore := s.JWTStore

	accessToken, err := jwtStore.GenerateAccessToken(
		strconv.Itoa(result.ID),
		strconv.Itoa(result.CompanyId),
	)
	if err != nil {
		responses.RenderInternalServerError(w)
		return
	}

	refreshToken, err := jwtStore.GenerateRefreshToken(strconv.Itoa(result.ID))
	if err != nil {
		responses.RenderInternalServerError(w)
		return
	}

	resp := sessions.AuthenticationModel{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       result.ID,
		CompanyId:    result.CompanyId,
	}

	http.SetCookie(w, &http.Cookie{
		Name:     RefreshTokenCookieName,
		Value:    refreshToken,
		HttpOnly: true,
		MaxAge:   60*60*24*7 + 1, // 7 days and 1 second.

	})
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&resp)
}

// HandleRefreshSession refreshes token with token in the cookies.
func (s *Server) HandleRefreshSession(w http.ResponseWriter, r *http.Request) {
	jwtStore := s.JWTStore

	c, err := r.Cookie(RefreshTokenCookieName)
	if err != nil {
		responses.RenderUnauthenticated(w, responses.UnauthenticatedMessage)
		return
	}

	tok, err := jwtStore.ParseToken(c.Value)
	if err != nil {
		responses.RenderUnauthenticated(w, responses.UnauthenticatedMessage)
		return
	}

	if tok.TokenType != auth.RefreshTokenKey {
		responses.RenderUnauthenticated(w, responses.UnauthenticatedMessage)
		return
	}

	convertedId, err := strconv.Atoi(tok.Subject)
	if err != nil {
		responses.RenderInternalServerError(w)
		return
	}

	result, err := s.SessionRepository.FindById(r.Context(), convertedId)
	if err != nil {
		responses.RenderUnauthenticated(w, responses.UnauthenticatedMessage)
		return
	}

	accessToken, err := jwtStore.GenerateAccessToken(
		strconv.Itoa(result.ID),
		strconv.Itoa(result.CompanyId),
	)
	if err != nil {
		responses.RenderInternalServerError(w)
		return
	}

	refreshToken, err := jwtStore.GenerateRefreshToken(strconv.Itoa(result.ID))
	if err != nil {
		responses.RenderInternalServerError(w)
		return
	}

	resp := sessions.AuthenticationModel{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       result.ID,
		CompanyId:    result.CompanyId,
	}

	http.SetCookie(w, &http.Cookie{
		Name:     RefreshTokenCookieName,
		Value:    refreshToken,
		HttpOnly: true,
		MaxAge:   60*60*24*7 + 1, // 7 days and 1 second.

	})
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&resp)
}
