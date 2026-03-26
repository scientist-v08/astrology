package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/scientist-v08/bphs/models"
	"github.com/scientist-v08/bphs/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// =============================================
// Request DTOs
// =============================================

type RegisterRequest struct {
	Username   string `json:"username" binding:"required,alphanum,min=4,max=20"`
	Password   string `json:"password" binding:"required,min=8"`
	FullName   string `json:"full_name" binding:"required,min=2,max=100"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required"`
}

// =============================================
// Response DTOs
// =============================================

type AuthResponse struct {
	Username              string    `json:"username"`
	FullName              string    `json:"full_name"`
}

type LoginResponse struct {
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	Username              string    `json:"username"`
	FullName              string    `json:"full_name"`
}

// ────────────────────────────────────────────────
// Request / Response DTOs for refresh
// ────────────────────────────────────────────────

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	Username              string    `json:"username"`
	FullName              string    `json:"full_name"`
}

// =============================================
// Controller / Handler
// =============================================

type AuthController struct {
	db           *gorm.DB
	tokenMaker   token.Maker
	accessDur    time.Duration
	refreshDur   time.Duration
}

func NewAuthController(db *gorm.DB, maker token.Maker, accessDuration, refreshDuration time.Duration) *AuthController {
	return &AuthController{
		db:         db,
		tokenMaker: maker,
		accessDur:  accessDuration,
		refreshDur: refreshDuration,
	}
}

// Register godoc
// @Summary      Create a new user
// @Description  Registers a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  RegisterRequest  true  "User registration data"
// @Success      201   {object}  AuthResponse
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /register [post]
func (ctrl *AuthController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if username already exists
	var existingUser models.User
	if err := ctrl.db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "username already taken"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := models.User{
		Username:       req.Username,
		HashedPassword: string(hashedPassword),
		FullName:       req.FullName,
		CreatedAt:      time.Now(),
		PasswordChangedAt: time.Time{}, // zero time = never changed
	}

	if err := ctrl.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	resp := AuthResponse{
		Username:              user.Username,
		FullName:              user.FullName,
	}

	c.JSON(http.StatusCreated, resp)
}

// Login godoc
// @Summary      Authenticate user
// @Description  Logs in a user and returns access + refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  LoginRequest  true  "Login credentials"
// @Success      200   {object}  AuthResponse
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := ctrl.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	// Create tokens (same as register)
	accessToken, accessPayload, err := ctrl.tokenMaker.CreateToken(
		user.Username,
		ctrl.accessDur,
		token.TokenTypeAccessToken,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create access token"})
		return
	}

	refreshToken, refreshPayload, err := ctrl.tokenMaker.CreateToken(
		user.Username,
		ctrl.refreshDur,
		token.TokenTypeRefreshToken,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create refresh token"})
		return
	}

	session := models.Session{
		ID:           refreshPayload.ID,          // the UUID from NewPayload
		Username:     user.Username,
		RefreshToken: refreshToken,               // the actual paseto string
		UserAgent:    c.Request.UserAgent(),
		ClientIP:     c.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
		CreatedAt:    time.Now(),
	}

	if err := ctrl.db.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store session"})
		return
	}

	resp := LoginResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		Username:              user.Username,
		FullName:              user.FullName,
	}

	c.JSON(http.StatusOK, resp)
}

// ────────────────────────────────────────────────
// Refresh godoc
// @Summary      Refresh access token using refresh token (with rotation)
// @Description  Validates refresh token, rotates it, issues new access + refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  RefreshTokenRequest  true  "Refresh token"
// @Success      200   {object}  RefreshTokenResponse
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      403   {object}  map[string]string  "blocked/used/mismatched"
// @Failure      500   {object}  map[string]string
// @Router       /refresh [post]
func (ctrl *AuthController) Refresh(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Verify the presented refresh token (PASETO signature + expiry + type)
	payload, err := ctrl.tokenMaker.VerifyToken(req.RefreshToken, token.TokenTypeRefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	// 2. Fetch session from DB using the token ID (jti equivalent)
	sessionID := payload.ID
	session, err := ctrl.getSession(sessionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "session not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// 3. Security checks
	if session.IsBlocked {
		c.JSON(http.StatusForbidden, gin.H{"error": "session is blocked"})
		return
	}

	if time.Now().After(session.ExpiresAt) {
		ctrl.db.Model(&session).Update("is_blocked", true)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token expired"})
		return
	}

	if session.RefreshToken != req.RefreshToken {
		// Mismatch → possible reuse attempt (stolen token used elsewhere)
		ctrl.db.Model(&session).Update("is_blocked", true)
		c.JSON(http.StatusForbidden, gin.H{"error": "refresh token mismatch - possible reuse detected"})
		return
	}

	// Device/session binding check
	currentUA := c.Request.UserAgent()
	currentIP := c.ClientIP()
	if session.UserAgent != currentUA || session.ClientIP != currentIP {
		c.JSON(http.StatusForbidden, gin.H{"error": "IP mismatch - possible reuse detected"})
		return
	}

	// 4. Get user info (for response)
	var user models.User
	if err := ctrl.db.Where("username = ?", payload.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	// 5. ROTATION: create NEW access + refresh tokens
	newAccessToken, newAccessPayload, err := ctrl.tokenMaker.CreateToken(
		payload.Username,
		ctrl.accessDur,
		token.TokenTypeAccessToken,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create access token"})
		return
	}

	newRefreshToken, newRefreshPayload, err := ctrl.tokenMaker.CreateToken(
		payload.Username,
		ctrl.refreshDur,
		token.TokenTypeRefreshToken,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create refresh token"})
		return
	}

	// 6. Update / replace session (rotation)
	newSession := models.Session{
		ID:           newRefreshPayload.ID, // new UUID from payload
		Username:     payload.Username,
		RefreshToken: newRefreshToken,
		UserAgent:    currentUA,
		ClientIP:     currentIP,
		IsBlocked:    false,
		ExpiresAt:    newRefreshPayload.ExpiredAt,
		CreatedAt:    time.Now(),
	}

	// Transaction: delete old + create new (atomic)
	tx := ctrl.db.Begin()
	if err := tx.Delete(&models.Session{}, "id = ?", session.ID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to invalidate old session"})
		return
	}
	if err := tx.Create(&newSession).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store new session"})
		return
	}
	tx.Commit()

	// 7. Response
	resp := RefreshTokenResponse{
		AccessToken:           newAccessToken,
		RefreshToken:          newRefreshToken,
		AccessTokenExpiresAt:  newAccessPayload.ExpiredAt,
		RefreshTokenExpiresAt: newRefreshPayload.ExpiredAt,
		Username:              user.Username,
		FullName:              user.FullName,
	}

	c.JSON(http.StatusOK, resp)
}

// Helper - you can also put this in a repository layer later
func (ctrl *AuthController) getSession(id uuid.UUID) (models.Session, error) {
	var s models.Session
	err := ctrl.db.Where("id = ?", id).First(&s).Error
	return s, err
}