package transport 

import (
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler( ep endpoints.Set) http.Handler {
	m := http.NewServeMux() 

	m.Handle("/healthz" , httptransport.NewServer(
		ep.ServiceStatusEndpoint , 
		decodeHTTPServiceStatusRequest , 
		encodeResponse
	))

	m.Handle("/get" , httptransport.NewServer(
		ep.GetEndpoint  ,
		decodeHTTPGetRequest , 
		encodeResponse
	))

	m.Handle("/status" , httptransport.NewServer(
		ep.StatusEndpoint , 
		decodeHTTPStatusRequest , 
		encodeResponse
	))

	m.Handle("/watermark" , httptransport.NewServer(
		ep.WatermarkEndpoint , 
		decodeHTTPWatermarkRequest ,
		encodeResponse 
	))

	m.Handle("/addDocument" , httptransport.NewServer(
		ep.AddDocumentEndpoint , 
		decodeHTTPAddDocumentRequest , 
		encodeResponse 
	))

	return m 
}


func decodeHTTPGetRequest(_ context.Context , r *http.Request) (interface{} , error){
var req endpoints.GetRequest 

if r.ContentLength == 0 {
	logger.Log("Get request with no body")
	return req , nil 
}

err := json.NewDecoder(r.Body).Decode(&req)

if err != nil {
return nil , err 
}

return req , nil 

}

func decodeHTTPStatusRequest(ctx context.Context , r *http.Request) (interface{} , error) {
	var req endpoints.StatusRequest 

	err := json.NewDecoder(r.Body).Decode(&req) 

	if err != nil {
		return nil , err 
	}

	return req , nil 
}

func decodeHTTPWatermarkRequest(ctx context.Context , r *http.Request) (interface{} , error) {
	var req endpoints.WatermarkRequest 

	err := json.NewDecoder(r.Body).Decode(&req) 

	if err != nil {
		return nil , err 
	}

	return req , nil 
}

func decodeHTTPAddDocumentRequest(ctx context.Context , r *http.Request) (interface{} , error) {
	var req endpoints.AddDocumentRequest 

	err := json.NewDecoder(r.Body).Decode(&req) 

	if err != nil {
		return nil , err 
	}

	return req , nil 
}

func decodeHTTPServiceStatusRequest(ctx context.Context , r *http.Request) (interface{} , error) {
	var req endpoints.ServiceStatusRequest 

	err := json.NewDecoder(r.Body).Decode(&req) 

	if err != nil {
		return nil , err 
	}

	return req , nil 
}

func encodeResponse(ctx context.Context , w http.ResponseWriter , response interface{}) error {
	if e , ok := response.(error) ; ok && e != nil {
		encodeError(ctx , e , w )
		return nil 
	}

	return json.NewEncoder(w).Encode(response)
}


func encodeError(ctx context.Context , e error  , w http.ResponseWriter) {
	w.Header().Set("Content-Type" , "application/json; charset=utf-8")

    switch e  {
	case util.ErrUnknown :
		w.WriteHeader(http.StatusNotFound)
	case util.ErrInvalidArgument : 
	w.WriteHeader(http.StatusBadRequest)
	default : 
	w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error" : e.Error() , 
	})
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}