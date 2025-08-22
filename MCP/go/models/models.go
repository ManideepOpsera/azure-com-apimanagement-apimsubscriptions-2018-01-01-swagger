package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// SubscriptionContractProperties represents the SubscriptionContractProperties schema from the OpenAPI specification
type SubscriptionContractProperties struct {
	Secondarykey string `json:"secondaryKey"` // Subscription secondary key.
	Startdate string `json:"startDate,omitempty"` // Subscription activation date. The setting is for audit purposes only and the subscription is not automatically activated. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	State string `json:"state"` // Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	Statecomment string `json:"stateComment,omitempty"` // Optional subscription comment added by an administrator.
	Enddate string `json:"endDate,omitempty"` // Date when subscription was cancelled or expired. The setting is for audit purposes only and the subscription is not automatically cancelled. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Expirationdate string `json:"expirationDate,omitempty"` // Subscription expiration date. The setting is for audit purposes only and the subscription is not automatically expired. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Primarykey string `json:"primaryKey"` // Subscription primary key.
	Userid string `json:"userId"` // The user resource identifier of the subscription owner. The value is a valid relative URL in the format of /users/{uid} where {uid} is a user identifier.
	Createddate string `json:"createdDate,omitempty"` // Subscription creation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Displayname string `json:"displayName,omitempty"` // The name of the subscription, or null if the subscription has no name.
	Notificationdate string `json:"notificationDate,omitempty"` // Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Productid string `json:"productId"` // The product resource identifier of the subscribed product. The value is a valid relative URL in the format of /products/{productId} where {productId} is a product identifier.
}

// SubscriptionCreateParameterProperties represents the SubscriptionCreateParameterProperties schema from the OpenAPI specification
type SubscriptionCreateParameterProperties struct {
	Primarykey string `json:"primaryKey,omitempty"` // Primary subscription key. If not specified during request key will be generated automatically.
	Productid string `json:"productId"` // Product (product id path) for which subscription is being created in form /products/{productId}
	Secondarykey string `json:"secondaryKey,omitempty"` // Secondary subscription key. If not specified during request key will be generated automatically.
	State string `json:"state,omitempty"` // Initial subscription state. If no value is specified, subscription is created with Submitted state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	Userid string `json:"userId"` // User (user id path) for whom subscription is being created in form /users/{uid}
	Displayname string `json:"displayName"` // Subscription name.
}

// SubscriptionCreateParameters represents the SubscriptionCreateParameters schema from the OpenAPI specification
type SubscriptionCreateParameters struct {
	Properties SubscriptionCreateParameterProperties `json:"properties,omitempty"` // Parameters supplied to the Create subscription operation.
}

// SubscriptionUpdateParameterProperties represents the SubscriptionUpdateParameterProperties schema from the OpenAPI specification
type SubscriptionUpdateParameterProperties struct {
	Userid string `json:"userId,omitempty"` // User identifier path: /users/{uid}
	Displayname string `json:"displayName,omitempty"` // Subscription name.
	Expirationdate string `json:"expirationDate,omitempty"` // Subscription expiration date. The setting is for audit purposes only and the subscription is not automatically expired. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Primarykey string `json:"primaryKey,omitempty"` // Primary subscription key.
	Productid string `json:"productId,omitempty"` // Product identifier path: /products/{productId}
	Secondarykey string `json:"secondaryKey,omitempty"` // Secondary subscription key.
	State string `json:"state,omitempty"` // Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	Statecomment string `json:"stateComment,omitempty"` // Comments describing subscription state change by the administrator.
}

// SubscriptionUpdateParameters represents the SubscriptionUpdateParameters schema from the OpenAPI specification
type SubscriptionUpdateParameters struct {
	Properties SubscriptionUpdateParameterProperties `json:"properties,omitempty"` // Parameters supplied to the Update subscription operation.
}

// SubscriptionCollection represents the SubscriptionCollection schema from the OpenAPI specification
type SubscriptionCollection struct {
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
	Value []SubscriptionContract `json:"value,omitempty"` // Page values.
}

// SubscriptionContract represents the SubscriptionContract schema from the OpenAPI specification
type SubscriptionContract struct {
	Name string `json:"name,omitempty"` // Resource name.
	TypeField string `json:"type,omitempty"` // Resource type for API Management resource.
	Id string `json:"id,omitempty"` // Resource ID.
}
