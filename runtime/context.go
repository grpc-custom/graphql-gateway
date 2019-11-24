package runtime

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/textproto"
	"strings"

	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

const (
	MetadataPrefix = "graphqlgateway-"
)

func Context(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}

func AnnotateContext(ctx context.Context, req *http.Request) context.Context {
	md := annotateContext(ctx, req)
	return metadata.NewOutgoingContext(ctx, md)
}

func annotateContext(ctx context.Context, req *http.Request) metadata.MD {
	var pairs []string
	for key, vals := range req.Header {
		for _, val := range vals {
			key = textproto.CanonicalMIMEHeaderKey(key)
			if isPermanentHTTPHeader(key) {
				pairs = append(pairs, strings.ToLower(MetadataPrefix+key), val)
			}
		}
	}
	if auth := req.Header.Get(authorizationHeader); auth != "" {
		pairs = append(pairs, strings.ToLower(authorizationHeader), auth)
	}
	if addr := remoteAddr(req); addr != "" {
		pairs = append(pairs, strings.ToLower(xForwardedForHeader), addr)
	}
	if host := remoteHost(req); host != "" {
		pairs = append(pairs, strings.ToLower(xForwardedHostHeader), host)
	}
	md := metadata.Pairs(pairs...)
	return md
}

func remoteAddr(r *http.Request) string {
	addr := r.RemoteAddr
	if addr == "" {
		return ""
	}
	remoteIP, _, err := net.SplitHostPort(addr)
	if err != nil {
		grpclog.Infof("invalid remote addr: %s", addr)
		return ""
	}
	fwd := r.Header.Get(xForwardedForHeader)
	if fwd == "" {
		return remoteIP
	}
	return fmt.Sprintf("%s, %s", fwd, remoteIP)
}

func remoteHost(r *http.Request) string {
	host := r.Header.Get(xForwardedHostHeader)
	if host != "" {
		return host
	}
	return r.Host
}

func isPermanentHTTPHeader(h string) bool {
	switch h {
	case
		"Accept",
		"Accept-Charset",
		"Accept-Language",
		"Accept-Ranges",
		"Authorization",
		"Cache-Control",
		"Content-Type",
		"Cookie",
		"Date",
		"Expect",
		"From",
		"Host",
		"If-Match",
		"If-Modified-Since",
		"If-None-Match",
		"If-Schedule-Tag-Match",
		"If-Unmodified-Since",
		"Max-Forwards",
		"Origin",
		"Pragma",
		"Referer",
		"User-Agent",
		"Via",
		"Warning":
		return true
	}
	return false
}
