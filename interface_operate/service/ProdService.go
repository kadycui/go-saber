package service

import "log"

type ProdService struct {
}

func NewProdService() *ProdService {
	return &ProdService{}
}

func (this *ProdService) Save() {
	log.Println("商品保存成功")
}
