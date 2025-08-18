package constants

import "net/http"

const CookieName string = "token"
const CookieSameSite http.SameSite = http.SameSiteStrictMode
const CookiePath string = "/api"
