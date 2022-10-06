package main

import (
	"fmt"
	"errors"
)

func monthSalary(hours float64, hourPay float64) (salary float64, err error){
	if hours < 80 || hourPay < 0 {
		err = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return
	}
	salary = hours * hourPay
	if salary >= 150000 {
		salary -= salary * 0.1
	}
	return
}

func bonus(bestSemesterSalary float64, duration float64, err1 error) (bonus float64, err error){
	if err1 != nil {
		err = fmt.Errorf("%w", err1)
		return
	}
	if bestSemesterSalary < 0 || duration < 0 {
		err = errors.New("error: el trabajador debe tener un salario o duración valida")
		return
	}
	bonus = bestSemesterSalary / 12 * duration
	return
}

func main() {
	var hoursPerMonth float64 = 79
	var hourPay float64 = 9.37
	var bestSalary float64 = 1008
	var duration float64 = 1.5

	salary, err := monthSalary(hoursPerMonth, hourPay)
	bonus, err1 := bonus(bestSalary, duration, err)

	if err == nil && err1 == nil {
		fmt.Printf("El salario es de %.2f, con un bono de %.2f, dando un total de %.2f \n", salary, bonus, salary + bonus)
	} else if errors.Unwrap(err1) == nil {
		fmt.Println(err1)
	} else {
		fmt.Println(err1)
	}
}
