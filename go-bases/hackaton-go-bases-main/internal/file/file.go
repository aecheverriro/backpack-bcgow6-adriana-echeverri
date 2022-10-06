package file

import (
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
	"os"
	"strings"
	"strconv"
)

type File struct {
	path string
}

const (
	lengthFile = 1000
)

func (f *File) Read() (tickets [lengthFile]service.Ticket, err error) {
	data, err := os.ReadFile(f.path)

	if err != nil {
		return
	}

	rows := strings.Split(string(data[:]), "\n")

	for index, value := range rows {
		information := strings.Split(string(value[:]), ",")
		id, err := strconv.Atoi(information[0])
		if err != nil {
			return
		}
		price, err := strconv.Atoi(information[5])
		if err != nil {
			return
		}
		tickets[index] = service.Ticket{
			id,
			information[1],
			information[2],
			information[3],
			information[4],
			price,
		}
	}
	
	return
}

func (f *File) Write(ticket service.Ticket) (err error) {
	file, err := os.OpenFile(f.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777))

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
