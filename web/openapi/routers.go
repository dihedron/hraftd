/*
 * Brokerd API
 *
 * This API allows to interact with the `brokerd` daemon in its three different capacities:  1. as a **key/value store for properties**, holding key/value pairs; 2. as a **Raft cluster member**, through cluster managemenet APIs thata llow to check the state and   interact with the cluster (e.g. moving the master to another node, forcing a sync-up of the   cluster nodes, etc.) 3. as a **relational store for virtual machines and network ports**, as reported by OpenStack via its    notification exchanges inside RabbitMQ.  The **Properties API** allows to manage the lifecycle of key/value pairs; the kes part can encode a pseudo-hierarchical organisation of the information by adopting conventional characters as field separators, the way it is usually done in Java properties files, e.g.: ``` key-part-1.key-part-2.key-part-3....key-part-N=value ```  where each part of the key (`key-part-X`) can encode some part of a taxonomy.  The **Cluster API** provides a way to interact with the Raft cluster that guarantees that the Finite State Machines (FSM) holding the state of the several `brokerd` instances running on different  OpenStack controller nodes are all kept in sync and moving in lock-step. Through the API it is possible  to check the health and the status (*leader*, *follower*) of the nodes in the cluster, move the cluster leadership from the current master to a different node, force a sync-up of the cluster nodes, trigger the snaphotting of the current FSM state, etc.  The **Store API** provides a way to interact with the SQLite database holding information about  Virtual Machines and Network Ports.
 *
 * API version: 1.0
 * Contact: support@example.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/api/v1/",
		Index,
	},

	{
		"ListNodes",
		http.MethodGet,
		"/api/v1/cluster/nodes",
		ListNodes,
	},

	{
		"CreateProperty",
		http.MethodPost,
		"/api/v1/properties",
		CreateProperty,
	},

	{
		"DeleteProperties",
		http.MethodDelete,
		"/api/v1/properties",
		DeleteProperties,
	},

	{
		"DeleteProperty",
		http.MethodDelete,
		"/api/v1/settings/:key",
		DeleteProperty,
	},

	{
		"GetProperty",
		http.MethodGet,
		"/api/v1/settings/:key",
		GetProperty,
	},

	{
		"ListProperties",
		http.MethodGet,
		"/api/v1/properties",
		ListProperties,
	},

	{
		"UpdateProperty",
		http.MethodPut,
		"/api/v1/settings/:key",
		UpdateProperty,
	},
}