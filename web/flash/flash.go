package flash

import (
	"encoding/base64"
	"net/http"
	"time"
)

// Set will save the msg in a cookie
func Set(w http.ResponseWriter, name string, value []byte) {
	c := &http.Cookie{
		Name:     name,
		Value:    encode(value),
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}

// Get will retrive the flash message from cookie
func Get(w http.ResponseWriter, r *http.Request, name string) ([]byte, error) {
	c, err := r.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil, nil
		default:
			return nil, err
		}
	}
	value, err := decode(c.Value)
	if err != nil {
		return nil, err
	}
	dc := &http.Cookie{
		Name:     name,
		HttpOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(1, 0),
	}
	http.SetCookie(w, dc)
	return value, nil
}

// -------------------------

func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
