package DTO

import "../../../bin/gorm"

type MessageDTO struct
{
	gorm.Model
	Data string

}

func  ( dto MessageDTO)  TableName() string {

	return "messages_consume"
}
