/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tsuru

type PlatformInfo struct {
	Platform *Platform `json:"platform,omitempty"`

	Images []string `json:"images,omitempty"`
}