package ssh

import (
	"log"

	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

func pathConfigLease(b *backend) *framework.Path {
	log.Printf("Vishal: ssh.pathConfigLease\n")
	return &framework.Path{
		Pattern: "config/lease",
		Fields: map[string]*framework.FieldSchema{
			"lease": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "Default lease for roles.",
			},

			"lease_max": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "Maximum time a credential is valid for.",
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.WriteOperation: b.pathLeaseWrite,
		},

		HelpSynopsis:    pathConfigLeaseHelpSyn,
		HelpDescription: pathConfigLeaseHelpDesc,
	}
}

func (b *backend) pathLeaseWrite(req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	log.Printf("Vishal: ssh.pathLeaseWrite\n")
	return nil, nil
}

const pathConfigLeaseHelpSyn = `
Configure the default lease information for SSH one time keys.
`

const pathConfigLeaseHelpDesc = `
This configures the default lease information used for SSH one time keys
generated by this backend. The lease specifies the duration that a
credential will be valid for, as well as the maximum session for
a set of credentials.

The format for the lease is "1h" or integer and then unit. The longest
unit is hour.
`