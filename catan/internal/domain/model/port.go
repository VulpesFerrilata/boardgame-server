package model

import "gorm.io/gorm"

type Port struct {
	gorm.Model
	Q        int      
	R        int      
	PortType FieldType 
	FieldQ   int      
	FieldR   int    
}