package pro

func MD5Encode(password string) string

func HashPassword(password string) (string, error)

func CheckPasswordHash(password, hash string) bool
