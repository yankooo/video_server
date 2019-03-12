/**
 *  @author: yanKoo
 *  @Date: 2019/3/10 21:41
 *  @Description:
 */
package defs

// 错误结构体
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"` // 用来查询每个错误的方式
}

type ErrorResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC: 400, Error: Err{Error: "Request body is not correct.", ErrorCode: "001"}} // 不能解析这个请求
	ErrorNotAuthUser            = ErrorResponse{HttpSC: 401, Error: Err{Error: "User authentication failed.", ErrorCode: "002"}}  // 用户不合法，不存在
	ErrorDBError                = ErrorResponse{HttpSC: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}                // 数据库操作错误
	ErrorInternalFaults         = ErrorResponse{HttpSC: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
