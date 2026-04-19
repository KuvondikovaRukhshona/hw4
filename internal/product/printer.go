package product

import (
	"fmt"
	"strconv"
	"strings"
)

type Product struct {
	Name    string
	Brand   string
	Price   int
	InStock bool
}

const (
	template = `===== Alifshop =====
Товар:    {name}
Бренд:    {brand}
Цена:     {price} сум
В наличии: {inStock}
Рассрочка: 12 мес → {firstAmount} сум/мес
====================`
)

func formatSum(tiyins int) string {
	sums := tiyins / 100
	remainder := tiyins % 100

	s := strconv.Itoa(sums)

	var buf []byte
	for i, ch := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			buf = append(buf, ' ')
		}
		buf = append(buf, byte(ch))
	}
	if remainder == 0 {
		return string(buf)
	}
	return fmt.Sprintf("%s.%02d", string(buf), remainder)
}
func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", formatSum(product.Price))
	text = strings.ReplaceAll(text, "{firstAmount}", formatSum(firstAmount))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	return text
}
