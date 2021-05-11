/*
 * Siren.
 *
 * Documentation of our Siren API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client
import (
	"time"
)

type Template struct {
	Body string `json:"body,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Id int32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Variables []Variable `json:"variables,omitempty"`
}
