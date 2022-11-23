package models

type Specification struct {
    
    Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
    
}