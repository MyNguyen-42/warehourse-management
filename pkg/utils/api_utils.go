package utils

//func GetPaging(paging common.Paging) common.Paging {
//	if paging.Limit > 1000 || paging.Limit <= 0 {
//		paging.Limit = int64(100)
//	}
//	if paging.Offset <= 0 {
//		paging.Offset = int64(0)
//	}
//
//	return paging
//}
//
//func ParseJsonRequest(ctx echo.Context, out interface{}) error {
//	var jsonMap = map[string]interface{}{}
//	err := ctx.Bind(&jsonMap)
//	if err != nil {
//		// if requested body is empty and it's a request for list, return default
//		if ctx.Request().ContentLength == 0 {
//			return nil
//		}
//		return errcode.Error(errcode.InvalidRequest)
//	}
//	var requestDataByte []byte
//	requestDataByte, _ = json.Marshal(jsonMap)
//
//	err = json.Unmarshal(requestDataByte, out)
//	if err != nil {
//		return errcode.Error(errcode.InvalidRequest)
//	}
//	return nil
//}
