package cmds

import (
	"fmt"
	"net"
	"os"

	"github.com/appscode/go/log"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"gomodules.xyz/cert"
	"gomodules.xyz/cert/certstore"
)

func NewCmdCreateServer(certDir string) *cobra.Command {
	var (
		org    []string
		prefix string
		sans   = cert.AltNames{
			IPs: []net.IP{net.ParseIP("127.0.0.1")},
		}
		overwrite bool
	)
	cmd := &cobra.Command{
		Use:               "server-cert",
		Short:             "Generate server certificate pair",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 1 {
				log.Fatalln("Multiple server name found.")
			}
			if len(args) == 0 {
				sans.DNSNames = merge("server", sans.DNSNames)
			} else if len(args) == 1 {
				sans.DNSNames = merge(args[0], sans.DNSNames)
			}

			cfg := cert.Config{
				AltNames: sans,
			}

			store, err := certstore.NewCertStore(afero.NewOsFs(), certDir, org...)
			if err != nil {
				fmt.Printf("Failed to create certificate store. Reason: %v.", err)
				os.Exit(1)
			}

			var p []string
			if prefix != "" {
				p = append(p, prefix)
			}
			if store.IsExists(Filename(cfg), p...) && !overwrite {
				fmt.Printf("Server certificate found at %s. Do you want to overwrite?", store.Location())
				os.Exit(1)
			}

			if err = store.LoadCA(p...); err != nil {
				fmt.Printf("CA certificates not found in %s. Run `init ca`", store.Location())
				os.Exit(1)
			}

			crt, key, err := store.NewServerCertPair(cfg.AltNames)
			if err != nil {
				fmt.Printf("Failed to generate server certificate pair. Reason: %v.", err)
				os.Exit(1)
			}
			err = store.Write(Filename(cfg), crt, key)
			if err != nil {
				fmt.Printf("Failed to init server certificate pair. Reason: %v.", err)
				os.Exit(1)
			}
			fmt.Println("Wrote server certificates in ", store.Location())
			os.Exit(0)
		},
	}

	cmd.Flags().StringVar(&certDir, "cert-dir", certDir, "Path to directory where pki files are stored.")
	cmd.Flags().IPSliceVar(&sans.IPs, "ips", sans.IPs, "Alternative IP addresses")
	cmd.Flags().StringSliceVar(&sans.DNSNames, "domains", sans.DNSNames, "Alternative Domain names")
	cmd.Flags().StringSliceVarP(&org, "organization", "o", org, "Name of client organizations.")
	cmd.Flags().StringVarP(&prefix, "prefix", "p", prefix, "Prefix added to certificate files")
	cmd.Flags().BoolVar(&overwrite, "overwrite", overwrite, "Overwrite existing cert/key pair")
	return cmd
}
