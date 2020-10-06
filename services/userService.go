package services

func UserRegister() (map[string]interface{},error) {
	result := make(map[string]interface{})
	result["name"] = "This is Sai"
	return result, nil
}