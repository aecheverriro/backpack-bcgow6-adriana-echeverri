package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

const (
	lengthFile = 1000
)

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

type File struct {
	path string
}

func (f *File) Read() ([lengthFile]Ticket, error){
	data, err := os.ReadFile(f.path)
	var tickets [lengthFile]Ticket
	
	if err == nil {
		rows := strings.Split(string(data[:]), "\n")

		for index, value := range rows {
			information := strings.Split(string(value[:]), ",")
			id := information[0]
			price := information[5]
			idFormatted, err1 := strconv.Atoi(id)
			priceFormatted, err2 := strconv.Atoi(price)

			if err1 != nil {
				return tickets, err1
			}
			if err2 != nil {
				return tickets, err2
			}

			tickets[index] = Ticket{
				idFormatted,
				information[1],
				information[2],
				information[3],
				information[4],
				priceFormatted,
			}
		}
	}
	
	return tickets, err
}

func (f *File) Write(ticket Ticket) (err error) {
	file, err := os.OpenFile(f.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		return
	}

	content := fmt.Sprintf(
		"\n%v,%v,%v,%v,%v,%v",
		ticket.Id,
		ticket.Names,
		ticket.Email,
		ticket.Destination,
		ticket.Date,
		ticket.Price)
	_, err = file.WriteString(content)
	return 
}

func main() {
	file := File{"./tickets.csv"}	
	data, err := file.Read()
	fmt.Println(data[0].Id)
	fmt.Println(err)
}