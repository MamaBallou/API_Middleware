package middleware_api

import (
	"middleware/api/utils"
    "net/http"
)

func JwtMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        requestType := utils.ExtractRequestType(r)

        if requestType != "Connexion"{
            // Récupérer le token JWT de l'en-tête "Authorization".
            tokenString := utils.ExtractBearerToken(r)

            // Vérifier la validité du token JWT.
            if tokenString == "" {
                utils.ToJson(w, "JWT missing", http.StatusNotAcceptable)
                return
            }

            claims, err := utils.JwtExtract(r)
            if err != nil {
                utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
                return
            }

            if application_type, ok := claims["application_type"].(string); ok {
                if application_type != "Consultation des freezebee"{
                    utils.ToJson(w, "Application Type incorrect", http.StatusUnprocessableEntity)
                    return
                }
            }

            if application_version, ok := claims["application_version"].(string); ok {
                if application_version != "1.0.0"{
                    utils.ToJson(w, "Application Version incorrect", http.StatusUnprocessableEntity)
                    return
                }
            }

            // Vérifier si l'email dans le token correspond à un JWT stocké dans la base de données.
            if user_permission, ok := claims["user_permission"].(string); ok {
                if user_permission == "Standard" && (r.Method == "POST" || r.Method == "PATCH" || r.Method == "DELETE"){
                    utils.ToJson(w, "Unauthorized to do this method", http.StatusUnprocessableEntity)
                    return
                }
            }
        }

        // Si l'accès est autorisé, passez la requête au gestionnaire de route suivant.
        next.ServeHTTP(w, r)
    })
}
