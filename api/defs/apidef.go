/**
 *  @author: yanKoo
 *  @Date: 2019/3/10 21:40
 *  @Description:
 */
package defs

// requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}
