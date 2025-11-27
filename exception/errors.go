package exception

import "errors"

var (
    ErrOutboundAdd         = errors.New("stok tidak mencukupi untuk outbound")
    ErrProductId         = errors.New("produk tidak ditemukan")
    ErrProductSku         = errors.New("SKU tidak boleh kosong")
)
