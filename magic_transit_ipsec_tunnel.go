package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

// Magic Transit IPsec Tunnel Error messages.
const (
	errMagicTransitIPsecTunnelNotModified = "When trying to modify IPsec tunnel, API returned modified: false"
	errMagicTransitIPsecTunnelNotDeleted  = "When trying to delete IPsec tunnel, API returned deleted: false"
)

// MagicTransitIPsecTunnel contains information about an IPsec tunnel.
type MagicTransitIPsecTunnel struct {
	ID                 string     `json:"id,omitempty"`
	CreatedOn          *time.Time `json:"created_on,omitempty"`
	ModifiedOn         *time.Time `json:"modified_on,omitempty"`
	Name               string     `json:"name"`
	CustomerEndpoint   string     `json:"customer_endpoint"`
	CloudflareEndpoint string     `json:"cloudflare_endpoint"`
	InterfaceAddress   string     `json:"interface_address"`
	Description        string     `json:"description,omitempty"`
}

// ListMagicTransitIPSecTunnelsResponse contains a response including IPsec tunnels.
type ListMagicTransitIPSecTunnelsResponse struct {
	Response
	Result struct {
		IPsecTunnels []MagicTransitIPsecTunnel `json:"ipsec_tunnels"`
	} `json:"result"`
}

// GetMagicTransitIPSecTunnelResponse contains a response including zero or one IPsec tunnels.
type GetMagicTransitIPSecTunnelResponse struct {
	Response
	Result struct {
		IPsecTunnel MagicTransitIPsecTunnel `json:"ipsec_tunnel"`
	} `json:"result"`
}

// CreateMagicTransitIPsecTunnelsRequest is an array of IPsec tunnels to create.
type CreateMagicTransitIPsecTunnelsRequest struct {
	IPsecTunnels []MagicTransitIPsecTunnel `json:"ipsec_tunnels"`
}

// UpdateMagicTransitIPSecTunnelResponse contains a response after updating an IPsec Tunnel.
type UpdateMagicTransitIPsecTunnelResponse struct {
	Response
	Result struct {
		Modified            bool                    `json:"modified"`
		ModifiedIPsecTunnel MagicTransitIPsecTunnel `json:"modified_ipsec_tunnel"`
	} `json:"result"`
}

// DeleteMagicTransitIPsecTunnelResponse contains a response after deleting an IPsec Tunnel.
type DeleteMagicTransitIPsecTunnelResponse struct {
	Response
	Result struct {
		Deleted            bool                    `json:"deleted"`
		DeletedIPsecTunnel MagicTransitIPsecTunnel `json:"deleted_ipsec_tunnel"`
	} `json:"result"`
}

// ListMagicTransitIPsecTunnels lists all IPsec tunnels for a given account
//
// API reference: https://api.cloudflare.com/#magic-ipsec-tunnels-list-ipsec-tunnels
func (api *API) ListMagicTransitIPsecTunnels(ctx context.Context) ([]MagicTransitIPsecTunnel, error) {
	if err := api.checkAccountID(); err != nil {
		return []MagicTransitIPsecTunnel{}, err
	}

	uri := fmt.Sprintf("/accounts/%s/magic/ipsec_tunnels", api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []MagicTransitIPsecTunnel{}, err
	}

	result := ListMagicTransitIPSecTunnelsResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []MagicTransitIPsecTunnel{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result.IPsecTunnels, nil
}

// GetMagicTransitIPsecTunnel returns zero or one IPsec tunnel
//
// API reference: https://api.cloudflare.com/#magic-ipsec-tunnels-ipsec-tunnel-details
func (api *API) GetMagicTransitIPsecTunnel(ctx context.Context, id string) (MagicTransitIPsecTunnel, error) {
	if err := api.checkAccountID(); err != nil {
		return MagicTransitIPsecTunnel{}, err
	}

	uri := fmt.Sprintf("/accounts/%s/magic/ipsec_tunnels/%s", api.AccountID, id)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return MagicTransitIPsecTunnel{}, err
	}

	result := GetMagicTransitIPSecTunnelResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return MagicTransitIPsecTunnel{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result.IPsecTunnel, nil
}

// CreateMagicTransitIPsecTunnels creates one or more IPsec tunnels
//
// API reference: https://api.cloudflare.com/#magic-ipsec-tunnels-create-ipsec-tunnels
func (api *API) CreateMagicTransitIPsecTunnels(ctx context.Context, tunnels []MagicTransitIPsecTunnel) ([]MagicTransitIPsecTunnel, error) {
	if err := api.checkAccountID(); err != nil {
		return []MagicTransitIPsecTunnel{}, err
	}

	uri := fmt.Sprintf("/accounts/%s/magic/ipsec_tunnels", api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, CreateMagicTransitIPsecTunnelsRequest{
		IPsecTunnels: tunnels,
	})

	if err != nil {
		return []MagicTransitIPsecTunnel{}, err
	}

	result := ListMagicTransitIPSecTunnelsResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []MagicTransitIPsecTunnel{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result.IPsecTunnels, nil
}

// UpdateMagicTransitIPsecTunnel updates an IPsec tunnel
//
// API reference: https://api.cloudflare.com/#magic-ipsec-tunnels-update-ipsec-tunnel
func (api *API) UpdateMagicTransitIPsecTunnel(ctx context.Context, id string, tunnel MagicTransitIPsecTunnel) (MagicTransitIPsecTunnel, error) {
	if err := api.checkAccountID(); err != nil {
		return MagicTransitIPsecTunnel{}, err
	}

	uri := fmt.Sprintf("/accounts/%s/magic/ipsec_tunnels/%s", api.AccountID, id)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, tunnel)

	if err != nil {
		return MagicTransitIPsecTunnel{}, err
	}

	result := UpdateMagicTransitIPsecTunnelResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return MagicTransitIPsecTunnel{}, errors.Wrap(err, errUnmarshalError)
	}

	if !result.Result.Modified {
		return MagicTransitIPsecTunnel{}, errors.New(errMagicTransitIPsecTunnelNotModified)
	}

	return result.Result.ModifiedIPsecTunnel, nil
}

// DeleteMagicTransitIPsecTunnel deletes an IPsec Tunnel
//
// API reference: https://api.cloudflare.com/#magic-ipsec-tunnels-delete-ipsec-tunnel
func (api *API) DeleteMagicTransitIPsecTunnel(ctx context.Context, id string) (MagicTransitIPsecTunnel, error) {
	if err := api.checkAccountID(); err != nil {
		return MagicTransitIPsecTunnel{}, err
	}

	uri := fmt.Sprintf("/accounts/%s/magic/ipsec_tunnels/%s", api.AccountID, id)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)

	if err != nil {
		return MagicTransitIPsecTunnel{}, err
	}

	result := DeleteMagicTransitIPsecTunnelResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return MagicTransitIPsecTunnel{}, errors.Wrap(err, errUnmarshalError)
	}

	if !result.Result.Deleted {
		return MagicTransitIPsecTunnel{}, errors.New(errMagicTransitIPsecTunnelNotDeleted)
	}

	return result.Result.DeletedIPsecTunnel, nil
}
