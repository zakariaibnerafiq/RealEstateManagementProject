package user

import (
	"fmt"
	"learninggo/config"
	"learninggo/service/auth"
	"learninggo/types"
	"learninggo/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/users/{id}", auth.WithJWTAuth(h.handleGetUser, h.store)).Methods(http.MethodGet)
	router.HandleFunc("/update/email", auth.WithJWTAuth(h.handleUpdateEmail, h.store)).Methods(http.MethodPut)
	router.HandleFunc("/update/username", auth.WithJWTAuth(h.handleUpdateUserName, h.store)).Methods(http.MethodPut)
	router.HandleFunc("/update/password", auth.WithJWTAuth(h.handleUpdatePassword, h.store)).Methods(http.MethodPut)
	router.HandleFunc("/update/userimage", auth.WithJWTAuth(h.handleUpdateUserImage, h.store)).Methods(http.MethodPut)
	router.HandleFunc("/update/phone", auth.WithJWTAuth(h.handleUpdatePhone, h.store)).Methods(http.MethodPut)
	router.HandleFunc("/admin/users", auth.WithJWTAuth(h.handleGetAllUsers, h.store)).Methods(http.MethodGet)
	router.HandleFunc("/dashboard/getuserdocument/{id}", auth.WithJWTAuth(h.handleDashGetDocument, h.store)).Methods("GET")
	router.HandleFunc("/dashboard/submituserdocument", auth.WithJWTAuth(h.handleDashSubmitDocument, h.store)).Methods("POST")
	router.HandleFunc("/dashboard/verifyuser", auth.WithJWTAuth(h.handleVerifyUser, h.store)).Methods("POST")
	router.HandleFunc("/dashboard/pendingusers/{page}", auth.WithJWTAuth(h.handleGetPendingUsers, h.store)).Methods("GET")

	
	// router.HandleFunc("/test", h.handletest).Methods("POST")

}

func (h *Handler) handleUpdateUserImage(w http.ResponseWriter, r *http.Request) {

	var payload types.UpdateUserImagePayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	dir := "./uploads/profile"
	user, err := h.store.GetUserByID(payload.ID)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	if user.ImagePath != "null" {
		err = utils.DeleteImage(user.ImagePath, dir)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	}

	// Decode Base64 image data
	fileName, err := utils.SaveImage(payload.Image, dir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.store.UpdateUserImage(payload.ID, fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
	fmt.Printf("File uploaded successfully: %s\n", fileName)

}
func (h *Handler) handleUpdateUserName(w http.ResponseWriter, r *http.Request) {
	contextValues := r.Context().Value(auth.UserKey).(types.UserContext)
	userID := contextValues.ID
	userRole := contextValues.Role

	var user types.UpdateUserNamePayload

	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Fetch user from store to check existing data and for validation
	u, err := h.store.GetUserByID(user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	// Validate payload structure
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Ensure the user performing the update is authorized
	if user.ID != userID {
		if userRole != "admin" {
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("forbidden"))
			return
		} else if !auth.ComparePasswords(u.Password, user.Password) {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid password"))
			return
		}
	}

	// Update the user's name
	err = h.store.UpdateUserName(user.ID, user.Name)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUpdateEmail(w http.ResponseWriter, r *http.Request) {
	contextValues := r.Context().Value(auth.UserKey).(types.UserContext)
	userID := contextValues.ID
	userRole := contextValues.Role

	var user types.UpdateUserEmailPayload

	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Fetch user from store to check existing data and for validation
	u, err := h.store.GetUserByID(user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	// Validate payload structure
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Ensure the user performing the update is authorized
	if user.ID != userID {
		if userRole != "admin" {
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("forbidden"))
			return
		} else if !auth.ComparePasswords(u.Password, user.Password) {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid password"))
			return
		}
	}

	// Update the user's name
	err = h.store.UpdateUserEmail(user.ID, user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUpdatePassword(w http.ResponseWriter, r *http.Request) {
	contextValues := r.Context().Value(auth.UserKey).(types.UserContext)
	userID := contextValues.ID
	userRole := contextValues.Role

	var user types.UpdateUserPasswordPayload

	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Fetch user from store to check existing data and for validation
	u, err := h.store.GetUserByID(user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	// Validate payload structure
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Ensure the user performing the update is authorized
	if user.ID != userID {
		if userRole != "admin" {
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("forbidden"))
			return
		} else if !auth.ComparePasswords(u.Password, user.OldPassword) {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid password"))
			return
		}
	}

	// Update the user's name
	err = h.store.UpdateUserPassword(user.ID, user.NewPassword)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleUpdatePhone(w http.ResponseWriter, r *http.Request) {
	contextValues := r.Context().Value(auth.UserKey).(types.UserContext)
	userID := contextValues.ID
	userRole := contextValues.Role

	var user types.UpdateUserPhonePayload

	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Fetch user from store to check existing data and for validation
	u, err := h.store.GetUserByID(user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	// Validate payload structure
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Ensure the user performing the update is authorized
	if user.ID != userID {
		if userRole != "admin" {
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("forbidden"))
			return
		} else if !auth.ComparePasswords(u.Password, user.Password) {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid password"))
			return
		}
	}

	// Update the user's name
	err = h.store.UpdateUserPhone(user.ID, user.Phone)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
// admin privilages
func (h *Handler) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	contextValues := r.Context().Value(auth.UserKey).(types.UserContext)
	userRole := contextValues.Role

	if userRole != "admin" {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("forbidden"))
		return
	}
	users, err := h.store.GetAllUsers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, users)
}
// user privilages

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	// decode token and get user id

	contextValues := r.Context().Value(auth.UserKey).(types.UserContext)

	userID := contextValues.ID
	userRole := contextValues.Role
	log.Println(userRole)
	vars := mux.Vars(r)

	id := vars["id"]
	log.Println("ID", id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if idInt != userID {
		if userRole != "admin" {
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("forbidden"))
			return
		}
	}
	user, err := h.store.GetUserByID(idInt)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}
	log.Println("User", user.Email)
	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUserPayload

	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	log.Println("User", user)
	u, err := h.store.GetUserByEmail(user.Email)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePasswords(u.Password, user.Password) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})

	log.Println("Login handler")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get json payload
	var user types.RegisterUserPayload

	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	//  check if the email is already registered
	_, err := h.store.GetUserByEmail(user.Email)

	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	log.Println("Register handler")
	// if not, create the user

	err = h.store.CreateUser(&types.User{
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		Phone:     user.Phone,
		Role:      "client",
		ImagePath: "null",
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// return success
	utils.WriteJSON(w, http.StatusCreated, nil)

}

func (h *Handler) handleGetPendingUsers (w http.ResponseWriter, r *http.Request) {
	contextValues := r.Context().Value(auth.UserKey).(types.UserContext)
	if contextValues.Role != "admin" {
		http.Error(w, "you must be an admin to view all properties", http.StatusUnauthorized)
		return
	}

	page, err := strconv.Atoi(mux.Vars(r)["page"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid page number"))
		return
	}


	users,count, err := h.store.GetAllPendingUsers(page)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}


	// Prepare the response
	response := map[string]interface{}{
		"count": count,
		"users": users,
	}
	utils.WriteJSON(w, http.StatusOK, response)
}
