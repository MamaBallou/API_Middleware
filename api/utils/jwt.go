package utils

import (
    "strings"
    "net/http"
    "middleware/config"
    jwt "github.com/dgrijalva/jwt-go"
)

// Clé secrète pour la génération et la validation des tokens JWT.
var secretKey = config.JwtSecretKey

// Extrait les claims (informations) du JWT à partir de la requête HTTP.
func JwtExtract(r *http.Request) (map[string]interface{}, error) {
    tokenString := ExtractBearerToken(r)
    claims := jwt.MapClaims{}
    _, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error) {
        return secretKey, nil
    })
    if err != nil {
        return nil, err
    }
    return claims, nil
}

// Extrait le token JWT du format "Bearer" à partir de l'en-tête HTTP.
func ExtractBearerToken(r *http.Request) string {
    headerAuthorization := r.Header.Get("Authorization")
    bearerToken := strings.Split(headerAuthorization, " ")
    if len(bearerToken) == 2 {
        return bearerToken[1]
    }
    return ""
}

// Extrait la session utilisateur à partir de l'en-tête HTTP.
func ExtractRequestType(r *http.Request) string {
    headerUserSession := r.Header.Get("RequestType")
    return headerUserSession
}