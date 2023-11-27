package config

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	userOperations "web-push/internal/database"
	"web-push/internal/utils"
	restOperations "web-push/pkg/controller"
)

type Routing struct {
	UserOperationsSrv userOperations.UserOperations
}

func (r Routing) InitRouting(ctx context.Context) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/users/create", restOperations.SaveUser(ctx, r.UserOperationsSrv)).Methods("POST", "OPTIONS")

	err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"POST", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "App-Version", "Device-Language", "Device-Name", "Device-ID", "User-Agent", "Referer"}),
	)(router)))

	if err != nil {
		utils.LogFatalAndStop(ctx, utils.GetLogDetails("[APP_ROUTING]", "[INIT_ROUTING]", fmt.Errorf("error at listen and serve: %w", err).Error()))
	}
}
