package handler

import (
	"errors"
	"fmt"
	"math/rand"
	"mitra/domain"
	"mitra/store"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/render"
)

func generateSecret(length, lower, upper, digits, symbols int) string {
	var (
		lowerCharSet = "abcdedfghijklmnopqrst"
		upperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digitsSet    = "0123456789"
		symbolsSet   = "!@#$%&*"
		allCharSet   = lowerCharSet + upperCharSet + digitsSet + symbolsSet
	)

	var passwd strings.Builder

	for i := 0; i < lower; i++ {
		random := rand.Intn(len(lowerCharSet))
		passwd.WriteString(string(lowerCharSet[random]))
	}

	for i := 0; i < upper; i++ {
		random := rand.Intn(len(upperCharSet))
		passwd.WriteString(string(upperCharSet[random]))
	}

	for i := 0; i < digits; i++ {
		random := rand.Intn(len(digitsSet))
		passwd.WriteString(string(digitsSet[random]))
	}

	for i := 0; i < symbols; i++ {
		random := rand.Intn(len(symbolsSet))
		passwd.WriteString(string(symbolsSet[random]))
	}

	remaining := length - lower - upper - digits - symbols
	for i := 0; i < remaining; i++ {
		random := rand.Intn(len(allCharSet))
		passwd.WriteString(string(allCharSet[random]))
	}

	inRune := []rune(passwd.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}

func generateCompletionCode() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000)*1000 + rand.Intn(1000)
}

type UserHandler struct {
	Store *store.Store
}

func NewUserHandler(store *store.Store) *UserHandler {
	return &UserHandler{Store: store}
}

type CreateUserRequest struct {
	domain.RequestBody
	ExternalID string `json:"external_id"`
}

type CreateUserResponse struct {
	ID         int    `json:"id"`
	ExternalID string `json:"external_id"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if h == nil {
		render.Render(w, r, NewErrResponseRenderer(nil, http.StatusInternalServerError))
		return
	}

	p := &CreateUserRequest{}
	if err := render.Bind(r, p); err != nil {
		fmt.Println(err)
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	ctx := r.Context()
	secret := generateSecret(36, 6, 6, 6, 6)

	fu, err := h.Store.Auth.RegisterUser(ctx, p.ExternalID, secret)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, store.ErrUIDAlreadyExists) {
			render.Render(w, r, NewErrResponseRenderer(err, http.StatusConflict))
			return
		}
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusInternalServerError))
		return
	}

	user, err := h.Store.User.CreateUser(ctx, fu)
	if err != nil {
		fmt.Println(err)
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusInternalServerError))
		return
	}

	token, err := h.Store.Auth.GenerateToken(ctx, fu.FirebaseUID)
	if err != nil {
		fmt.Println(err)
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusInternalServerError))
		return
	}

	c := domain.CompletionCode{UserID: user.ID, Code: generateCompletionCode()}
	if err := h.Store.User.SetCompletionCode(ctx, &c); err != nil {
		fmt.Println(err)
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusInternalServerError))
		return
	}

	render.Render(w, r, NewResponseRenderer(
		&domain.User{
			ID:         user.ID,
			ExternalID: user.ExternalID,
			Token:      token,
		},
		http.StatusOK,
	))
}

type IssueCompletionCodeRequest struct {
	domain.RequestBody
	UserID int `db:"user_id"`
}

func (h *UserHandler) GetCompletionCode(w http.ResponseWriter, r *http.Request) {
	if h == nil {
		render.Render(w, r, NewErrResponseRenderer(nil, http.StatusInternalServerError))
		return
	}

	ctx := r.Context()
	param := r.URL.Query().Get("user")
	userID, err := strconv.Atoi(param)
	if err != nil {
		render.Render(w, r, NewErrResponseRenderer(ErrInvalidParameter, http.StatusBadRequest))
		return
	}

	code, err := h.Store.User.GetCompletionCode(ctx, userID)
	if err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusInternalServerError))
		return
	}

	render.Render(w, r, NewResponseRenderer(code, http.StatusOK))
}

func (h *UserHandler) IssueCompletionCode(w http.ResponseWriter, r *http.Request) {
	if h == nil {
		render.Render(w, r, NewErrResponseRenderer(nil, http.StatusInternalServerError))
		return
	}

	p := &IssueCompletionCodeRequest{}
	if err := render.Bind(r, p); err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusBadRequest))
		return
	}

	ctx := r.Context()
	code, err := h.Store.User.GetCompletionCode(ctx, p.UserID)
	if err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusInternalServerError))
		return
	}

	if code > 0 {
		render.Render(w, r, NewResponseRenderer(code, http.StatusOK))
		return
	}

	rand.Seed(time.Now().UnixNano())
	code = rand.Intn(100000)

	err = h.Store.User.SetCompletionCode(ctx, &domain.CompletionCode{UserID: p.UserID, Code: code})
	if err != nil {
		render.Render(w, r, NewErrResponseRenderer(err, http.StatusInternalServerError))
		return
	}

	render.Render(w, r, NewResponseRenderer(code, http.StatusOK))
}
