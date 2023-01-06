package keycloak

import "github.com/Nerzal/gocloak/v7"

type Keycloak struct {
	Gocloak      gocloak.GoCloak // Keycloak client
	ClientId     string          // ClientId specified in Keycloak
	ClientSecret string          // ClientSecret secret specified in Keycloak
	Realm        string          // Realm specified in Keycloak
}

func NewKeycloak() *Keycloak {
	return &Keycloak{
		Gocloak:      gocloak.NewClient("http://localhost:8180"),
		ClientId:     "task-client",
		ClientSecret: "tP9729gPzNk3N6835oKdzUxMgpq0Q8dz",
		Realm:        "task-realm",
	}
}




