package polvo_client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"os"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	v1 "pkg.aiocean.dev/polvogo/aiocean/polvo/v1"
)

func GetPolvoClient (ctx context.Context) (v1.PolvoServiceClient, error){

	conn, err := newGrpcConnection(ctx, os.Getenv("POLVO_SERVICE_ADDRESS"))
	if err != nil {
		return nil, errors.WithMessage(err, "failed to dial")
	}

	client := v1.NewPolvoServiceClient(conn)

	return client, nil
}

func newGrpcConnection(ctx context.Context, host string) (*grpc.ClientConn, error) {

	var opts []grpc.DialOption

	if strings.HasSuffix(host, ":443") {
		opts = append(opts, grpc.WithAuthority(host))
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}

		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	grpcConn, err := grpc.DialContext(ctx, host, opts...)
	if err != nil {
		return nil, err
	}

	return grpcConn, nil
}
