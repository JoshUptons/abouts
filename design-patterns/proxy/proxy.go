package proxy

/*
A proxy a node placed in between the requester, and the resource requested.
The proxy node acts as a filter, firewall, and caching point for the endpoint.
This keeps traffic limited to valid requests via the firewall, could act as a
rate limiter or load balancer for traffic to be moved between free servers.
Or could act as a caching layer to help keep frequent, static requests from
bogging down the resource's endpoint.
*/

/*
Example:
We have an API that is interacting with a database to store user information.
We can use a proxy handler to control flow to and from the database, as well as
cache some of the common requests.
*/

/*
I will create my own request structures for simplicity.
*/

var (
	allowedOrigins = myArr{
		Items: []string{
			"localhost",
			"google.com",
		},
	}
	allowedMethods = myArr{
		Items: []string{
			"GET",
		},
	}
	bannedOrigins = myArr{
		Items: []string{
			"banned.com",
		},
	}
)

type myArr struct {
	Items []string
}

type myHeaders struct {
	Origin      string
	Method      string
	ContentType string `json:"Content-Type" url:"Content-Type"`
}

type myRequest struct {
	Headers myHeaders
	Body    string
}

/*
a handler function would not really return a string, but just for simplicity
I am only going to demonstrate how a handler might process a request.
The function myHandler acts as a proxy between a requester and the resource,
in this case a database
*/

func myHandler(r *myRequest) string {
	verified := verifyHeaders(r.Headers)
	if verified == false {
		return "404"
	}

	var data string

	switch r.Headers.Method {
	// we only allow get requests, but I'll put this in just as a placeholder
	case "GET":
		/*
			mock database call
			would return the appropriate response
			if the object in the database cannot be
			accessed or not found etc.
		*/

		// return db.get(r.Body)
	}

	return data
}

func verifyHeaders(h myHeaders) bool {
	// simple verificaction layer
	if h.ContentType != "application/json" {
		return false
	}

	if isNotIn(h.Method, &allowedMethods) {
		return false
	}

	if isNotIn(h.Origin, &allowedOrigins) {
		return false
	}

	return true
}

func isNotIn(item string, arr *myArr) bool {
	for _, i := range arr.Items {
		if i == item {
			return false
		}
	}
	return true
}
