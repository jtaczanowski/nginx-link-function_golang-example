package main

/*
#include "ngx_link_func_module.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export ngx_link_func_init_cycle
func ngx_link_func_init_cycle(ctx interface{}) {
	fmt.Println("initializing go app")
}

//export ngx_link_func_exit_cycle
func ngx_link_func_exit_cycle(ctx interface{}) {
	fmt.Println("exiting go app")
}

//export add_http_header_to_request
func add_http_header_to_request(ctx *C.ngx_link_func_ctx_t) {
	// add new HTTP header to request headers
	C.ngx_link_func_add_header_in(ctx, C.CString("headerKey"), C.ulong(len("headerKey")), C.CString("headerVal"), C.ulong(len("headerVal")))
}

//export return_custom_http_headers_and_response
func return_custom_http_headers_and_response(ctx *C.ngx_link_func_ctx_t) {
	// get User-Agent header
	userAgent := C.ngx_link_func_get_header(ctx, C.CString("User-Agent"), C.ulong(len("User-Agent")))
	userAgentString := C.GoString((*C.char)(unsafe.Pointer(userAgent)))

	// add HTTP header
	C.ngx_link_func_add_header_out(ctx, C.CString("headerKey"), C.ulong(len("headerKey")), C.CString("headerVal"), C.ulong(len("headerVal")))

	// serve HTTP response
	C.ngx_link_func_write_resp(
		ctx,
		C.ulong(200),
		C.CString("200 OK"),
		C.CString("text/plain"),
		C.CString(userAgentString),
		C.ulong(len(userAgentString)),
	)
}

func main() {
}
